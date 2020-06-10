// +build !ignore_autogenerated

// Code generated by operator-sdk-v0.13.0-x86_64-linux-gnu. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBCluster) DeepCopyInto(out *YBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBCluster.
func (in *YBCluster) DeepCopy() *YBCluster {
	if in == nil {
		return nil
	}
	out := new(YBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *YBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBClusterList) DeepCopyInto(out *YBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]YBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBClusterList.
func (in *YBClusterList) DeepCopy() *YBClusterList {
	if in == nil {
		return nil
	}
	out := new(YBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *YBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBClusterSpec) DeepCopyInto(out *YBClusterSpec) {
	*out = *in
	out.Image = in.Image
	out.TLS = in.TLS
	in.Master.DeepCopyInto(&out.Master)
	in.Tserver.DeepCopyInto(&out.Tserver)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBClusterSpec.
func (in *YBClusterSpec) DeepCopy() *YBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(YBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBClusterStatus) DeepCopyInto(out *YBClusterStatus) {
	*out = *in
	in.DataMoveChangeTime.DeepCopyInto(&out.DataMoveChangeTime)
	in.TSScaleDownChangeTime.DeepCopyInto(&out.TSScaleDownChangeTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBClusterStatus.
func (in *YBClusterStatus) DeepCopy() *YBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(YBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBGFlagSpec) DeepCopyInto(out *YBGFlagSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBGFlagSpec.
func (in *YBGFlagSpec) DeepCopy() *YBGFlagSpec {
	if in == nil {
		return nil
	}
	out := new(YBGFlagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBImageSpec) DeepCopyInto(out *YBImageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBImageSpec.
func (in *YBImageSpec) DeepCopy() *YBImageSpec {
	if in == nil {
		return nil
	}
	out := new(YBImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBMasterSpec) DeepCopyInto(out *YBMasterSpec) {
	*out = *in
	out.Storage = in.Storage
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Gflags != nil {
		in, out := &in.Gflags, &out.Gflags
		*out = make([]YBGFlagSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBMasterSpec.
func (in *YBMasterSpec) DeepCopy() *YBMasterSpec {
	if in == nil {
		return nil
	}
	out := new(YBMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBRootCASpec) DeepCopyInto(out *YBRootCASpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBRootCASpec.
func (in *YBRootCASpec) DeepCopy() *YBRootCASpec {
	if in == nil {
		return nil
	}
	out := new(YBRootCASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBStorageSpec) DeepCopyInto(out *YBStorageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBStorageSpec.
func (in *YBStorageSpec) DeepCopy() *YBStorageSpec {
	if in == nil {
		return nil
	}
	out := new(YBStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBTLSSpec) DeepCopyInto(out *YBTLSSpec) {
	*out = *in
	out.RootCA = in.RootCA
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBTLSSpec.
func (in *YBTLSSpec) DeepCopy() *YBTLSSpec {
	if in == nil {
		return nil
	}
	out := new(YBTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YBTServerSpec) DeepCopyInto(out *YBTServerSpec) {
	*out = *in
	out.Storage = in.Storage
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Gflags != nil {
		in, out := &in.Gflags, &out.Gflags
		*out = make([]YBGFlagSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YBTServerSpec.
func (in *YBTServerSpec) DeepCopy() *YBTServerSpec {
	if in == nil {
		return nil
	}
	out := new(YBTServerSpec)
	in.DeepCopyInto(out)
	return out
}
