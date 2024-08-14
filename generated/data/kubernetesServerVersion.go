package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesServerVersion = `{
  "block": {
    "attributes": {
      "build_date": {
        "computed": true,
        "description": "Kubernetes server build date",
        "description_kind": "plain",
        "type": "string"
      },
      "compiler": {
        "computed": true,
        "description": "Compiler used to build Kubernetes",
        "description_kind": "plain",
        "type": "string"
      },
      "git_commit": {
        "computed": true,
        "description": "Git commit SHA",
        "description_kind": "plain",
        "type": "string"
      },
      "git_tree_state": {
        "computed": true,
        "description": "Git commit tree state",
        "description_kind": "plain",
        "type": "string"
      },
      "git_version": {
        "computed": true,
        "description": "Composite version and git commit sha",
        "description_kind": "plain",
        "type": "string"
      },
      "go_version": {
        "computed": true,
        "description": "Go compiler version",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "major": {
        "computed": true,
        "description": "Major Kubernetes version",
        "description_kind": "plain",
        "type": "string"
      },
      "minor": {
        "computed": true,
        "description": "Minor Kubernetes version",
        "description_kind": "plain",
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description": "Platform",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "Composite Kubernetes server version",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "This data source reads the versioning information of the server and makes specific attributes available to Terraform. Read more at [version info reference](https://pkg.go.dev/k8s.io/apimachinery/pkg/version#Info)",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesServerVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesServerVersion), &result)
	return &result
}
