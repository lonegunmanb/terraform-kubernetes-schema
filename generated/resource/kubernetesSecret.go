package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesSecret = `{
  "block": {
    "attributes": {
      "binary_data": {
        "description": "A map of the secret data in base64 encoding. Use this for binary data.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "binary_data_wo": {
        "description": "A write-only map of the secret data in base64 encoding. Use this for binary data.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ],
        "write_only": true
      },
      "binary_data_wo_revision": {
        "description": "The current revision of the write-only \"binary_data_wo\" attribute. Incrementing this integer value will cause Terraform to update the write-only value.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data": {
        "computed": true,
        "description": "A map of the secret data.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "data_wo": {
        "description": "A map write-only of the secret data.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ],
        "write_only": true
      },
      "data_wo_revision": {
        "description": "The current revision of the write-only \"data_wo\" attribute. Incrementing this integer value will cause Terraform to update the write-only value.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "immutable": {
        "description": "Ensures that data stored in the Secret cannot be updated (only object metadata can be modified).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "type": {
        "description": "Type of secret",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wait_for_service_account_token": {
        "description": "Terraform will wait for the service account token to be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "An unstructured key value map stored with the secret that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "generate_name": {
              "description": "Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "generation": {
              "computed": true,
              "description": "A sequence number representing a specific generation of the desired state.",
              "description_kind": "plain",
              "type": "number"
            },
            "labels": {
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the secret. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the secret, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description": "Namespace defines the space within which name of the secret must be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this secret that can be used by clients to determine when secret has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this secret. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard secret's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description": "The resource provides mechanisms to inject containers with sensitive information, such as passwords, while keeping containers agnostic of Kubernetes. Secrets can be used to store sensitive information either as individual properties or coarse-grained entries like entire files or JSON blobs. The resource will by default create a secret which is available to any pod in the specified (or default) namespace.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesSecret), &result)
	return &result
}
