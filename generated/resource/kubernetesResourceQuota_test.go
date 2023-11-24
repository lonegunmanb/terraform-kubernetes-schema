package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-kubernetes-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestKubernetesResourceQuotaSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.KubernetesResourceQuotaSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
