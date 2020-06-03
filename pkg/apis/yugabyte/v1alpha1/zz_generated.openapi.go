// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/yugabyte/v1alpha1.YBCluster":       schema_pkg_apis_yugabyte_v1alpha1_YBCluster(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBClusterSpec":   schema_pkg_apis_yugabyte_v1alpha1_YBClusterSpec(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBClusterStatus": schema_pkg_apis_yugabyte_v1alpha1_YBClusterStatus(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBGFlagSpec":     schema_pkg_apis_yugabyte_v1alpha1_YBGFlagSpec(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBImageSpec":     schema_pkg_apis_yugabyte_v1alpha1_YBImageSpec(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBMasterSpec":    schema_pkg_apis_yugabyte_v1alpha1_YBMasterSpec(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBRootCASpec":    schema_pkg_apis_yugabyte_v1alpha1_YBRootCASpec(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBStorageSpec":   schema_pkg_apis_yugabyte_v1alpha1_YBStorageSpec(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBTLSSpec":       schema_pkg_apis_yugabyte_v1alpha1_YBTLSSpec(ref),
		"./pkg/apis/yugabyte/v1alpha1.YBTServerSpec":   schema_pkg_apis_yugabyte_v1alpha1_YBTServerSpec(ref),
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBCluster is the Schema for the ybclusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/yugabyte/v1alpha1.YBClusterSpec", "./pkg/apis/yugabyte/v1alpha1.YBClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBClusterSpec defines the desired state of YBCluster",
				Properties: map[string]spec.Schema{
					"replicationFactor": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBImageSpec"),
						},
					},
					"tls": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBTLSSpec"),
						},
					},
					"master": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBMasterSpec"),
						},
					},
					"tserver": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBTServerSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/yugabyte/v1alpha1.YBImageSpec", "./pkg/apis/yugabyte/v1alpha1.YBMasterSpec", "./pkg/apis/yugabyte/v1alpha1.YBTLSSpec", "./pkg/apis/yugabyte/v1alpha1.YBTServerSpec"},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBClusterStatus defines the observed state of YBCluster",
				Properties: map[string]spec.Schema{
					"masterReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"tserverReplicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"dataMoveCond": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"dataMoveChangeTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"tServerScaleDownCond": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"tSScaleDownChangeTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"masterReplicas", "tserverReplicas", "dataMoveCond", "dataMoveChangeTime", "tServerScaleDownCond", "tSScaleDownChangeTime"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBGFlagSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBGFlagSpec defines key-value pairs for each GFlag.",
				Properties: map[string]spec.Schema{
					"key": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBImageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBImageSpec defines docker image specific attributes.",
				Properties: map[string]spec.Schema{
					"repository": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"tag": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"pullPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBMasterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBMasterSpec defines attributes for YBMaster pods.",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"masterUIPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"masterRPCPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"enableLoadBalancer": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"podManagementPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"storage": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBStorageSpec"),
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"gflags": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBGFlagSpec"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/yugabyte/v1alpha1.YBGFlagSpec", "./pkg/apis/yugabyte/v1alpha1.YBStorageSpec", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBRootCASpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBRootCASpec defines Root CA cert & key attributes required for enabling TLS encryption.",
				Properties: map[string]spec.Schema{
					"cert": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"key": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBStorageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBStorageSpec defines storage specific attributes for YBMaster/YBTserver pods.",
				Properties: map[string]spec.Schema{
					"count": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"size": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"storageClass": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBTLSSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBTLSSpec defines TLS encryption specific attributes",
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"rootCA": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBRootCASpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/yugabyte/v1alpha1.YBRootCASpec"},
	}
}

func schema_pkg_apis_yugabyte_v1alpha1_YBTServerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "YBTServerSpec defines attributes for YBTServer pods.",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"tserverUIPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"tserverRPCPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ycqlPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"yedisPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ysqlPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"enableLoadBalancer": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"podManagementPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"storage": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBStorageSpec"),
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"gflags": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/yugabyte/v1alpha1.YBGFlagSpec"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/yugabyte/v1alpha1.YBGFlagSpec", "./pkg/apis/yugabyte/v1alpha1.YBStorageSpec", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}
