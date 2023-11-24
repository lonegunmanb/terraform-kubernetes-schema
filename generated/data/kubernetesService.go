package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesService = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spec": {
        "computed": true,
        "description": "Spec defines the behavior of a service. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocate_load_balancer_node_ports": "bool",
              "cluster_ip": "string",
              "cluster_ips": [
                "list",
                "string"
              ],
              "external_ips": [
                "set",
                "string"
              ],
              "external_name": "string",
              "external_traffic_policy": "string",
              "health_check_node_port": "number",
              "internal_traffic_policy": "string",
              "ip_families": [
                "list",
                "string"
              ],
              "ip_family_policy": "string",
              "load_balancer_class": "string",
              "load_balancer_ip": "string",
              "load_balancer_source_ranges": [
                "set",
                "string"
              ],
              "port": [
                "list",
                [
                  "object",
                  {
                    "app_protocol": "string",
                    "name": "string",
                    "node_port": "number",
                    "port": "number",
                    "protocol": "string",
                    "target_port": "string"
                  }
                ]
              ],
              "publish_not_ready_addresses": "bool",
              "selector": [
                "map",
                "string"
              ],
              "session_affinity": "string",
              "session_affinity_config": [
                "list",
                [
                  "object",
                  {
                    "client_ip": [
                      "list",
                      [
                        "object",
                        {
                          "timeout_seconds": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "type": "string"
            }
          ]
        ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesService), &result)
	return &result
}
