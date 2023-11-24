package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-kubernetes-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestKubernetesPodV1Schema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.KubernetesPodV1Schema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
