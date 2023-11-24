package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-kubernetes-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestKubernetesServiceAccountSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.KubernetesServiceAccountSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
