package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesResource = `{
  "block": {
    "attributes": {
      "api_version": {
        "description": "The resource apiVersion.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kind": {
        "description": "The resource kind.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "object": {
        "computed": true,
        "description": "The response from the API server.",
        "description_kind": "plain",
        "optional": true,
        "type": "dynamic"
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "name": {
              "description": "The resource name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "namespace": {
              "description": "The resource namespace.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Metadata for the resource",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func KubernetesResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesResource), &result)
	return &result
}
