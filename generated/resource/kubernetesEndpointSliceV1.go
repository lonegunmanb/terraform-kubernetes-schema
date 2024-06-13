package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesEndpointSliceV1 = `{
  "block": {
    "attributes": {
      "address_type": {
        "description": "address_type specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "endpoint": {
        "block": {
          "attributes": {
            "addresses": {
              "description": "addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "hostname": {
              "description": "hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_name": {
              "description": "nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zone": {
              "description": "zone is the name of the Zone this endpoint exists in.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "condition": {
              "block": {
                "attributes": {
                  "ready": {
                    "description": "ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "serving": {
                    "description": "serving is identical to ready except that it is set regardless of the terminating state of endpoints.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "terminating": {
                    "description": "terminating indicates that this endpoint is terminating.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "condition contains information about the current status of the endpoint.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "target_ref": {
              "block": {
                "attributes": {
                  "field_path": {
                    "description": "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Name of the referent.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description": "Namespace of the referent.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource_version": {
                    "description": "Specific resourceVersion to which this reference is made, if any.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uid": {
                    "description": "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "targetRef is a reference to a Kubernetes object that represents this endpoint.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "endpoint is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.",
          "description_kind": "plain"
        },
        "max_items": 1000,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "An unstructured key value map stored with the endpoint_slice that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the endpoint_slice. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the endpoint_slice, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description": "Namespace defines the space within which name of the endpoint_slice must be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this endpoint_slice that can be used by clients to determine when endpoint_slice has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this endpoint_slice. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard endpoint_slice's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "port": {
        "block": {
          "attributes": {
            "app_protocol": {
              "description": "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "name represents the name of this port. All ports in an EndpointSlice must have a unique name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "port represents the port number of the endpoint.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol": {
              "description": "protocol represents the IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "port specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. Each slice may include a maximum of 100 ports.",
          "description_kind": "plain"
        },
        "max_items": 100,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "An EndpointSlice contains references to a set of network endpoints.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesEndpointSliceV1Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesEndpointSliceV1), &result)
	return &result
}
