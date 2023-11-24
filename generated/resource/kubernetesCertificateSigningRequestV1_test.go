package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-kubernetes-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestKubernetesCertificateSigningRequestV1Schema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.KubernetesCertificateSigningRequestV1Schema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
