package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesMutatingWebhookConfigurationV1 = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "webhook": {
        "computed": true,
        "description": "Webhooks is a list of webhooks and the affected resources and operations.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "admission_review_versions": [
                "list",
                "string"
              ],
              "client_config": [
                "list",
                [
                  "object",
                  {
                    "ca_bundle": "string",
                    "service": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "namespace": "string",
                          "path": "string",
                          "port": "number"
                        }
                      ]
                    ],
                    "url": "string"
                  }
                ]
              ],
              "failure_policy": "string",
              "match_policy": "string",
              "name": "string",
              "namespace_selector": [
                "list",
                [
                  "object",
                  {
                    "match_expressions": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "operator": "string",
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "match_labels": [
                      "map",
                      "string"
                    ]
                  }
                ]
              ],
              "object_selector": [
                "list",
                [
                  "object",
                  {
                    "match_expressions": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "operator": "string",
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "match_labels": [
                      "map",
                      "string"
                    ]
                  }
                ]
              ],
              "reinvocation_policy": "string",
              "rule": [
                "list",
                [
                  "object",
                  {
                    "api_groups": [
                      "list",
                      "string"
                    ],
                    "api_versions": [
                      "list",
                      "string"
                    ],
                    "operations": [
                      "list",
                      "string"
                    ],
                    "resources": [
                      "list",
                      "string"
                    ],
                    "scope": "string"
                  }
                ]
              ],
              "side_effects": "string",
              "timeout_seconds": "number"
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
              "description": "An unstructured key value map stored with the mutating webhook configuration that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the mutating webhook configuration. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the mutating webhook configuration, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this mutating webhook configuration that can be used by clients to determine when mutating webhook configuration has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this mutating webhook configuration. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard mutating webhook configuration's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "A Mutating Webhook Configuration configures a [mutating admission webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#what-are-admission-webhooks). This data source allows you to pull data about a given mutating webhook configuration based on its name.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesMutatingWebhookConfigurationV1Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesMutatingWebhookConfigurationV1), &result)
	return &result
}
