package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesNetworkPolicy = `{
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
              "description": "An unstructured key value map stored with the network policy that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the network policy. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the network policy, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description": "Namespace defines the space within which name of the network policy must be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this network policy that can be used by clients to determine when network policy has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this network policy. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard network policy's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "policy_types": {
              "description": "policyTypes is a list of rule types that the NetworkPolicy relates to. Valid options are [\"Ingress\"], [\"Egress\"], or [\"Ingress\", \"Egress\"]. If this field is not specified, it will default based on the existence of ingress or egress rules; policies that contain an egress section are assumed to affect egress, and all policies (whether or not they contain an ingress section) are assumed to affect ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ \"Egress\" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include \"Egress\" (since such a policy would not include an egress section and would otherwise default to just [ \"Ingress\" ]). This field is beta-level in 1.8",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "egress": {
              "block": {
                "block_types": {
                  "ports": {
                    "block": {
                      "attributes": {
                        "port": {
                          "description": "port represents the port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "protocol": {
                          "description": "protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "ports is a list of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "to": {
                    "block": {
                      "block_types": {
                        "ip_block": {
                          "block": {
                            "attributes": {
                              "cidr": {
                                "description": "cidr is a string representing the IPBlock Valid examples are \"192.168.1.0/24\" or \"2001:db8::/64\"",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "except": {
                                "description": "except is a slice of CIDRs that should not be included within an IPBlock Valid examples are \"192.168.1.0/24\" or \"2001:db8::/64\" Except values will be rejected if they are outside the cidr range",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "ipBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "namespace_selector": {
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
                            "description": "namespaceSelector selects namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.\n\nIf podSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the namespaces selected by namespaceSelector. Otherwise it selects all pods in the namespaces selected by namespaceSelector.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "pod_selector": {
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
                            "description": "podSelector is a label selector which selects pods. This field follows standard label selector semantics; if present but empty, it selects all pods.\n\nIf namespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the pods matching podSelector in the policy's own namespace.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "to is a list of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "egress is a list of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ingress": {
              "block": {
                "block_types": {
                  "from": {
                    "block": {
                      "block_types": {
                        "ip_block": {
                          "block": {
                            "attributes": {
                              "cidr": {
                                "description": "cidr is a string representing the IPBlock Valid examples are \"192.168.1.0/24\" or \"2001:db8::/64\"",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "except": {
                                "description": "except is a slice of CIDRs that should not be included within an IPBlock Valid examples are \"192.168.1.0/24\" or \"2001:db8::/64\" Except values will be rejected if they are outside the cidr range",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "ipBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "namespace_selector": {
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
                            "description": "namespaceSelector selects namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.\n\nIf podSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the namespaces selected by namespaceSelector. Otherwise it selects all pods in the namespaces selected by namespaceSelector.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "pod_selector": {
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
                            "description": "podSelector is a label selector which selects pods. This field follows standard label selector semantics; if present but empty, it selects all pods.\n\nIf namespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the pods matching podSelector in the policy's own namespace.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "from is a list of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "ports": {
                    "block": {
                      "attributes": {
                        "port": {
                          "description": "port represents the port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "protocol": {
                          "description": "protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "ports is a list of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "ingress is a list of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "pod_selector": {
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
                "description": "podSelector selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "spec represents the specification of the desired behavior for this NetworkPolicy.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesNetworkPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesNetworkPolicy), &result)
	return &result
}
