package ybcluster

import (
	"context"
	"fmt"

	"github.com/operator-framework/operator-sdk/pkg/status"
	yugabytev1alpha1 "github.com/yugabyte/yugabyte-k8s-operator/pkg/apis/yugabyte/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

// scaleTServers determines if TServers are going to be scaled up or
// scaled down. If scale down operation is required, it blacklists
// TServer pods. Retruns boolean indicating StatefulSet should be
// updated or not.
func (r *ReconcileYBCluster) scaleTServers(currentReplicas int32, cluster *yugabytev1alpha1.YBCluster) (bool, error) {
	// Ignore new/changed replica count if scale down
	// operation is in progress
	if cluster.Status.Conditions.IsFalseFor(scalingDownTServersCondition) {
		cluster.Status.TargetedTServerReplicas = cluster.Spec.Tserver.Replicas
		if err := r.client.Status().Update(context.TODO(), cluster); err != nil {
			return false, err
		}
	}
	scaleDownBy := currentReplicas - cluster.Status.TargetedTServerReplicas

	if scaleDownBy > 0 {
		logger.Infof("scaling down TServer replicas by %d.", scaleDownBy)

		tserverScaleCond := status.Condition{
			Type:    scalingDownTServersCondition,
			Status:  corev1.ConditionTrue,
			Reason:  status.ConditionReason("ScaleDownInProgress"),
			Message: "one or more TServer(s) are scaling down",
		}
		logger.Infof("updating Status condition %s: %s", tserverScaleCond.Type, tserverScaleCond.Status)
		cluster.Status.Conditions.SetCondition(tserverScaleCond)
		if err := r.client.Status().Update(context.TODO(), cluster); err != nil {
			return false, err
		}

		if err := r.blacklistPods(cluster, scaleDownBy); err != nil {
			return false, err
		}
	}

	return allowTServerStsUpdate(scaleDownBy, cluster)
}

// blacklistPods adds yugabyte.com/blacklist: true annotation to the
// TServer pods
func (r *ReconcileYBCluster) blacklistPods(cluster *yugabytev1alpha1.YBCluster, cnt int32) error {
	scalingDownTo := cluster.Status.TargetedTServerReplicas
	tserverReplicas := scalingDownTo + cnt
	for podNum := tserverReplicas - 1; podNum >= scalingDownTo; podNum-- {
		pod := &corev1.Pod{}
		err := r.client.Get(context.TODO(), types.NamespacedName{
			Namespace: cluster.GetNamespace(),
			Name:      fmt.Sprintf("%s-%d", tserverName, podNum),
		}, pod)
		if err != nil {
			return err
		}

		if pod.Annotations == nil {
			pod.SetAnnotations(map[string]string{blacklistAnnotation: "true"})
		} else if _, ok := pod.Annotations[blacklistAnnotation]; !ok {
			pod.Annotations[blacklistAnnotation] = "true"
		}

		if err = r.client.Update(context.TODO(), pod); err != nil {
			return err
		}
	}
	return nil
}

// allowTServerStsUpdate decides if TServer StatefulSet should be
// updated or not. Update is allowed for scale up directly. If it's a
// scale down, then it checks if data move operation has been
// completed.
func allowTServerStsUpdate(scaleDownBy int32, cluster *yugabytev1alpha1.YBCluster) (bool, error) {
	// Allow scale up directly
	if scaleDownBy <= 0 {
		return true, nil
	}

	dataMoveCond := cluster.Status.Conditions.GetCondition(movingDataCondition)
	tserverScaleDownCond := cluster.Status.Conditions.GetCondition(scalingDownTServersCondition)
	if dataMoveCond == nil || tserverScaleDownCond == nil {
		err := fmt.Errorf("status condition %s or %s is nil", movingDataCondition, scalingDownTServersCondition)
		logger.Error(err)
		return false, err
	}

	// Allow scale down if data move operation has been completed
	if dataMoveCond.IsFalse() {
		// TODO(bhavin192): add heartbeat time so that we can
		// handle the 0 tablets on a TServer case. Should have
		// gap of at least 5 minutes.
		if dataMoveCond.LastTransitionTime.After(tserverScaleDownCond.LastTransitionTime.Time) {
			return true, nil
		}
	}
	return false, nil
}
