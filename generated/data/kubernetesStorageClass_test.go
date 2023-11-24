package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-kubernetes-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestKubernetesStorageClassSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.KubernetesStorageClassSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
