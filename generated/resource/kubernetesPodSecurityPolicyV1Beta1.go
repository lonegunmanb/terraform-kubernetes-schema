package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesPodSecurityPolicyV1Beta1 = `{
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
              "description": "An unstructured key value map stored with the podsecuritypolicy that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "generation": {
              "computed": true,
              "description": "A sequence number representing a specific generation of the desired state.",
              "description_kind": "plain",
              "type": "number"
            },
            "labels": {
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the podsecuritypolicy. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the podsecuritypolicy, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this podsecuritypolicy that can be used by clients to determine when podsecuritypolicy has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this podsecuritypolicy. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard podsecuritypolicy's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "allow_privilege_escalation": {
              "computed": true,
              "description": "allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allowed_capabilities": {
              "computed": true,
              "description": "allowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allowed_proc_mount_types": {
              "description": "AllowedProcMountTypes is an allowlist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allowed_unsafe_sysctls": {
              "description": "allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in \"*\" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to allowlist all allowed unsafe sysctls explicitly to avoid rejection.\n\nExamples: e.g. \"foo/*\" allows \"foo/bar\", \"foo/baz\", etc. e.g. \"foo.*\" allows \"foo.bar\", \"foo.baz\", etc.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "default_add_capabilities": {
              "description": "defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "default_allow_privilege_escalation": {
              "computed": true,
              "description": "defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "forbidden_sysctls": {
              "description": "forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in \"*\" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.\n\nExamples: e.g. \"foo/*\" forbids \"foo/bar\", \"foo/baz\", etc. e.g. \"foo.*\" forbids \"foo.bar\", \"foo.baz\", etc.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "host_ipc": {
              "computed": true,
              "description": "hostIPC determines if the policy allows the use of HostIPC in the pod spec.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "host_network": {
              "computed": true,
              "description": "hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "host_pid": {
              "computed": true,
              "description": "hostPID determines if the policy allows the use of HostPID in the pod spec.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "privileged": {
              "computed": true,
              "description": "privileged determines if a pod can request to be run as privileged.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "read_only_root_filesystem": {
              "computed": true,
              "description": "readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "required_drop_capabilities": {
              "computed": true,
              "description": "requiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "volumes": {
              "computed": true,
              "description": "volumes is an allowlist of volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use '*'.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "allowed_flex_volumes": {
              "block": {
                "attributes": {
                  "driver": {
                    "description": "driver is the name of the Flexvolume driver.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "allowedFlexVolumes is an allowlist of Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the \"volumes\" field.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "allowed_host_paths": {
              "block": {
                "attributes": {
                  "path_prefix": {
                    "description": "pathPrefix is the path prefix that the host volume must match. It does not support ` + "`" + `*` + "`" + `. Trailing slashes are trimmed when validating the path prefix with a host path.\n\nExamples: ` + "`" + `/foo` + "`" + ` would allow ` + "`" + `/foo` + "`" + `, ` + "`" + `/foo/` + "`" + ` and ` + "`" + `/foo/bar` + "`" + ` ` + "`" + `/foo` + "`" + ` would not allow ` + "`" + `/food` + "`" + ` or ` + "`" + `/etc/foo` + "`" + `",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "read_only": {
                    "description": "when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "allowedHostPaths is an allowlist of host paths. Empty indicates that all host paths may be used.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "fs_group": {
              "block": {
                "attributes": {
                  "rule": {
                    "description": "rule is the strategy that will dictate what FSGroup is used in the SecurityContext.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "range": {
                    "block": {
                      "attributes": {
                        "max": {
                          "description": "max is the end of the range, inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "min": {
                          "description": "min is the start of the range, inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end. Required for MustRunAs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "host_ports": {
              "block": {
                "attributes": {
                  "max": {
                    "description": "max is the end of the range, inclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "min": {
                    "description": "min is the start of the range, inclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "hostPorts determines which host port ranges are allowed to be exposed.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "run_as_group": {
              "block": {
                "attributes": {
                  "rule": {
                    "description": "rule is the strategy that will dictate the allowable RunAsGroup values that may be set.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "range": {
                    "block": {
                      "attributes": {
                        "max": {
                          "description": "max is the end of the range, inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "min": {
                          "description": "min is the start of the range, inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "ranges are the allowed ranges of gids that may be used. If you would like to force a single gid then supply a single range with the same start and end. Required for MustRunAs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If this field is omitted, the pod's RunAsGroup can take any value. This field requires the RunAsGroup feature gate to be enabled.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "run_as_user": {
              "block": {
                "attributes": {
                  "rule": {
                    "description": "rule is the strategy that will dictate the allowable RunAsUser values that may be set.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "range": {
                    "block": {
                      "attributes": {
                        "max": {
                          "description": "max is the end of the range, inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "min": {
                          "description": "min is the start of the range, inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "ranges are the allowed ranges of uids that may be used. If you would like to force a single uid then supply a single range with the same start and end. Required for MustRunAs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "se_linux": {
              "block": {
                "attributes": {
                  "rule": {
                    "description": "rule is the strategy that will dictate the allowable labels that may be set.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "se_linux_options": {
                    "block": {
                      "attributes": {
                        "level": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "user": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "seLinux is the strategy that will dictate the allowable labels that may be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "supplemental_groups": {
              "block": {
                "attributes": {
                  "rule": {
                    "description": "rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "range": {
                    "block": {
                      "attributes": {
                        "max": {
                          "description": "max is the end of the range, inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "min": {
                          "description": "min is the start of the range, inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end. Required for MustRunAs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "supplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "spec defines the policy enforced.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "deprecated": true,
    "description": "A Pod Security Policy is a cluster-level resource that controls security sensitive aspects of the pod specification. The PodSecurityPolicy objects define a set of conditions that a pod must run with in order to be accepted into the system, as well as defaults for the related fields.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesPodSecurityPolicyV1Beta1Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesPodSecurityPolicyV1Beta1), &result)
	return &result
}
