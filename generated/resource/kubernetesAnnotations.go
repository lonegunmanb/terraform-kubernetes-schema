package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesAnnotations = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "A map of annotations to apply to the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "api_version": {
        "description": "The apiVersion of the resource to annotate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "field_manager": {
        "description": "Set the name of the field manager for the specified labels.",
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
      },
      "kind": {
        "description": "The kind of the resource to annotate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_annotations": {
        "description": "A map of annotations to apply to the resource template.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "name": {
              "description": "The name of the resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "namespace": {
              "description": "The namespace of the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
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

func KubernetesAnnotationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesAnnotations), &result)
	return &result
}
