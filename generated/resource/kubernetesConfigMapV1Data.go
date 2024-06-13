package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesConfigMapV1Data = `{
  "block": {
    "attributes": {
      "data": {
        "description": "The data we want to add to the ConfigMap.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "field_manager": {
        "description": "Set the name of the field manager for the specified labels.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force": {
        "description": "Force overwriting data that is managed outside of Terraform.",
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
              "description": "The name of the ConfigMap.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "namespace": {
              "description": "The namespace of the ConfigMap.",
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
    "description": "This resource allows Terraform to manage data within a pre-existing ConfigMap. This resource uses [field management](https://kubernetes.io/docs/reference/using-api/server-side-apply/#field-management) and [server-side apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/) to manage only the data that is defined in the Terraform configuration. Existing data not specified in the configuration will be ignored. If data specified in the config and is already managed by another client it will cause a conflict which can be overridden by setting ` + "`" + `force` + "`" + ` to true.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesConfigMapV1DataSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesConfigMapV1Data), &result)
	return &result
}
