package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesPod = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_state": {
        "description": "A list of the pod phases that indicate whether it was successfully created. Options: \"Pending\", \"Running\", \"Succeeded\", \"Failed\", \"Unknown\". Default: \"Running\". More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "An unstructured key value map stored with the pod that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the pod. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the pod, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description": "Namespace defines the space within which name of the pod must be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this pod that can be used by clients to determine when pod has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard pod's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "active_deadline_seconds": {
              "description": "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "automount_service_account_token": {
              "description": "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dns_policy": {
              "description": "Set DNS policy for containers within the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. Defaults to 'ClusterFirst'. More info: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_service_links": {
              "description": "Enables generating environment variables for service discovery. Defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "host_ipc": {
              "description": "Use the host's ipc namespace. Optional: Defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "host_network": {
              "description": "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "host_pid": {
              "description": "Use the host's pid namespace.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hostname": {
              "computed": true,
              "description": "Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_name": {
              "computed": true,
              "description": "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_selector": {
              "description": "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "priority_class_name": {
              "description": "If specified, indicates the pod's priority. \"system-node-critical\" and \"system-cluster-critical\" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "restart_policy": {
              "description": "Restart policy for all containers within the pod. One of Always, OnFailure, Never. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "runtime_class_name": {
              "description": "RuntimeClassName is a feature for selecting the container runtime configuration. The container runtime configuration is used to run a Pod's containers. More info: https://kubernetes.io/docs/concepts/containers/runtime-class",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduler_name": {
              "computed": true,
              "description": "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_account_name": {
              "computed": true,
              "description": "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: http://releases.k8s.io/HEAD/docs/design/service_accounts.md.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_process_namespace": {
              "description": "Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "subdomain": {
              "description": "If specified, the fully qualified Pod hostname will be \"...svc.\". If not specified, the pod will not have a domainname at all..",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "termination_grace_period_seconds": {
              "description": "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "affinity": {
              "block": {
                "block_types": {
                  "node_affinity": {
                    "block": {
                      "block_types": {
                        "preferred_during_scheduling_ignored_during_execution": {
                          "block": {
                            "attributes": {
                              "weight": {
                                "description": "weight is in the range 1-100",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "preference": {
                                "block": {
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
                                            "description": "Operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "values": {
                                            "description": "Values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "List of node selector requirements. The requirements are ANDed.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "match_fields": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description": "The label key that the selector applies to.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "operator": {
                                            "description": "A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + `, ` + "`" + `DoesNotExist` + "`" + `, ` + "`" + `Gt` + "`" + `, and ` + "`" + `Lt` + "`" + `.",
                                            "description_kind": "plain",
                                            "required": true,
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
                                        "description": "A list of node selector requirements by node's fields. The requirements are ANDed.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A node selector term, associated with the corresponding weight.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, RequiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding MatchExpressions; the node(s) with the highest sum are the most preferred.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "required_during_scheduling_ignored_during_execution": {
                          "block": {
                            "block_types": {
                              "node_selector_term": {
                                "block": {
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
                                            "description": "Operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "values": {
                                            "description": "Values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "List of node selector requirements. The requirements are ANDed.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "match_fields": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description": "The label key that the selector applies to.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "operator": {
                                            "description": "A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + `, ` + "`" + `DoesNotExist` + "`" + `, ` + "`" + `Gt` + "`" + `, and ` + "`" + `Lt` + "`" + `.",
                                            "description_kind": "plain",
                                            "required": true,
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
                                        "description": "A list of node selector requirements by node's fields. The requirements are ANDed.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "List of node selector terms. The terms are ORed.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a node label update), the system may or may not try to eventually evict the pod from its node.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Node affinity scheduling rules for the pod.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "pod_affinity": {
                    "block": {
                      "block_types": {
                        "preferred_during_scheduling_ignored_during_execution": {
                          "block": {
                            "attributes": {
                              "weight": {
                                "description": "weight associated with matching the corresponding podAffinityTerm, in the range 1-100",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "pod_affinity_term": {
                                "block": {
                                  "attributes": {
                                    "namespaces": {
                                      "description": "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "topology_key": {
                                      "description": "empty topology key is interpreted by the scheduler as 'all topologies'",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "label_selector": {
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
                                        "description": "A label query over a set of resources, in this case pods.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A pod affinity term, associated with the corresponding weight",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, RequiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding MatchExpressions; the node(s) with the highest sum are the most preferred.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "required_during_scheduling_ignored_during_execution": {
                          "block": {
                            "attributes": {
                              "namespaces": {
                                "description": "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "topology_key": {
                                "description": "empty topology key is interpreted by the scheduler as 'all topologies'",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "label_selector": {
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
                                  "description": "A label query over a set of resources, in this case pods.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each PodAffinityTerm are intersected, i.e. all terms must be satisfied.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.)",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "pod_anti_affinity": {
                    "block": {
                      "block_types": {
                        "preferred_during_scheduling_ignored_during_execution": {
                          "block": {
                            "attributes": {
                              "weight": {
                                "description": "weight associated with matching the corresponding podAffinityTerm, in the range 1-100",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "pod_affinity_term": {
                                "block": {
                                  "attributes": {
                                    "namespaces": {
                                      "description": "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "topology_key": {
                                      "description": "empty topology key is interpreted by the scheduler as 'all topologies'",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "label_selector": {
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
                                        "description": "A label query over a set of resources, in this case pods.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A pod affinity term, associated with the corresponding weight",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, RequiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding MatchExpressions; the node(s) with the highest sum are the most preferred.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "required_during_scheduling_ignored_during_execution": {
                          "block": {
                            "attributes": {
                              "namespaces": {
                                "description": "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "topology_key": {
                                "description": "empty topology key is interpreted by the scheduler as 'all topologies'",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "label_selector": {
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
                                  "description": "A label query over a set of resources, in this case pods.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each PodAffinityTerm are intersected, i.e. all terms must be satisfied.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.)",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional pod scheduling constraints.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "container": {
              "block": {
                "attributes": {
                  "args": {
                    "description": "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "command": {
                    "description": "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "image": {
                    "description": "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images/",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_pull_policy": {
                    "computed": true,
                    "description": "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images/#updating-images",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "stdin": {
                    "description": "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "stdin_once": {
                    "description": "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "termination_message_path": {
                    "description": "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "termination_message_policy": {
                    "computed": true,
                    "description": "Optional: Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tty": {
                    "description": "Whether this container should allocate a TTY for itself",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "working_dir": {
                    "description": "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "env": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the environment variable. Must be a C_IDENTIFIER",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "value_from": {
                          "block": {
                            "block_types": {
                              "config_map_key_ref": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description": "The key to select.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "name": {
                                      "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "optional": {
                                      "description": "Specify whether the ConfigMap or its key must be defined.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Selects a key of a ConfigMap.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "field_ref": {
                                "block": {
                                  "attributes": {
                                    "api_version": {
                                      "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "field_path": {
                                      "description": "Path of the field to select in the specified API version",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.podIP.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "resource_field_ref": {
                                "block": {
                                  "attributes": {
                                    "container_name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "divisor": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "resource": {
                                      "description": "Resource to select",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "secret_key_ref": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description": "The key of the secret to select from. Must be a valid secret key.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "name": {
                                      "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "optional": {
                                      "description": "Specify whether the Secret or its key must be defined.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Selects a key of a secret in the pod's namespace.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Source for the environment variable's value",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List of environment variables to set in the container. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "env_from": {
                    "block": {
                      "attributes": {
                        "prefix": {
                          "description": "An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "config_map_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "optional": {
                                "description": "Specify whether the ConfigMap must be defined",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "The ConfigMap to select from",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "optional": {
                                "description": "Specify whether the Secret must be defined",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "The Secret to select from",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "lifecycle": {
                    "block": {
                      "block_types": {
                        "post_start": {
                          "block": {
                            "block_types": {
                              "exec": {
                                "block": {
                                  "attributes": {
                                    "command": {
                                      "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "exec specifies the action to take.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "http_get": {
                                "block": {
                                  "attributes": {
                                    "host": {
                                      "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "path": {
                                      "description": "Path to access on the HTTP server.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "port": {
                                      "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "scheme": {
                                      "description": "Scheme to use for connecting to the host.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "http_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "The header field name",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "description": "The header field value",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Scheme to use for connecting to the host.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Specifies the http request to perform.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "tcp_socket": {
                                "block": {
                                  "attributes": {
                                    "port": {
                                      "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "post_start is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "pre_stop": {
                          "block": {
                            "block_types": {
                              "exec": {
                                "block": {
                                  "attributes": {
                                    "command": {
                                      "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "exec specifies the action to take.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "http_get": {
                                "block": {
                                  "attributes": {
                                    "host": {
                                      "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "path": {
                                      "description": "Path to access on the HTTP server.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "port": {
                                      "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "scheme": {
                                      "description": "Scheme to use for connecting to the host.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "http_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "The header field name",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "description": "The header field value",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Scheme to use for connecting to the host.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Specifies the http request to perform.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "tcp_socket": {
                                "block": {
                                  "attributes": {
                                    "port": {
                                      "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "pre_stop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Actions that the management system should take in response to container lifecycle events",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "liveness_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Minimum consecutive successes for the probe to be considered successful after having failed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "exec specifies the action to take.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number of the port to access on the container. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GRPC specifies an action involving a GRPC port.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_header": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Scheme to use for connecting to the host.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the http request to perform.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "port": {
                    "block": {
                      "attributes": {
                        "container_port": {
                          "description": "Number of port to expose on the pod's IP address. This must be a valid port number, 0 \u003c x \u003c 65536.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "host_ip": {
                          "description": "What host IP to bind the external port to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_port": {
                          "description": "Number of port to expose on the host. If specified, this must be a valid port number, 0 \u003c x \u003c 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "name": {
                          "description": "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "protocol": {
                          "description": "Protocol for port. Must be UDP or TCP. Defaults to \"TCP\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \"0.0.0.0\" address inside a container will be accessible from the network. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "readiness_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Minimum consecutive successes for the probe to be considered successful after having failed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "exec specifies the action to take.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number of the port to access on the container. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GRPC specifies an action involving a GRPC port.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_header": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Scheme to use for connecting to the host.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the http request to perform.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "resources": {
                    "block": {
                      "attributes": {
                        "limits": {
                          "computed": true,
                          "description": "Describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "requests": {
                          "computed": true,
                          "description": "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "security_context": {
                    "block": {
                      "attributes": {
                        "allow_privilege_escalation": {
                          "description": "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "privileged": {
                          "description": "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "read_only_root_filesystem": {
                          "description": "Whether this container has a read-only root filesystem. Default is false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "run_as_group": {
                          "description": "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "run_as_non_root": {
                          "description": "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "run_as_user": {
                          "description": "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "capabilities": {
                          "block": {
                            "attributes": {
                              "add": {
                                "description": "Added capabilities",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "drop": {
                                "description": "Removed capabilities",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "se_linux_options": {
                          "block": {
                            "attributes": {
                              "level": {
                                "description": "Level is SELinux level label that applies to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "role": {
                                "description": "Role is a SELinux role label that applies to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "Type is a SELinux type label that applies to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "user": {
                                "description": "User is a SELinux user label that applies to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "seccomp_profile": {
                          "block": {
                            "attributes": {
                              "localhost_profile": {
                                "description": "Localhost Profile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "Type indicates which kind of seccomp profile will be applied. Valid options are: Localhost, RuntimeDefault, Unconfined.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Security options the pod should run with. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "startup_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Minimum consecutive successes for the probe to be considered successful after having failed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "exec specifies the action to take.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number of the port to access on the container. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GRPC specifies an action involving a GRPC port.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_header": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Scheme to use for connecting to the host.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the http request to perform.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. This is an alpha feature enabled by the StartupProbe feature flag. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "volume_mount": {
                    "block": {
                      "attributes": {
                        "mount_path": {
                          "description": "Path within the container at which the volume should be mounted. Must not contain ':'.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "mount_propagation": {
                          "description": "Mount propagation mode. mount_propagation determines how mounts are propagated from the host to container and the other way around. Valid values are None (default), HostToContainer and Bidirectional.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "This must match the Name of a Volume.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "sub_path": {
                          "description": "Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Pod volumes to mount into the container's filesystem. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "dns_config": {
              "block": {
                "attributes": {
                  "nameservers": {
                    "description": "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "searches": {
                    "description": "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "option": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the option.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Value of the option. Optional: Defaults to empty.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. Optional: Defaults to empty",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "host_aliases": {
              "block": {
                "attributes": {
                  "hostnames": {
                    "description": "Hostnames for the IP address.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ip": {
                    "description": "IP address of the host file entry.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "List of hosts and IPs that will be injected into the pod's hosts file if specified. Optional: Defaults to empty.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "image_pull_secrets": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "init_container": {
              "block": {
                "attributes": {
                  "args": {
                    "description": "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "command": {
                    "description": "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "image": {
                    "description": "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images/",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_pull_policy": {
                    "computed": true,
                    "description": "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images/#updating-images",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "stdin": {
                    "description": "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "stdin_once": {
                    "description": "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "termination_message_path": {
                    "description": "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "termination_message_policy": {
                    "computed": true,
                    "description": "Optional: Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tty": {
                    "description": "Whether this container should allocate a TTY for itself",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "working_dir": {
                    "description": "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "env": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the environment variable. Must be a C_IDENTIFIER",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "value_from": {
                          "block": {
                            "block_types": {
                              "config_map_key_ref": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description": "The key to select.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "name": {
                                      "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "optional": {
                                      "description": "Specify whether the ConfigMap or its key must be defined.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Selects a key of a ConfigMap.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "field_ref": {
                                "block": {
                                  "attributes": {
                                    "api_version": {
                                      "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "field_path": {
                                      "description": "Path of the field to select in the specified API version",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.podIP.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "resource_field_ref": {
                                "block": {
                                  "attributes": {
                                    "container_name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "divisor": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "resource": {
                                      "description": "Resource to select",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "secret_key_ref": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description": "The key of the secret to select from. Must be a valid secret key.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "name": {
                                      "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "optional": {
                                      "description": "Specify whether the Secret or its key must be defined.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Selects a key of a secret in the pod's namespace.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Source for the environment variable's value",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List of environment variables to set in the container. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "env_from": {
                    "block": {
                      "attributes": {
                        "prefix": {
                          "description": "An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "config_map_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "optional": {
                                "description": "Specify whether the ConfigMap must be defined",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "The ConfigMap to select from",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "optional": {
                                "description": "Specify whether the Secret must be defined",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "The Secret to select from",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "lifecycle": {
                    "block": {
                      "block_types": {
                        "post_start": {
                          "block": {
                            "block_types": {
                              "exec": {
                                "block": {
                                  "attributes": {
                                    "command": {
                                      "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "exec specifies the action to take.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "http_get": {
                                "block": {
                                  "attributes": {
                                    "host": {
                                      "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "path": {
                                      "description": "Path to access on the HTTP server.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "port": {
                                      "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "scheme": {
                                      "description": "Scheme to use for connecting to the host.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "http_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "The header field name",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "description": "The header field value",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Scheme to use for connecting to the host.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Specifies the http request to perform.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "tcp_socket": {
                                "block": {
                                  "attributes": {
                                    "port": {
                                      "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "post_start is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "pre_stop": {
                          "block": {
                            "block_types": {
                              "exec": {
                                "block": {
                                  "attributes": {
                                    "command": {
                                      "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "exec specifies the action to take.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "http_get": {
                                "block": {
                                  "attributes": {
                                    "host": {
                                      "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "path": {
                                      "description": "Path to access on the HTTP server.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "port": {
                                      "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "scheme": {
                                      "description": "Scheme to use for connecting to the host.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "http_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "The header field name",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "description": "The header field value",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Scheme to use for connecting to the host.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Specifies the http request to perform.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "tcp_socket": {
                                "block": {
                                  "attributes": {
                                    "port": {
                                      "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "pre_stop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Actions that the management system should take in response to container lifecycle events",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "liveness_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Minimum consecutive successes for the probe to be considered successful after having failed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "exec specifies the action to take.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number of the port to access on the container. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GRPC specifies an action involving a GRPC port.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_header": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Scheme to use for connecting to the host.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the http request to perform.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "port": {
                    "block": {
                      "attributes": {
                        "container_port": {
                          "description": "Number of port to expose on the pod's IP address. This must be a valid port number, 0 \u003c x \u003c 65536.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "host_ip": {
                          "description": "What host IP to bind the external port to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_port": {
                          "description": "Number of port to expose on the host. If specified, this must be a valid port number, 0 \u003c x \u003c 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "name": {
                          "description": "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "protocol": {
                          "description": "Protocol for port. Must be UDP or TCP. Defaults to \"TCP\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \"0.0.0.0\" address inside a container will be accessible from the network. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "readiness_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Minimum consecutive successes for the probe to be considered successful after having failed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "exec specifies the action to take.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number of the port to access on the container. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GRPC specifies an action involving a GRPC port.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_header": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Scheme to use for connecting to the host.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the http request to perform.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "resources": {
                    "block": {
                      "attributes": {
                        "limits": {
                          "computed": true,
                          "description": "Describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "requests": {
                          "computed": true,
                          "description": "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "security_context": {
                    "block": {
                      "attributes": {
                        "allow_privilege_escalation": {
                          "description": "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "privileged": {
                          "description": "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "read_only_root_filesystem": {
                          "description": "Whether this container has a read-only root filesystem. Default is false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "run_as_group": {
                          "description": "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "run_as_non_root": {
                          "description": "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "run_as_user": {
                          "description": "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "capabilities": {
                          "block": {
                            "attributes": {
                              "add": {
                                "description": "Added capabilities",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "drop": {
                                "description": "Removed capabilities",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "se_linux_options": {
                          "block": {
                            "attributes": {
                              "level": {
                                "description": "Level is SELinux level label that applies to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "role": {
                                "description": "Role is a SELinux role label that applies to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "Type is a SELinux type label that applies to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "user": {
                                "description": "User is a SELinux user label that applies to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "seccomp_profile": {
                          "block": {
                            "attributes": {
                              "localhost_profile": {
                                "description": "Localhost Profile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "Type indicates which kind of seccomp profile will be applied. Valid options are: Localhost, RuntimeDefault, Unconfined.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Security options the pod should run with. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "startup_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Minimum consecutive successes for the probe to be considered successful after having failed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "exec specifies the action to take.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number of the port to access on the container. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GRPC specifies an action involving a GRPC port.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_header": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Scheme to use for connecting to the host.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the http request to perform.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. This is an alpha feature enabled by the StartupProbe feature flag. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "volume_mount": {
                    "block": {
                      "attributes": {
                        "mount_path": {
                          "description": "Path within the container at which the volume should be mounted. Must not contain ':'.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "mount_propagation": {
                          "description": "Mount propagation mode. mount_propagation determines how mounts are propagated from the host to container and the other way around. Valid values are None (default), HostToContainer and Bidirectional.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "This must match the Name of a Volume.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "sub_path": {
                          "description": "Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Pod volumes to mount into the container's filesystem. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "List of init containers belonging to the pod. Init containers always run to completion and each must complete successfully before the next is started. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "os": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name is the name of the operating system. The currently supported values are linux and windows.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specifies the OS of the containers in the pod.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "readiness_gate": {
              "block": {
                "attributes": {
                  "condition_type": {
                    "description": "refers to a condition in the pod's condition list with matching type.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to \"True\" More info: https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "security_context": {
              "block": {
                "attributes": {
                  "fs_group": {
                    "description": "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fs_group_change_policy": {
                    "description": "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "run_as_group": {
                    "description": "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "run_as_non_root": {
                    "description": "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "run_as_user": {
                    "description": "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "supplemental_groups": {
                    "description": "A list of groups applied to the first process run in each container, in addition to the container's primary GID. If unspecified, no groups will be added to any container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "number"
                    ]
                  }
                },
                "block_types": {
                  "se_linux_options": {
                    "block": {
                      "attributes": {
                        "level": {
                          "description": "Level is SELinux level label that applies to the container.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role": {
                          "description": "Role is a SELinux role label that applies to the container.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Type is a SELinux type label that applies to the container.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user": {
                          "description": "User is a SELinux user label that applies to the container.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "seccomp_profile": {
                    "block": {
                      "attributes": {
                        "localhost_profile": {
                          "description": "Localhost Profile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Type indicates which kind of seccomp profile will be applied. Valid options are: Localhost, RuntimeDefault, Unconfined.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "sysctl": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of a property to set.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Value of a property to set.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "holds a list of namespaced sysctls used for the pod.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "windows_options": {
                    "block": {
                      "attributes": {
                        "gmsa_credential_spec": {
                          "description": "GMSACredentialSpec is where the GMSA admission webhook inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "gmsa_credential_spec_name": {
                          "description": "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_process": {
                          "description": "HostProcess determines if a container should be run as a 'Host Process' container. Default value is false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "run_as_username": {
                          "description": "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "toleration": {
              "block": {
                "attributes": {
                  "effect": {
                    "description": "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
                    "description": "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "toleration_seconds": {
                    "description": "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "If specified, the pod's toleration. Optional: Defaults to empty",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "topology_spread_constraint": {
              "block": {
                "attributes": {
                  "max_skew": {
                    "description": "describes the degree to which pods may be unevenly distributed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "topology_key": {
                    "description": "the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "when_unsatisfiable": {
                    "description": "indicates how to deal with a pod if it doesn't satisfy the spread constraint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "label_selector": {
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
                      "description": "A label query over a set of resources, in this case pods.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "volume": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "aws_elastic_block_store": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "partition": {
                          "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "read_only": {
                          "description": "Whether to set the read-only property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "volume_id": {
                          "description": "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "azure_disk": {
                    "block": {
                      "attributes": {
                        "caching_mode": {
                          "description": "Host Caching mode: None, Read Only, Read Write.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "data_disk_uri": {
                          "description": "The URI the data disk in the blob storage",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "disk_name": {
                          "description": "The Name of the data disk in the blob storage",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kind": {
                          "computed": true,
                          "description": "The type for the data disk. Expected values: Shared, Dedicated, Managed. Defaults to Shared",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Represents an Azure Data Disk mount on the host and bind mount to the pod.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "azure_file": {
                    "block": {
                      "attributes": {
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "secret_name": {
                          "description": "The name of secret that contains Azure Storage Account Name and Key",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_namespace": {
                          "description": "The namespace of the secret that contains Azure Storage Account Name and Key. For Kubernetes up to 1.18.x the default is the same as the Pod. For Kubernetes 1.19.x and later the default is \"default\" namespace.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "share_name": {
                          "description": "Share Name",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an Azure File Service mount on the host and bind mount to the pod.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ceph_fs": {
                    "block": {
                      "attributes": {
                        "monitors": {
                          "description": "Monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "path": {
                          "description": "Used as the mounted root, rather than the full Ceph tree, default is /",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to ` + "`" + `false` + "`" + ` (read/write). More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "secret_file": {
                          "description": "The path to key ring for User, default is ` + "`" + `/etc/ceph/user.secret` + "`" + `. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user": {
                          "description": "User is the rados user name, default is admin. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents a Ceph FS mount on the host that shares a pod's lifetime",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "cinder": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "volume_id": {
                          "description": "Volume ID used to identify the volume in Cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "config_map": {
                    "block": {
                      "attributes": {
                        "default_mode": {
                          "description": "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "optional": {
                          "description": "Optional: Specify whether the ConfigMap or its keys must be defined.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "items": {
                          "block": {
                            "attributes": {
                              "key": {
                                "description": "The key to project.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "mode": {
                                "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "ConfigMap represents a configMap that should populate this volume",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "csi": {
                    "block": {
                      "attributes": {
                        "driver": {
                          "description": "the name of the volume driver to use. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to set the read-only property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: https://kubernetes.io/docs/concepts/storage/volumes#csi",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "volume_attributes": {
                          "description": "Attributes of the volume to publish.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "node_publish_secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents a CSI Volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#csi",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "downward_api": {
                    "block": {
                      "attributes": {
                        "default_mode": {
                          "description": "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "items": {
                          "block": {
                            "attributes": {
                              "mode": {
                                "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "field_ref": {
                                "block": {
                                  "attributes": {
                                    "api_version": {
                                      "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "field_path": {
                                      "description": "Path of the field to select in the specified API version",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              },
                              "resource_field_ref": {
                                "block": {
                                  "attributes": {
                                    "container_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "divisor": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "resource": {
                                      "description": "Resource to select",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "DownwardAPI represents downward API about the pod that should populate this volume",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "empty_dir": {
                    "block": {
                      "attributes": {
                        "medium": {
                          "description": "What type of storage medium should back this directory. The default is \"\" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "size_limit": {
                          "description": "Total amount of local storage required for this EmptyDir volume.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ephemeral": {
                    "block": {
                      "block_types": {
                        "volume_claim_template": {
                          "block": {
                            "block_types": {
                              "metadata": {
                                "block": {
                                  "attributes": {
                                    "annotations": {
                                      "description": "An unstructured key value map stored with the persistent volume claim that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "labels": {
                                      "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the persistent volume claim. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "spec": {
                                "block": {
                                  "attributes": {
                                    "access_modes": {
                                      "description": "A set of the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "storage_class_name": {
                                      "computed": true,
                                      "description": "Name of the storage class requested by the claim",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "volume_mode": {
                                      "computed": true,
                                      "description": "Defines what type of volume is required by the claim.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "volume_name": {
                                      "computed": true,
                                      "description": "The binding reference to the PersistentVolume backing this claim.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "resources": {
                                      "block": {
                                        "attributes": {
                                          "limits": {
                                            "description": "Map describing the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          },
                                          "requests": {
                                            "description": "Map describing the minimum amount of compute resources required. If this is omitted for a container, it defaults to ` + "`" + `limits` + "`" + ` if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "A list of the minimum resources the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    },
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
                                        "description": "A label query over volumes to consider for binding.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents an ephemeral volume that is handled by a normal storage driver. More info: https://kubernetes.io/docs/concepts/storage/ephemeral-volumes/#generic-ephemeral-volumes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "fc": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lun": {
                          "description": "FC target lun number",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "target_ww_ns": {
                          "description": "FC target worldwide names (WWNs)",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description": "Represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "flex_volume": {
                    "block": {
                      "attributes": {
                        "driver": {
                          "description": "Driver is the name of the driver to use for this volume.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". The default filesystem depends on FlexVolume script.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "options": {
                          "description": "Extra command options if any.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "read_only": {
                          "description": "Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents a generic volume resource that is provisioned/attached using an exec based plugin. This is an alpha feature and may change in future.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "flocker": {
                    "block": {
                      "attributes": {
                        "dataset_name": {
                          "description": "Name of the dataset stored as metadata -\u003e name on the dataset for Flocker should be considered as deprecated",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "dataset_uuid": {
                          "description": "UUID of the dataset. This is unique identifier of a Flocker dataset",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "gce_persistent_disk": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "partition": {
                          "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "pd_name": {
                          "description": "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "git_repo": {
                    "block": {
                      "attributes": {
                        "directory": {
                          "description": "Target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "repository": {
                          "description": "Repository URL",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "revision": {
                          "description": "Commit hash for the specified revision.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "GitRepo represents a git repository at a particular revision.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "glusterfs": {
                    "block": {
                      "attributes": {
                        "endpoints_name": {
                          "description": "The endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description": "The Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "host_path": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Type for HostPath volume. Allowed values are \"\" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "iscsi": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "iqn": {
                          "description": "Target iSCSI Qualified Name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "iscsi_interface": {
                          "description": "iSCSI interface name that uses an iSCSI transport. Defaults to 'default' (tcp).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lun": {
                          "description": "iSCSI target lun number.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "target_portal": {
                          "description": "iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "local": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#local",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a mounted local storage device such as a disk, partition or directory. Local volumes can only be used as a statically created PersistentVolume. Dynamic provisioning is not supported yet. More info: https://kubernetes.io/docs/concepts/storage/volumes#local",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "nfs": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "server": {
                          "description": "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "persistent_volume_claim": {
                    "block": {
                      "attributes": {
                        "claim_name": {
                          "description": "ClaimName is the name of a PersistentVolumeClaim in the same ",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Will force the ReadOnly setting in VolumeMounts.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "The specification of a persistent volume.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "photon_persistent_disk": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "pd_id": {
                          "description": "ID that identifies Photon Controller persistent disk",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a PhotonController persistent disk attached and mounted on kubelets host machine",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "projected": {
                    "block": {
                      "attributes": {
                        "default_mode": {
                          "description": "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "sources": {
                          "block": {
                            "block_types": {
                              "config_map": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "optional": {
                                      "description": "Optional: Specify whether the ConfigMap or it's keys must be defined.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "items": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description": "The key to project.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "mode": {
                                            "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "path": {
                                            "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "ConfigMap represents a configMap that should populate this volume",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "downward_api": {
                                "block": {
                                  "block_types": {
                                    "items": {
                                      "block": {
                                        "attributes": {
                                          "mode": {
                                            "description": "Mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "path": {
                                            "description": "Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "field_ref": {
                                            "block": {
                                              "attributes": {
                                                "api_version": {
                                                  "description": "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "field_path": {
                                                  "description": "Path of the field to select in the specified API version",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Selects a field of the pod: only annotations, labels, name and namespace are supported.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "resource_field_ref": {
                                            "block": {
                                              "attributes": {
                                                "container_name": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "divisor": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "resource": {
                                                  "description": "Resource to select",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "DownwardAPI represents downward API about the pod that should populate this volume",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "secret": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secrets",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "optional": {
                                      "description": "Optional: Specify whether the Secret or it's keys must be defined.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "items": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description": "The key to project.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "mode": {
                                            "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "path": {
                                            "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secrets",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "service_account_token": {
                                "block": {
                                  "attributes": {
                                    "audience": {
                                      "description": "Audience is the intended audience of the token",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "expiration_seconds": {
                                      "description": "ExpirationSeconds is the expected duration of validity of the service account token. It defaults to 1 hour and must be at least 10 minutes (600 seconds).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "path": {
                                      "description": "Path specifies a relative path to the mount point of the projected volume.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A projected service account token volume",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Source of the volume to project in the directory.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Projected represents a single volume that projects several volume sources into the same directory. More info: https://kubernetes.io/docs/concepts/storage/volumes/#projected",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "quobyte": {
                    "block": {
                      "attributes": {
                        "group": {
                          "description": "Group to map volume access to Default is no group",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "registry": {
                          "description": "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "user": {
                          "description": "User to map volume access to Defaults to serivceaccount user",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "volume": {
                          "description": "Volume is a string that references an already created Quobyte volume by name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "rbd": {
                    "block": {
                      "attributes": {
                        "ceph_monitors": {
                          "description": "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "fs_type": {
                          "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "keyring": {
                          "computed": true,
                          "description": "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "rados_user": {
                          "description": "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "rbd_image": {
                          "description": "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "rbd_pool": {
                          "description": "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "secret": {
                    "block": {
                      "attributes": {
                        "default_mode": {
                          "description": "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "optional": {
                          "description": "Optional: Specify whether the Secret or its keys must be defined.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "secret_name": {
                          "description": "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secrets",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "items": {
                          "block": {
                            "attributes": {
                              "key": {
                                "description": "The key to project.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "mode": {
                                "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secrets",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "vsphere_volume": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "volume_path": {
                          "description": "Path that identifies vSphere volume vmdk",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a vSphere volume attached and mounted on kubelets host machine",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Specification of the desired behavior of the pod.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func KubernetesPodSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesPod), &result)
	return &result
}
