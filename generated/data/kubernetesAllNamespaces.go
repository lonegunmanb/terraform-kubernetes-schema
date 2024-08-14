package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesAllNamespaces = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespaces": {
        "computed": true,
        "description": "List of all namespaces in a cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "This data source provides a mechanism for listing the names of all available namespaces in a Kubernetes cluster. It can be used to check for existence of a specific namespaces or to apply another resource to all or a subset of existing namespaces in a cluster.In Kubernetes, namespaces provide a scope for names and are intended as a way to divide cluster resources between multiple users.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesAllNamespacesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesAllNamespaces), &result)
	return &result
}
