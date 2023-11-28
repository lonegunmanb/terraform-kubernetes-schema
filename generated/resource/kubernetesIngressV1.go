package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesIngressV1 = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "load_balancer": [
                "list",
                [
                  "object",
                  {
                    "ingress": [
                      "list",
                      [
                        "object",
                        {
                          "hostname": "string",
                          "ip": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "wait_for_load_balancer": {
        "description": "Terraform will wait for the load balancer to have at least 1 endpoint before considering the resource created.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "An unstructured key value map stored with the ingress that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the ingress. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the ingress, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description": "Namespace defines the space within which name of the ingress must be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this ingress that can be used by clients to determine when ingress has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this ingress. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard ingress's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "ingress_class_name": {
              "computed": true,
              "description": "ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -\u003e IngressClass -\u003e Ingress resource). Although the ` + "`" + `kubernetes.io/ingress.class` + "`" + ` annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "default_backend": {
              "block": {
                "block_types": {
                  "resource": {
                    "block": {
                      "attributes": {
                        "api_group": {
                          "description": "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "kind": {
                          "description": "The kind of resource.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "The name of the User to bind to.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "service": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Specifies the name of the referenced service.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "port": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Specifies the name of the port of the referenced service.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "number": {
                                "description": "Specifies the numerical port of the referenced service.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Specifies the port of the referenced service.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rule": {
              "block": {
                "attributes": {
                  "host": {
                    "description": "host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \"host\" part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to\n   the IP in the Spec of the parent Ingress.\n2. The ` + "`" + `:` + "`" + ` delimiter is not respected because ports are not allowed.\n\t  Currently the port of an Ingress is implicitly :80 for http and\n\t  :443 for https.\nBoth these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.\n\nhost can be \"precise\" which is a domain name without the terminating dot of a network host (e.g. \"foo.bar.com\") or \"wildcard\", which is a domain name prefixed with a single wildcard label (e.g. \"*.foo.com\"). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == \"*\"). Requests will be matched against the Host field in the following way: 1. If host is precise, the request matches this rule if the http host header is equal to Host. 2. If host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "http": {
                    "block": {
                      "block_types": {
                        "path": {
                          "block": {
                            "attributes": {
                              "path": {
                                "description": "path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \"path\" part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value \"Exact\" or \"Prefix\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path_type": {
                                "description": "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is\n  done on a path element by element basis. A path element refers is the\n  list of labels in the path split by the '/' separator. A request is a\n  match for path p if every p is an element-wise prefix of p of the\n  request path. Note that if the last element of the path is a substring\n  of the last element in request path, it is not a match (e.g. /foo/bar\n  matches /foo/bar/baz, but does not match /foo/barbaz).\n* ImplementationSpecific: Interpretation of the Path matching is up to\n  the IngressClass. Implementations can treat this as a separate PathType\n  or treat it identically to Prefix or Exact path types.\nImplementations are required to support all path types.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "backend": {
                                "block": {
                                  "block_types": {
                                    "resource": {
                                      "block": {
                                        "attributes": {
                                          "api_group": {
                                            "description": "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "kind": {
                                            "description": "The kind of resource.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "name": {
                                            "description": "The name of the User to bind to.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "service": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Specifies the name of the referenced service.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "port": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Specifies the name of the port of the referenced service.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "number": {
                                                  "description": "Specifies the numerical port of the referenced service.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Specifies the port of the referenced service.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "min_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Backend defines the referenced service endpoint to which the traffic will be forwarded to.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "http is a list of http selectors pointing to backends. In the example: http:///? -\u003e backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "tls": {
              "block": {
                "attributes": {
                  "hosts": {
                    "description": "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "secret_name": {
                    "description": "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the \"Host\" header is used for routing.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
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
  "version": 0
}`

func KubernetesIngressV1Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesIngressV1), &result)
	return &result
}
