package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesResources = `{
  "block": {
    "attributes": {
      "api_version": {
        "description": "The resource apiVersion.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "field_selector": {
        "description": "A selector to restrict the list of returned objects by their fields.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kind": {
        "description": "The resource kind.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "label_selector": {
        "description": "A selector to restrict the list of returned objects by their labels.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "limit": {
        "description": "Limit is a maximum number of responses to return for a list call.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "namespace": {
        "description": "The resource namespace.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "objects": {
        "computed": true,
        "description": "The response from the API server.",
        "description_kind": "plain",
        "optional": true,
        "type": "dynamic"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func KubernetesResourcesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesResources), &result)
	return &result
}
