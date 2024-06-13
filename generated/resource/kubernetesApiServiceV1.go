package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesApiServiceV1 = `{
  "block": {
    "attributes": {
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
            "annotations": {
              "description": "An unstructured key value map stored with the api_service that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the api_service. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the api_service, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this api_service that can be used by clients to determine when api_service has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this api_service. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard api_service's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "ca_bundle": {
              "description": "CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group": {
              "description": "Group is the API group name this server hosts.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "group_priority_minimum": {
              "description": "GroupPriorityMinimum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object. (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "insecure_skip_tls_verify": {
              "description": "InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged. You should use the CABundle instead.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "version": {
              "description": "Version is the API version this server hosts. For example, ` + "`" + `v1` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version_priority": {
              "description": "VersionPriority controls the ordering of this API version inside of its group. Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is ` + "`" + `kube-like` + "`" + `, it will sort above non ` + "`" + `kube-like` + "`" + ` version strings, which are ordered lexicographically. ` + "`" + `Kube-like` + "`" + ` versions start with a ` + "`" + `v` + "`" + `, then are followed by a number (the major version), then optionally the string ` + "`" + `alpha` + "`" + ` or ` + "`" + `beta` + "`" + ` and another number (the minor version). These are sorted first by GA \u003e ` + "`" + `beta` + "`" + ` \u003e ` + "`" + `alpha` + "`" + ` (where GA is a version with no suffix such as ` + "`" + `beta` + "`" + ` or ` + "`" + `alpha` + "`" + `), and then by comparing major version, then minor version. An example sorted list of versions: ` + "`" + `v10` + "`" + `, ` + "`" + `v2` + "`" + `, ` + "`" + `v1` + "`" + `, ` + "`" + `v11beta2` + "`" + `, ` + "`" + `v10beta3` + "`" + `, ` + "`" + `v3beta1` + "`" + `, ` + "`" + `v12alpha1` + "`" + `, ` + "`" + `v11alpha2` + "`" + `, ` + "`" + `foo1` + "`" + `, ` + "`" + `foo10` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "service": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name is the name of the service.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description": "Namespace is the namespace of the service.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description": "If specified, the port on the service that is hosting the service. Defaults to 443 for backward compatibility. Should be a valid port number (1-65535, inclusive).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Service is a reference to the service for this API server. It must communicate on port 443. If the Service is nil, that means the handling for the API groupversion is handled locally on this server. The call will simply delegate to the normal handler chain to be fulfilled.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Spec contains information for locating and communicating with a server. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "An API Service is an abstraction which defines for locating and communicating with servers.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesApiServiceV1Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesApiServiceV1), &result)
	return &result
}
