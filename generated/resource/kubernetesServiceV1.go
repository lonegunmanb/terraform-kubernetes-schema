package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesServiceV1 = `{
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
              "description": "An unstructured key value map stored with the service that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the service. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the service, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description": "Namespace defines the space within which name of the service must be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this service that can be used by clients to determine when service has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this service. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard service's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "allocate_load_balancer_node_ports": {
              "description": "Defines if ` + "`" + `NodePorts` + "`" + ` will be automatically allocated for services with type ` + "`" + `LoadBalancer` + "`" + `. It may be set to ` + "`" + `false` + "`" + ` if the cluster load-balancer does not rely on ` + "`" + `NodePorts` + "`" + `.  If the caller requests specific ` + "`" + `NodePorts` + "`" + ` (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type ` + "`" + `LoadBalancer` + "`" + `. Default is ` + "`" + `true` + "`" + `. More info: https://kubernetes.io/docs/concepts/services-networking/service/#load-balancer-nodeport-allocation",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "cluster_ip": {
              "computed": true,
              "description": "The IP address of the service. It is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. ` + "`" + `None` + "`" + ` can be specified for headless services when proxying is not required. Ignored if type is ` + "`" + `ExternalName` + "`" + `. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cluster_ips": {
              "computed": true,
              "description": "List of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise creation of the service will fail. If this field is not specified, it will be initialized from the ` + "`" + `clusterIP` + "`" + ` field. If this field is specified, clients must ensure that ` + "`" + `clusterIPs[0]` + "`" + ` and ` + "`" + `clusterIP` + "`" + ` have the same value. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "external_ips": {
              "description": "A list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "external_name": {
              "description": "The external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid DNS name and requires ` + "`" + `type` + "`" + ` to be ` + "`" + `ExternalName` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_traffic_policy": {
              "computed": true,
              "description": "Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. ` + "`" + `Local` + "`" + ` preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. ` + "`" + `Cluster` + "`" + ` obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading. More info: https://kubernetes.io/docs/tutorials/services/source-ip/",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_node_port": {
              "computed": true,
              "description": "Specifies the Healthcheck NodePort for the service. Only effects when type is set to ` + "`" + `LoadBalancer` + "`" + ` and external_traffic_policy is set to ` + "`" + `Local` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "internal_traffic_policy": {
              "computed": true,
              "description": "Specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only. ` + "`" + `Cluster` + "`" + ` routes internal traffic to a Service to all endpoints. ` + "`" + `Local` + "`" + ` routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready. The default value is ` + "`" + `Cluster` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_families": {
              "computed": true,
              "description": "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "ip_family_policy": {
              "computed": true,
              "description": "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "load_balancer_class": {
              "description": "The class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix. This field can only be set when the Service type is ` + "`" + `LoadBalancer` + "`" + `. If not set, the default load balancer implementation is used. This field can only be set when creating or updating a Service to type ` + "`" + `LoadBalancer` + "`" + `. More info: https://kubernetes.io/docs/concepts/services-networking/service/#load-balancer-class",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "load_balancer_ip": {
              "description": "Only applies to ` + "`" + `type = LoadBalancer` + "`" + `. LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying this field when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "load_balancer_source_ranges": {
              "description": "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature. More info: http://kubernetes.io/docs/user-guide/services-firewalls",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "publish_not_ready_addresses": {
              "description": "When set to true, indicates that DNS implementations must publish the ` + "`" + `notReadyAddresses` + "`" + ` of subsets for the Endpoints associated with the Service. The default value is ` + "`" + `false` + "`" + `. The primary use case for setting this field is to use a StatefulSet's Headless Service to propagate ` + "`" + `SRV` + "`" + ` records for its Pods without respect to their readiness for purpose of peer discovery.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "selector": {
              "description": "Route service traffic to pods with label keys and values matching this selector. Only applies to types ` + "`" + `ClusterIP` + "`" + `, ` + "`" + `NodePort` + "`" + `, and ` + "`" + `LoadBalancer` + "`" + `. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "session_affinity": {
              "description": "Used to maintain session affinity. Supports ` + "`" + `ClientIP` + "`" + ` and ` + "`" + `None` + "`" + `. Defaults to ` + "`" + `None` + "`" + `. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Determines how the service is exposed. Defaults to ` + "`" + `ClusterIP` + "`" + `. Valid options are ` + "`" + `ExternalName` + "`" + `, ` + "`" + `ClusterIP` + "`" + `, ` + "`" + `NodePort` + "`" + `, and ` + "`" + `LoadBalancer` + "`" + `. ` + "`" + `ExternalName` + "`" + ` maps to the specified ` + "`" + `external_name` + "`" + `. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "port": {
              "block": {
                "attributes": {
                  "app_protocol": {
                    "description": "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The name of this port within the service. All ports within the service must have unique names. Optional if only one ServicePort is defined on this service.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_port": {
                    "computed": true,
                    "description": "The port on each node on which this service is exposed when ` + "`" + `type` + "`" + ` is ` + "`" + `NodePort` + "`" + ` or ` + "`" + `LoadBalancer` + "`" + `. Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the ` + "`" + `type` + "`" + ` of this service requires one. More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "port": {
                    "description": "The port that will be exposed by this service.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "protocol": {
                    "description": "The IP protocol for this port. Supports ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `. Default is ` + "`" + `TCP` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_port": {
                    "computed": true,
                    "description": "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. This field is ignored for services with ` + "`" + `cluster_ip = \"None\"` + "`" + `. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "session_affinity_config": {
              "block": {
                "block_types": {
                  "client_ip": {
                    "block": {
                      "attributes": {
                        "timeout_seconds": {
                          "computed": true,
                          "description": "Specifies the seconds of ` + "`" + `ClientIP` + "`" + ` type session sticky time. The value must be \u003e 0 and \u003c= 86400(for 1 day) if ` + "`" + `ServiceAffinity` + "`" + ` == ` + "`" + `ClientIP` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Contains the configurations of Client IP based session affinity.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Contains the configurations of session affinity. More info: https://kubernetes.io/docs/concepts/services-networking/service/#proxy-mode-ipvs",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Spec defines the behavior of a service. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description": "A Service is an abstraction which defines a logical set of pods and a policy by which to access them - sometimes called a micro-service.",
    "description_kind": "plain"
  },
  "version": 1
}`

func KubernetesServiceV1Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesServiceV1), &result)
	return &result
}
