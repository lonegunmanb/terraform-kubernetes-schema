package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesSecretV1Data = `{
  "block": {
    "attributes": {
      "data": {
        "description": "Data to be stored in the Kubernetes Secret.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "field_manager": {
        "description": "Set the name of the field manager for the specified labels",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force": {
        "description": "Flag to force updates to the Kubernetes Secret.",
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
              "description": "The name of the Secret.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "namespace": {
              "description": "The namespace of the Secret.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Metadata for the kubernetes Secret.",
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

func KubernetesSecretV1DataSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesSecretV1Data), &result)
	return &result
}
