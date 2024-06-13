package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesHorizontalPodAutoscalerV2Beta2 = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "An unstructured key value map stored with the horizontal pod autoscaler that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "generate_name": {
              "description": "Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "generation": {
              "computed": true,
              "description": "A sequence number representing a specific generation of the desired state.",
              "description_kind": "plain",
              "type": "number"
            },
            "labels": {
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the horizontal pod autoscaler. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the horizontal pod autoscaler, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description": "Namespace defines the space within which name of the horizontal pod autoscaler must be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this horizontal pod autoscaler that can be used by clients to determine when horizontal pod autoscaler has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this horizontal pod autoscaler. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard horizontal pod autoscaler's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "max_replicas": {
              "description": "Upper limit for the number of pods that can be set by the autoscaler.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_replicas": {
              "description": "Lower limit for the number of pods that can be set by the autoscaler, defaults to ` + "`" + `1` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_cpu_utilization_percentage": {
              "computed": true,
              "description": "Target average CPU utilization (represented as a percentage of requested CPU) over all the pods. If not specified the default autoscaling policy will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "behavior": {
              "block": {
                "block_types": {
                  "scale_down": {
                    "block": {
                      "attributes": {
                        "select_policy": {
                          "description": "Used to specify which policy should be used. If not set, the default value Max is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "stabilization_window_seconds": {
                          "description": "Number of seconds for which past recommendations should be considered while scaling up or scaling down. This value must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "policy": {
                          "block": {
                            "attributes": {
                              "period_seconds": {
                                "description": "Period specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "type": {
                                "description": "Type is used to specify the scaling policy: Percent or Pods",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "Value contains the amount of change which is permitted by the policy. It must be greater than zero.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description": "List of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the scaling rule will be discarded as invalid.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Scaling policy for scaling Down",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "scale_up": {
                    "block": {
                      "attributes": {
                        "select_policy": {
                          "description": "Used to specify which policy should be used. If not set, the default value Max is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "stabilization_window_seconds": {
                          "description": "Number of seconds for which past recommendations should be considered while scaling up or scaling down. This value must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "policy": {
                          "block": {
                            "attributes": {
                              "period_seconds": {
                                "description": "Period specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "type": {
                                "description": "Type is used to specify the scaling policy: Percent or Pods",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "Value contains the amount of change which is permitted by the policy. It must be greater than zero.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description": "List of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the scaling rule will be discarded as invalid.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Scaling policy for scaling Up",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Behavior configures the scaling behavior of the target in both Up and Down directions (` + "`" + `scale_up` + "`" + ` and ` + "`" + `scale_down` + "`" + ` fields respectively).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "metric": {
              "block": {
                "attributes": {
                  "type": {
                    "description": "type is the type of metric source. It should be one of \"ContainerResource\", \"External\", \"Object\", \"Pods\" or \"Resource\", each mapping to a matching field in the object. Note: \"ContainerResource\" type is available on when the feature-gate HPAContainerMetrics is enabled",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "container_resource": {
                    "block": {
                      "attributes": {
                        "container": {
                          "description": "name of the container in the pods of the scaling target",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "name of the resource in question",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "target": {
                          "block": {
                            "attributes": {
                              "average_utilization": {
                                "description": "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "average_value": {
                                "description": "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "type represents whether the metric type is Utilization, Value, or AverageValue",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "value is the target value of the metric (as a quantity).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "target specifies the target value for the given metric",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "external": {
                    "block": {
                      "block_types": {
                        "metric": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "name is the name of the given metric",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "selector": {
                                "block": {
                                  "attributes": {
                                    "match_labels": {
                                      "description": "A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of ` + "`" + `match_expressions` + "`" + `, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "match_expressions": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description": "The label key that the selector applies to.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "operator": {
                                            "description": "A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + ` and ` + "`" + `DoesNotExist` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "values": {
                                            "description": "An array of string values. If the operator is ` + "`" + `In` + "`" + ` or ` + "`" + `NotIn` + "`" + `, the values array must be non-empty. If the operator is ` + "`" + `Exists` + "`" + ` or ` + "`" + `DoesNotExist` + "`" + `, the values array must be empty. This array is replaced during a strategic merge patch.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "A list of label selector requirements. The requirements are ANDed.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "metric identifies the target metric by name and selector",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "target": {
                          "block": {
                            "attributes": {
                              "average_utilization": {
                                "description": "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "average_value": {
                                "description": "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "type represents whether the metric type is Utilization, Value, or AverageValue",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "value is the target value of the metric (as a quantity).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "target specifies the target value for the given metric",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "object": {
                    "block": {
                      "block_types": {
                        "described_object": {
                          "block": {
                            "attributes": {
                              "api_version": {
                                "description": "API version of the referent",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "kind": {
                                "description": "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "name": {
                                "description": "Name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "metric": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "name is the name of the given metric",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "selector": {
                                "block": {
                                  "attributes": {
                                    "match_labels": {
                                      "description": "A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of ` + "`" + `match_expressions` + "`" + `, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "match_expressions": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description": "The label key that the selector applies to.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "operator": {
                                            "description": "A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + ` and ` + "`" + `DoesNotExist` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "values": {
                                            "description": "An array of string values. If the operator is ` + "`" + `In` + "`" + ` or ` + "`" + `NotIn` + "`" + `, the values array must be non-empty. If the operator is ` + "`" + `Exists` + "`" + ` or ` + "`" + `DoesNotExist` + "`" + `, the values array must be empty. This array is replaced during a strategic merge patch.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "A list of label selector requirements. The requirements are ANDed.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "metric identifies the target metric by name and selector",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "target": {
                          "block": {
                            "attributes": {
                              "average_utilization": {
                                "description": "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "average_value": {
                                "description": "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "type represents whether the metric type is Utilization, Value, or AverageValue",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "value is the target value of the metric (as a quantity).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "target specifies the target value for the given metric",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "pods": {
                    "block": {
                      "block_types": {
                        "metric": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "name is the name of the given metric",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "selector": {
                                "block": {
                                  "attributes": {
                                    "match_labels": {
                                      "description": "A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of ` + "`" + `match_expressions` + "`" + `, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "match_expressions": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description": "The label key that the selector applies to.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "operator": {
                                            "description": "A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + ` and ` + "`" + `DoesNotExist` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "values": {
                                            "description": "An array of string values. If the operator is ` + "`" + `In` + "`" + ` or ` + "`" + `NotIn` + "`" + `, the values array must be non-empty. If the operator is ` + "`" + `Exists` + "`" + ` or ` + "`" + `DoesNotExist` + "`" + `, the values array must be empty. This array is replaced during a strategic merge patch.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "A list of label selector requirements. The requirements are ANDed.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "metric identifies the target metric by name and selector",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "target": {
                          "block": {
                            "attributes": {
                              "average_utilization": {
                                "description": "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "average_value": {
                                "description": "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "type represents whether the metric type is Utilization, Value, or AverageValue",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "value is the target value of the metric (as a quantity).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "target specifies the target value for the given metric",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "resource": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "name is the name of the resource in question.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "target": {
                          "block": {
                            "attributes": {
                              "average_utilization": {
                                "description": "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "average_value": {
                                "description": "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "type represents whether the metric type is Utilization, Value, or AverageValue",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "value is the target value of the metric (as a quantity).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Target specifies the target value for the given metric",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used). The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods. Ergo, metrics used must decrease as the pod count is increased, and vice-versa. See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "scale_target_ref": {
              "block": {
                "attributes": {
                  "api_version": {
                    "description": "API version of the referent",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kind": {
                    "description": "Kind of the referent. e.g. ` + "`" + `ReplicationController` + "`" + `. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Reference to scaled resource. e.g. Replication Controller",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "Horizontal Pod Autoscaler automatically scales the number of pods in a replication controller, deployment or replica set based on observed CPU utilization.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesHorizontalPodAutoscalerV2Beta2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesHorizontalPodAutoscalerV2Beta2), &result)
	return &result
}
