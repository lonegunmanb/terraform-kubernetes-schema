package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesManifest = `{
  "block": {
    "attributes": {
      "computed_fields": {
        "description": "List of manifest fields whose values can be altered by the API server during 'apply'. Defaults to: [\"metadata.annotations\", \"metadata.labels\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "manifest": {
        "description": "A Kubernetes manifest describing the desired state of the resource in HCL format.",
        "description_kind": "plain",
        "required": true,
        "type": "dynamic"
      },
      "object": {
        "computed": true,
        "description": "The resulting resource state, as returned by the API server after applying the desired state from ` + "`" + `manifest` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "dynamic"
      },
      "wait_for": {
        "deprecated": true,
        "description": "A map of attribute paths and desired patterns to be matched. After each apply the provider will wait for all attributes listed here to reach a value that matches the desired pattern.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "object",
          {
            "fields": [
              "map",
              "string"
            ]
          }
        ]
      }
    },
    "block_types": {
      "field_manager": {
        "block": {
          "attributes": {
            "force_conflicts": {
              "description": "Force changes against conflicts.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description": "The name to use for the field manager when creating and updating the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configure field manager options.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "Timeout for the create operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "Timeout for the delete operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "Timeout for the update operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "wait": {
        "block": {
          "attributes": {
            "fields": {
              "description": "A map of paths to fields to wait for a specific field value.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "rollout": {
              "description": "Wait for rollout to complete on resources that support ` + "`" + `kubectl rollout status` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "condition": {
              "block": {
                "attributes": {
                  "status": {
                    "description": "The condition status.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "The type of condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Configure waiter options.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func KubernetesManifestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesManifest), &result)
	return &result
}
