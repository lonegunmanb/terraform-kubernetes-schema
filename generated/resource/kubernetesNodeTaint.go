package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesNodeTaint = `{
  "block": {
    "attributes": {
      "field_manager": {
        "description": "Set the name of the field manager for the node taint",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force": {
        "description": "Force overwriting annotations that were created or edited outside of Terraform.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
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
            "name": {
              "description": "The name of the node",
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
      "taint": {
        "block": {
          "attributes": {
            "effect": {
              "description": "The taint effect",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key": {
              "description": "The taint key",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The taint value",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "[Node affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity) is a property of Pods that attracts them to a set of [nodes](https://kubernetes.io/docs/concepts/architecture/nodes/) (either as a preference or a hard requirement). Taints are the opposite -- they allow a node to repel a set of pods.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesNodeTaintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesNodeTaint), &result)
	return &result
}
