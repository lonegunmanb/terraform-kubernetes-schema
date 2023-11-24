package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesPersistentVolume = `{
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
              "description": "An unstructured key value map stored with the persistent volume that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the persistent volume. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the persistent volume, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this persistent volume that can be used by clients to determine when persistent volume has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this persistent volume. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard persistent volume's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "access_modes": {
              "description": "Contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "capacity": {
              "description": "A description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity",
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            },
            "mount_options": {
              "description": "A list of mount options, e.g. [\"ro\", \"soft\"]. Not validated - mount will simply fail if one is invalid.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "persistent_volume_reclaim_policy": {
              "description": "What happens to a persistent volume when released from its claim. Valid options are Retain (default) and Recycle. Recycling must be supported by the volume plugin underlying this persistent volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_class_name": {
              "description": "A description of the persistent volume's class. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#class",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_mode": {
              "description": "Defines if a volume is intended to be used with a formatted filesystem. or to remain in raw block state.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "claim_ref": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The name of the PersistentVolumeClaim",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description": "The namespace of the PersistentVolumeClaim. Uses 'default' namespace if none is specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A reference to the persistent volume claim details for statically managed PVs. More Info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#binding",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "node_affinity": {
              "block": {
                "block_types": {
                  "required": {
                    "block": {
                      "block_types": {
                        "node_selector_term": {
                          "block": {
                            "block_types": {
                              "match_expressions": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description": "The label key that the selector applies to.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "operator": {
                                      "description": "A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + `, ` + "`" + `DoesNotExist` + "`" + `, ` + "`" + `Gt` + "`" + `, and ` + "`" + `Lt` + "`" + `.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "values": {
                                      "description": "An array of string values. If the operator is ` + "`" + `In` + "`" + ` or ` + "`" + `NotIn` + "`" + `, the values array must be non-empty. If the operator is ` + "`" + `Exists` + "`" + ` or ` + "`" + `DoesNotExist` + "`" + `, the values array must be empty. This array is replaced during a strategic merge patch.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "A list of node selector requirements by node's labels. The requirements are ANDed.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "match_fields": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description": "The label key that the selector applies to.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "operator": {
                                      "description": "A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + `, ` + "`" + `DoesNotExist` + "`" + `, ` + "`" + `Gt` + "`" + `, and ` + "`" + `Lt` + "`" + `.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "values": {
                                      "description": "An array of string values. If the operator is ` + "`" + `In` + "`" + ` or ` + "`" + `NotIn` + "`" + `, the values array must be non-empty. If the operator is ` + "`" + `Exists` + "`" + ` or ` + "`" + `DoesNotExist` + "`" + `, the values array must be empty. This array is replaced during a strategic merge patch.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "A list of node selector requirements by node's fields. The requirements are ANDed.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A description of the persistent volume's node affinity. More info: https://kubernetes.io/docs/concepts/storage/volumes/#local",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "persistent_volume_source": {
              "block": {
                "block_types": {
                  "aws_elastic_block_store": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "partition": {
                          "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "read_only": {
                          "description": "Whether to set the read-only property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "volume_id": {
                          "description": "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "azure_disk": {
                    "block": {
                      "attributes": {
                        "caching_mode": {
                          "description": "Host Caching mode: None, Read Only, Read Write.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "data_disk_uri": {
                          "description": "The URI the data disk in the blob storage",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "disk_name": {
                          "description": "The Name of the data disk in the blob storage",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kind": {
                          "computed": true,
                          "description": "The type for the data disk. Expected values: Shared, Dedicated, Managed. Defaults to Shared",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Represents an Azure Data Disk mount on the host and bind mount to the pod.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "azure_file": {
                    "block": {
                      "attributes": {
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "secret_name": {
                          "description": "The name of secret that contains Azure Storage Account Name and Key",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_namespace": {
                          "description": "The namespace of the secret that contains Azure Storage Account Name and Key. For Kubernetes up to 1.18.x the default is the same as the Pod. For Kubernetes 1.19.x and later the default is \"default\" namespace.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "share_name": {
                          "description": "Share Name",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an Azure File Service mount on the host and bind mount to the pod.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ceph_fs": {
                    "block": {
                      "attributes": {
                        "monitors": {
                          "description": "Monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "path": {
                          "description": "Used as the mounted root, rather than the full Ceph tree, default is /",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to ` + "`" + `false` + "`" + ` (read/write). More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "secret_file": {
                          "description": "The path to key ring for User, default is ` + "`" + `/etc/ceph/user.secret` + "`" + `. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user": {
                          "description": "User is the rados user name, default is admin. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents a Ceph FS mount on the host that shares a pod's lifetime",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "cinder": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "volume_id": {
                          "description": "Volume ID used to identify the volume in Cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "csi": {
                    "block": {
                      "attributes": {
                        "driver": {
                          "description": "the name of the volume driver to use. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to set the read-only property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: https://kubernetes.io/docs/concepts/storage/volumes#csi",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "volume_attributes": {
                          "description": "Attributes of the volume to publish.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "volume_handle": {
                          "description": "A string value that uniquely identifies the volume. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "controller_expand_secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerExpandVolume call.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "controller_publish_secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "node_publish_secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "node_stage_secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents a CSI Volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#csi",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "fc": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lun": {
                          "description": "FC target lun number",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "target_ww_ns": {
                          "description": "FC target worldwide names (WWNs)",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description": "Represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "flex_volume": {
                    "block": {
                      "attributes": {
                        "driver": {
                          "description": "Driver is the name of the driver to use for this volume.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". The default filesystem depends on FlexVolume script.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "options": {
                          "description": "Extra command options if any.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "read_only": {
                          "description": "Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents a generic volume resource that is provisioned/attached using an exec based plugin. This is an alpha feature and may change in future.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "flocker": {
                    "block": {
                      "attributes": {
                        "dataset_name": {
                          "description": "Name of the dataset stored as metadata -\u003e name on the dataset for Flocker should be considered as deprecated",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "dataset_uuid": {
                          "description": "UUID of the dataset. This is unique identifier of a Flocker dataset",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "gce_persistent_disk": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "partition": {
                          "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "pd_name": {
                          "description": "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "glusterfs": {
                    "block": {
                      "attributes": {
                        "endpoints_name": {
                          "description": "The endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description": "The Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "host_path": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Type for HostPath volume. Allowed values are \"\" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "iscsi": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "iqn": {
                          "description": "Target iSCSI Qualified Name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "iscsi_interface": {
                          "description": "iSCSI interface name that uses an iSCSI transport. Defaults to 'default' (tcp).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lun": {
                          "description": "iSCSI target lun number.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "target_portal": {
                          "description": "iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "local": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#local",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a mounted local storage device such as a disk, partition or directory. Local volumes can only be used as a statically created PersistentVolume. Dynamic provisioning is not supported yet. More info: https://kubernetes.io/docs/concepts/storage/volumes#local",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "nfs": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "server": {
                          "description": "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "photon_persistent_disk": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "pd_id": {
                          "description": "ID that identifies Photon Controller persistent disk",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a PhotonController persistent disk attached and mounted on kubelets host machine",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "quobyte": {
                    "block": {
                      "attributes": {
                        "group": {
                          "description": "Group to map volume access to Default is no group",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "registry": {
                          "description": "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "user": {
                          "description": "User to map volume access to Defaults to serivceaccount user",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "volume": {
                          "description": "Volume is a string that references an already created Quobyte volume by name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "rbd": {
                    "block": {
                      "attributes": {
                        "ceph_monitors": {
                          "description": "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "fs_type": {
                          "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "keyring": {
                          "computed": true,
                          "description": "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "rados_user": {
                          "description": "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "rbd_image": {
                          "description": "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "rbd_pool": {
                          "description": "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "read_only": {
                          "description": "Whether to force the read-only setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "secret_ref": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "vsphere_volume": {
                    "block": {
                      "attributes": {
                        "fs_type": {
                          "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "volume_path": {
                          "description": "Path that identifies vSphere volume vmdk",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a vSphere volume attached and mounted on kubelets host machine",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The specification of a persistent volume.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Spec of the persistent volume owned by the cluster",
          "description_kind": "plain"
        },
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

func KubernetesPersistentVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesPersistentVolume), &result)
	return &result
}
