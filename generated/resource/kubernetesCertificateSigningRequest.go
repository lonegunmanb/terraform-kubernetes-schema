package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesCertificateSigningRequest = `{
  "block": {
    "attributes": {
      "auto_approve": {
        "description": "Automatically approve the CertificateSigningRequest",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "certificate": {
        "computed": true,
        "description": "If request was approved, the controller will place the issued certificate here.",
        "description_kind": "plain",
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
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "An unstructured key value map stored with the certificate signing request that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the certificate signing request. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the certificate signing request, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this certificate signing request that can be used by clients to determine when certificate signing request has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this certificate signing request. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard certificate signing request's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "request": {
              "description": "Base64-encoded PKCS#10 CSR data",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "signer_name": {
              "description": "Requested signer for the request. It is a qualified name in the form: ` + "`" + `scope-hostname.io/name` + "`" + `.If empty, it will be defaulted: 1. If it's a kubelet client certificate, it is assigned ` + "`" + `kubernetes.io/kube-apiserver-client-kubelet` + "`" + `.2. If it's a kubelet serving certificate, it is assigned ` + "`" + `kubernetes.io/kubelet-serving` + "`" + `.3. Otherwise, it is assigned ` + "`" + `kubernetes.io/legacy-unknown` + "`" + `. Distribution of trust for signers happens out of band.You can select on this field using ` + "`" + `spec.signerName` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "usages": {
              "description": "allowedUsages specifies a set of usage contexts the key will be valid for. See:\n\thttps://tools.ietf.org/html/rfc5280#section-4.2.1.3\n\thttps://tools.ietf.org/html/rfc5280#section-4.2.1.12\n\nValid values are:\n \"signing\",\n \"digital signature\",\n \"content commitment\",\n \"key encipherment\",\n \"key agreement\",\n \"data encipherment\",\n \"cert sign\",\n \"crl sign\",\n \"encipher only\",\n \"decipher only\",\n \"any\",\n \"server auth\",\n \"client auth\",\n \"code signing\",\n \"email protection\",\n \"s/mime\",\n \"ipsec end system\",\n \"ipsec tunnel\",\n \"ipsec user\",\n \"timestamping\",\n \"ocsp signing\",\n \"microsoft sgc\",\n \"netscape sgc\"",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "Describes a certificate signing request",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesCertificateSigningRequestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesCertificateSigningRequest), &result)
	return &result
}
