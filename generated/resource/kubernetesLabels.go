package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesLabels = `{
  "block": {
    "attributes": {
      "api_version": {
        "description": "The apiVersion of the resource to label.",
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
        "description": "Force overwriting labels that were created or edited outside of Terraform.",
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
        "description": "The kind of the resource to label.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "description": "A map of labels to apply to the resource.",
        "description_kind": "plain",
        "required": true,
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
    "description": "This resource allows Terraform to manage the labels for a resource that already exists. This resource uses [field management](https://kubernetes.io/docs/reference/using-api/server-side-apply/#field-management) and [server-side apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/) to manage only the labels that are defined in the Terraform configuration. Existing labels not specified in the configuration will be ignored. If a label specified in the config and is already managed by another client it will cause a conflict which can be overridden by setting ` + "`" + `force` + "`" + ` to true.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesLabelsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesLabels), &result)
	return &result
}
