// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/kubemove/v1alpha1.DataSync":          schema_pkg_apis_kubemove_v1alpha1_DataSync(ref),
		"./pkg/apis/kubemove/v1alpha1.DataSyncSpec":      schema_pkg_apis_kubemove_v1alpha1_DataSyncSpec(ref),
		"./pkg/apis/kubemove/v1alpha1.DataSyncStatus":    schema_pkg_apis_kubemove_v1alpha1_DataSyncStatus(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveEngine":        schema_pkg_apis_kubemove_v1alpha1_MoveEngine(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveEngineSpec":    schema_pkg_apis_kubemove_v1alpha1_MoveEngineSpec(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveEngineStatus":  schema_pkg_apis_kubemove_v1alpha1_MoveEngineStatus(ref),
		"./pkg/apis/kubemove/v1alpha1.MovePair":          schema_pkg_apis_kubemove_v1alpha1_MovePair(ref),
		"./pkg/apis/kubemove/v1alpha1.MovePairSpec":      schema_pkg_apis_kubemove_v1alpha1_MovePairSpec(ref),
		"./pkg/apis/kubemove/v1alpha1.MovePairStatus":    schema_pkg_apis_kubemove_v1alpha1_MovePairStatus(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveReverse":       schema_pkg_apis_kubemove_v1alpha1_MoveReverse(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveReverseSpec":   schema_pkg_apis_kubemove_v1alpha1_MoveReverseSpec(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveReverseStatus": schema_pkg_apis_kubemove_v1alpha1_MoveReverseStatus(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveSwitch":        schema_pkg_apis_kubemove_v1alpha1_MoveSwitch(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveSwitchSpec":    schema_pkg_apis_kubemove_v1alpha1_MoveSwitchSpec(ref),
		"./pkg/apis/kubemove/v1alpha1.MoveSwitchStatus":  schema_pkg_apis_kubemove_v1alpha1_MoveSwitchStatus(ref),
	}
}

func schema_pkg_apis_kubemove_v1alpha1_DataSync(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataSync is the Schema for the datasyncs API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/kubemove/v1alpha1.DataSyncSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/kubemove/v1alpha1.DataSyncStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/kubemove/v1alpha1.DataSyncSpec", "./pkg/apis/kubemove/v1alpha1.DataSyncStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_DataSyncSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataSyncSpec defines the desired state of DataSync",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"plugin": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"moveEngine": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"backup": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"restore": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"namespace", "plugin", "moveEngine", "backup", "restore", "config"},
			},
		},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_DataSyncStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataSyncStatus defines the observed state of DataSync",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"stage": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"completionTime": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"stage", "status", "completionTime", "reason"},
			},
		},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveEngine(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveEngine is the Schema for the moveengines API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/kubemove/v1alpha1.MoveEngineSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/kubemove/v1alpha1.MoveEngineStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/kubemove/v1alpha1.MoveEngineSpec", "./pkg/apis/kubemove/v1alpha1.MoveEngineStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveEngineSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveEngineSpec defines the desired state of MoveEngine",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"movePair": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"remoteNamespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"syncPeriod": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mode": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"plugin": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"includeResources": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"pluginParameters": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
						},
					},
				},
				Required: []string{"movePair", "namespace", "remoteNamespace", "syncPeriod", "mode", "plugin", "includeResources"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/runtime.RawExtension"},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveEngineStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveEngineStatus defines the observed state of MoveEngine",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"Status": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"LastStatus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"SyncedTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"LastSyncedTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"DataSync": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"DataSyncStatus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"Status", "LastStatus", "SyncedTime", "LastSyncedTime", "DataSync", "DataSyncStatus"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MovePair(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MovePair is the Schema for the movepairs API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/kubemove/v1alpha1.MovePairSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/kubemove/v1alpha1.MovePairStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/kubemove/v1alpha1.MovePairSpec", "./pkg/apis/kubemove/v1alpha1.MovePairStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MovePairSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MovePairSpec defines the desired state of MovePair",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("k8s.io/client-go/tools/clientcmd/api/v1.Config"),
						},
					},
				},
				Required: []string{"config"},
			},
		},
		Dependencies: []string{
			"k8s.io/client-go/tools/clientcmd/api/v1.Config"},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MovePairStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MovePairStatus defines the observed state of MovePair",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"state"},
			},
		},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveReverse(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveReverse is the Schema for the movereverses API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/kubemove/v1alpha1.MoveReverseSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/kubemove/v1alpha1.MoveReverseStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/kubemove/v1alpha1.MoveReverseSpec", "./pkg/apis/kubemove/v1alpha1.MoveReverseStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveReverseSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveReverseSpec defines the desired state of MoveReverse",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveReverseStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveReverseStatus defines the observed state of MoveReverse",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveSwitch(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveSwitch is the Schema for the moveswitches API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/kubemove/v1alpha1.MoveSwitchSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/kubemove/v1alpha1.MoveSwitchStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/kubemove/v1alpha1.MoveSwitchSpec", "./pkg/apis/kubemove/v1alpha1.MoveSwitchStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveSwitchSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveSwitchSpec defines the desired state of MoveSwitch",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"moveEngine": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"active": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"moveEngine", "active"},
			},
		},
	}
}

func schema_pkg_apis_kubemove_v1alpha1_MoveSwitchStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MoveSwitchStatus defines the observed state of MoveSwitch",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"Stage": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"Status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"Reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"Volume": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/kubemove/v1alpha1.VolumeStatus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"Stage", "Status", "Reason", "Volume"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/kubemove/v1alpha1.VolumeStatus"},
	}
}
