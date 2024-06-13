package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesEnv = `{
  "block": {
    "attributes": {
      "api_version": {
        "description": "Resource API version",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "container": {
        "description": "Name of the container for which we are updating the environment variables.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "field_manager": {
        "description": "Set the name of the field manager for the specified environment variables.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force": {
        "description": "Force overwriting environments that were created or edited outside of Terraform.",
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
      "init_container": {
        "description": "Name of the initContainer for which we are updating the environment variables.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kind": {
        "description": "Resource Kind",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "env": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the environment variable. Must be a C_IDENTIFIER",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "value_from": {
              "block": {
                "block_types": {
                  "config_map_key_ref": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description": "The key to select.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "optional": {
                          "description": "Specify whether the ConfigMap or its key must be defined.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Selects a key of a ConfigMap.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "field_ref": {
                    "block": {
                      "attributes": {
                        "api_version": {
                          "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "field_path": {
                          "description": "Path of the field to select in the specified API version",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.podIP.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "resource_field_ref": {
                    "block": {
                      "attributes": {
                        "container_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "divisor": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource": {
                          "description": "Resource to select",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "secret_key_ref": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description": "The key of the secret to select from. Must be a valid secret key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "optional": {
                          "description": "Specify whether the Secret or its key must be defined.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Selects a key of a secret in the pod's namespace.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Source for the environment variable's value",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "List of custom values used to represent environment variables",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
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
    "description": "This resource provides a way to manage environment variables in resources that were created outside of Terraform. This resource provides functionality similar to the ` + "`" + `kubectl set env` + "`" + ` command.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesEnvSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesEnv), &result)
	return &result
}
