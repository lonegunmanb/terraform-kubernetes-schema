package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesNodes = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nodes": {
        "computed": true,
        "description": "List of nodes in a cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "metadata": [
                "list",
                [
                  "object",
                  {
                    "annotations": [
                      "map",
                      "string"
                    ],
                    "generation": "number",
                    "labels": [
                      "map",
                      "string"
                    ],
                    "name": "string",
                    "resource_version": "string",
                    "uid": "string"
                  }
                ]
              ],
              "spec": [
                "list",
                [
                  "object",
                  {
                    "pod_cidr": "string",
                    "pod_cidrs": [
                      "list",
                      "string"
                    ],
                    "provider_id": "string",
                    "taints": [
                      "list",
                      [
                        "object",
                        {
                          "effect": "string",
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "unschedulable": "bool"
                  }
                ]
              ],
              "status": [
                "list",
                [
                  "object",
                  {
                    "addresses": [
                      "list",
                      [
                        "object",
                        {
                          "address": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "allocatable": [
                      "map",
                      "string"
                    ],
                    "capacity": [
                      "map",
                      "string"
                    ],
                    "conditions": [
                      "list",
                      [
                        "object",
                        {
                          "last_heartbeat_time": "string",
                          "last_transition_time": "string",
                          "message": "string",
                          "reason": "string",
                          "status": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "node_info": [
                      "list",
                      [
                        "object",
                        {
                          "architecture": "string",
                          "boot_id": "string",
                          "container_runtime_version": "string",
                          "kernel_version": "string",
                          "kube_proxy_version": "string",
                          "kubelet_version": "string",
                          "machine_id": "string",
                          "operating_system": "string",
                          "os_image": "string",
                          "system_uuid": "string"
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
            "labels": {
              "description": "Select nodes with these labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description": "Metadata fields to narrow node selection.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "This data source provides a mechanism for listing the names of nodes in a kubernetes cluster.By default, all nodes in the cluster are returned, but queries by node label are also supported. It can be used to check for the existence of a specific node or to lookup a node to apply a taint with the ` + "`" + `kubernetes_node_taint` + "`" + ` resource.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesNodesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesNodes), &result)
	return &result
}
