package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const kubernetesPod = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spec": {
        "computed": true,
        "description": "Specification of the desired behavior of the pod.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active_deadline_seconds": "number",
              "affinity": [
                "list",
                [
                  "object",
                  {
                    "node_affinity": [
                      "list",
                      [
                        "object",
                        {
                          "preferred_during_scheduling_ignored_during_execution": [
                            "list",
                            [
                              "object",
                              {
                                "preference": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "match_expressions": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "operator": "string",
                                            "values": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ],
                                      "match_fields": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "operator": "string",
                                            "values": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "weight": "number"
                              }
                            ]
                          ],
                          "required_during_scheduling_ignored_during_execution": [
                            "list",
                            [
                              "object",
                              {
                                "node_selector_term": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "match_expressions": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "operator": "string",
                                            "values": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ],
                                      "match_fields": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "operator": "string",
                                            "values": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "pod_affinity": [
                      "list",
                      [
                        "object",
                        {
                          "preferred_during_scheduling_ignored_during_execution": [
                            "list",
                            [
                              "object",
                              {
                                "pod_affinity_term": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "label_selector": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "match_expressions": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "key": "string",
                                                  "operator": "string",
                                                  "values": [
                                                    "set",
                                                    "string"
                                                  ]
                                                }
                                              ]
                                            ],
                                            "match_labels": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ],
                                      "namespaces": [
                                        "set",
                                        "string"
                                      ],
                                      "topology_key": "string"
                                    }
                                  ]
                                ],
                                "weight": "number"
                              }
                            ]
                          ],
                          "required_during_scheduling_ignored_during_execution": [
                            "list",
                            [
                              "object",
                              {
                                "label_selector": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "match_expressions": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "operator": "string",
                                            "values": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ],
                                      "match_labels": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "namespaces": [
                                  "set",
                                  "string"
                                ],
                                "topology_key": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "pod_anti_affinity": [
                      "list",
                      [
                        "object",
                        {
                          "preferred_during_scheduling_ignored_during_execution": [
                            "list",
                            [
                              "object",
                              {
                                "pod_affinity_term": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "label_selector": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "match_expressions": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "key": "string",
                                                  "operator": "string",
                                                  "values": [
                                                    "set",
                                                    "string"
                                                  ]
                                                }
                                              ]
                                            ],
                                            "match_labels": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ],
                                      "namespaces": [
                                        "set",
                                        "string"
                                      ],
                                      "topology_key": "string"
                                    }
                                  ]
                                ],
                                "weight": "number"
                              }
                            ]
                          ],
                          "required_during_scheduling_ignored_during_execution": [
                            "list",
                            [
                              "object",
                              {
                                "label_selector": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "match_expressions": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "operator": "string",
                                            "values": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ],
                                      "match_labels": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "namespaces": [
                                  "set",
                                  "string"
                                ],
                                "topology_key": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "automount_service_account_token": "bool",
              "container": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "command": [
                      "list",
                      "string"
                    ],
                    "env": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "value": "string",
                          "value_from": [
                            "list",
                            [
                              "object",
                              {
                                "config_map_key_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "key": "string",
                                      "name": "string",
                                      "optional": "bool"
                                    }
                                  ]
                                ],
                                "field_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "api_version": "string",
                                      "field_path": "string"
                                    }
                                  ]
                                ],
                                "resource_field_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "container_name": "string",
                                      "divisor": "string",
                                      "resource": "string"
                                    }
                                  ]
                                ],
                                "secret_key_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "key": "string",
                                      "name": "string",
                                      "optional": "bool"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "env_from": [
                      "list",
                      [
                        "object",
                        {
                          "config_map_ref": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "optional": "bool"
                              }
                            ]
                          ],
                          "prefix": "string",
                          "secret_ref": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "optional": "bool"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "image": "string",
                    "image_pull_policy": "string",
                    "lifecycle": [
                      "list",
                      [
                        "object",
                        {
                          "post_start": [
                            "list",
                            [
                              "object",
                              {
                                "exec": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "command": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "http_get": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "host": "string",
                                      "http_header": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "name": "string",
                                            "value": "string"
                                          }
                                        ]
                                      ],
                                      "path": "string",
                                      "port": "string",
                                      "scheme": "string"
                                    }
                                  ]
                                ],
                                "tcp_socket": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "port": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "pre_stop": [
                            "list",
                            [
                              "object",
                              {
                                "exec": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "command": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "http_get": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "host": "string",
                                      "http_header": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "name": "string",
                                            "value": "string"
                                          }
                                        ]
                                      ],
                                      "path": "string",
                                      "port": "string",
                                      "scheme": "string"
                                    }
                                  ]
                                ],
                                "tcp_socket": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "port": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "liveness_probe": [
                      "list",
                      [
                        "object",
                        {
                          "exec": [
                            "list",
                            [
                              "object",
                              {
                                "command": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "failure_threshold": "number",
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "service": "string"
                              }
                            ]
                          ],
                          "http_get": [
                            "list",
                            [
                              "object",
                              {
                                "host": "string",
                                "http_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "path": "string",
                                "port": "string",
                                "scheme": "string"
                              }
                            ]
                          ],
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "success_threshold": "number",
                          "tcp_socket": [
                            "list",
                            [
                              "object",
                              {
                                "port": "string"
                              }
                            ]
                          ],
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "name": "string",
                    "port": [
                      "list",
                      [
                        "object",
                        {
                          "container_port": "number",
                          "host_ip": "string",
                          "host_port": "number",
                          "name": "string",
                          "protocol": "string"
                        }
                      ]
                    ],
                    "readiness_probe": [
                      "list",
                      [
                        "object",
                        {
                          "exec": [
                            "list",
                            [
                              "object",
                              {
                                "command": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "failure_threshold": "number",
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "service": "string"
                              }
                            ]
                          ],
                          "http_get": [
                            "list",
                            [
                              "object",
                              {
                                "host": "string",
                                "http_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "path": "string",
                                "port": "string",
                                "scheme": "string"
                              }
                            ]
                          ],
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "success_threshold": "number",
                          "tcp_socket": [
                            "list",
                            [
                              "object",
                              {
                                "port": "string"
                              }
                            ]
                          ],
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "resources": [
                      "list",
                      [
                        "object",
                        {
                          "limits": [
                            "map",
                            "string"
                          ],
                          "requests": [
                            "map",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "security_context": [
                      "list",
                      [
                        "object",
                        {
                          "allow_privilege_escalation": "bool",
                          "capabilities": [
                            "list",
                            [
                              "object",
                              {
                                "add": [
                                  "list",
                                  "string"
                                ],
                                "drop": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "privileged": "bool",
                          "read_only_root_filesystem": "bool",
                          "run_as_group": "string",
                          "run_as_non_root": "bool",
                          "run_as_user": "string",
                          "se_linux_options": [
                            "list",
                            [
                              "object",
                              {
                                "level": "string",
                                "role": "string",
                                "type": "string",
                                "user": "string"
                              }
                            ]
                          ],
                          "seccomp_profile": [
                            "list",
                            [
                              "object",
                              {
                                "localhost_profile": "string",
                                "type": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "startup_probe": [
                      "list",
                      [
                        "object",
                        {
                          "exec": [
                            "list",
                            [
                              "object",
                              {
                                "command": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "failure_threshold": "number",
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "service": "string"
                              }
                            ]
                          ],
                          "http_get": [
                            "list",
                            [
                              "object",
                              {
                                "host": "string",
                                "http_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "path": "string",
                                "port": "string",
                                "scheme": "string"
                              }
                            ]
                          ],
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "success_threshold": "number",
                          "tcp_socket": [
                            "list",
                            [
                              "object",
                              {
                                "port": "string"
                              }
                            ]
                          ],
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "stdin": "bool",
                    "stdin_once": "bool",
                    "termination_message_path": "string",
                    "termination_message_policy": "string",
                    "tty": "bool",
                    "volume_mount": [
                      "list",
                      [
                        "object",
                        {
                          "mount_path": "string",
                          "mount_propagation": "string",
                          "name": "string",
                          "read_only": "bool",
                          "sub_path": "string"
                        }
                      ]
                    ],
                    "working_dir": "string"
                  }
                ]
              ],
              "dns_config": [
                "list",
                [
                  "object",
                  {
                    "nameservers": [
                      "list",
                      "string"
                    ],
                    "option": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "searches": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "dns_policy": "string",
              "enable_service_links": "bool",
              "host_aliases": [
                "list",
                [
                  "object",
                  {
                    "hostnames": [
                      "list",
                      "string"
                    ],
                    "ip": "string"
                  }
                ]
              ],
              "host_ipc": "bool",
              "host_network": "bool",
              "host_pid": "bool",
              "hostname": "string",
              "image_pull_secrets": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "init_container": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "command": [
                      "list",
                      "string"
                    ],
                    "env": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "value": "string",
                          "value_from": [
                            "list",
                            [
                              "object",
                              {
                                "config_map_key_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "key": "string",
                                      "name": "string",
                                      "optional": "bool"
                                    }
                                  ]
                                ],
                                "field_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "api_version": "string",
                                      "field_path": "string"
                                    }
                                  ]
                                ],
                                "resource_field_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "container_name": "string",
                                      "divisor": "string",
                                      "resource": "string"
                                    }
                                  ]
                                ],
                                "secret_key_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "key": "string",
                                      "name": "string",
                                      "optional": "bool"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "env_from": [
                      "list",
                      [
                        "object",
                        {
                          "config_map_ref": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "optional": "bool"
                              }
                            ]
                          ],
                          "prefix": "string",
                          "secret_ref": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "optional": "bool"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "image": "string",
                    "image_pull_policy": "string",
                    "lifecycle": [
                      "list",
                      [
                        "object",
                        {
                          "post_start": [
                            "list",
                            [
                              "object",
                              {
                                "exec": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "command": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "http_get": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "host": "string",
                                      "http_header": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "name": "string",
                                            "value": "string"
                                          }
                                        ]
                                      ],
                                      "path": "string",
                                      "port": "string",
                                      "scheme": "string"
                                    }
                                  ]
                                ],
                                "tcp_socket": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "port": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "pre_stop": [
                            "list",
                            [
                              "object",
                              {
                                "exec": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "command": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "http_get": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "host": "string",
                                      "http_header": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "name": "string",
                                            "value": "string"
                                          }
                                        ]
                                      ],
                                      "path": "string",
                                      "port": "string",
                                      "scheme": "string"
                                    }
                                  ]
                                ],
                                "tcp_socket": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "port": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "liveness_probe": [
                      "list",
                      [
                        "object",
                        {
                          "exec": [
                            "list",
                            [
                              "object",
                              {
                                "command": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "failure_threshold": "number",
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "service": "string"
                              }
                            ]
                          ],
                          "http_get": [
                            "list",
                            [
                              "object",
                              {
                                "host": "string",
                                "http_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "path": "string",
                                "port": "string",
                                "scheme": "string"
                              }
                            ]
                          ],
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "success_threshold": "number",
                          "tcp_socket": [
                            "list",
                            [
                              "object",
                              {
                                "port": "string"
                              }
                            ]
                          ],
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "name": "string",
                    "port": [
                      "list",
                      [
                        "object",
                        {
                          "container_port": "number",
                          "host_ip": "string",
                          "host_port": "number",
                          "name": "string",
                          "protocol": "string"
                        }
                      ]
                    ],
                    "readiness_probe": [
                      "list",
                      [
                        "object",
                        {
                          "exec": [
                            "list",
                            [
                              "object",
                              {
                                "command": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "failure_threshold": "number",
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "service": "string"
                              }
                            ]
                          ],
                          "http_get": [
                            "list",
                            [
                              "object",
                              {
                                "host": "string",
                                "http_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "path": "string",
                                "port": "string",
                                "scheme": "string"
                              }
                            ]
                          ],
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "success_threshold": "number",
                          "tcp_socket": [
                            "list",
                            [
                              "object",
                              {
                                "port": "string"
                              }
                            ]
                          ],
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "resources": [
                      "list",
                      [
                        "object",
                        {
                          "limits": [
                            "map",
                            "string"
                          ],
                          "requests": [
                            "map",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "security_context": [
                      "list",
                      [
                        "object",
                        {
                          "allow_privilege_escalation": "bool",
                          "capabilities": [
                            "list",
                            [
                              "object",
                              {
                                "add": [
                                  "list",
                                  "string"
                                ],
                                "drop": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "privileged": "bool",
                          "read_only_root_filesystem": "bool",
                          "run_as_group": "string",
                          "run_as_non_root": "bool",
                          "run_as_user": "string",
                          "se_linux_options": [
                            "list",
                            [
                              "object",
                              {
                                "level": "string",
                                "role": "string",
                                "type": "string",
                                "user": "string"
                              }
                            ]
                          ],
                          "seccomp_profile": [
                            "list",
                            [
                              "object",
                              {
                                "localhost_profile": "string",
                                "type": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "startup_probe": [
                      "list",
                      [
                        "object",
                        {
                          "exec": [
                            "list",
                            [
                              "object",
                              {
                                "command": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "failure_threshold": "number",
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "service": "string"
                              }
                            ]
                          ],
                          "http_get": [
                            "list",
                            [
                              "object",
                              {
                                "host": "string",
                                "http_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "path": "string",
                                "port": "string",
                                "scheme": "string"
                              }
                            ]
                          ],
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "success_threshold": "number",
                          "tcp_socket": [
                            "list",
                            [
                              "object",
                              {
                                "port": "string"
                              }
                            ]
                          ],
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "stdin": "bool",
                    "stdin_once": "bool",
                    "termination_message_path": "string",
                    "termination_message_policy": "string",
                    "tty": "bool",
                    "volume_mount": [
                      "list",
                      [
                        "object",
                        {
                          "mount_path": "string",
                          "mount_propagation": "string",
                          "name": "string",
                          "read_only": "bool",
                          "sub_path": "string"
                        }
                      ]
                    ],
                    "working_dir": "string"
                  }
                ]
              ],
              "node_name": "string",
              "node_selector": [
                "map",
                "string"
              ],
              "os": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "priority_class_name": "string",
              "readiness_gate": [
                "list",
                [
                  "object",
                  {
                    "condition_type": "string"
                  }
                ]
              ],
              "restart_policy": "string",
              "runtime_class_name": "string",
              "scheduler_name": "string",
              "security_context": [
                "list",
                [
                  "object",
                  {
                    "fs_group": "string",
                    "fs_group_change_policy": "string",
                    "run_as_group": "string",
                    "run_as_non_root": "bool",
                    "run_as_user": "string",
                    "se_linux_options": [
                      "list",
                      [
                        "object",
                        {
                          "level": "string",
                          "role": "string",
                          "type": "string",
                          "user": "string"
                        }
                      ]
                    ],
                    "seccomp_profile": [
                      "list",
                      [
                        "object",
                        {
                          "localhost_profile": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "supplemental_groups": [
                      "set",
                      "number"
                    ],
                    "sysctl": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "windows_options": [
                      "list",
                      [
                        "object",
                        {
                          "gmsa_credential_spec": "string",
                          "gmsa_credential_spec_name": "string",
                          "host_process": "bool",
                          "run_as_username": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "service_account_name": "string",
              "share_process_namespace": "bool",
              "subdomain": "string",
              "termination_grace_period_seconds": "number",
              "toleration": [
                "list",
                [
                  "object",
                  {
                    "effect": "string",
                    "key": "string",
                    "operator": "string",
                    "toleration_seconds": "string",
                    "value": "string"
                  }
                ]
              ],
              "topology_spread_constraint": [
                "list",
                [
                  "object",
                  {
                    "label_selector": [
                      "list",
                      [
                        "object",
                        {
                          "match_expressions": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "operator": "string",
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "match_labels": [
                            "map",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "match_label_keys": [
                      "set",
                      "string"
                    ],
                    "max_skew": "number",
                    "min_domains": "number",
                    "node_affinity_policy": "string",
                    "node_taints_policy": "string",
                    "topology_key": "string",
                    "when_unsatisfiable": "string"
                  }
                ]
              ],
              "volume": [
                "list",
                [
                  "object",
                  {
                    "aws_elastic_block_store": [
                      "list",
                      [
                        "object",
                        {
                          "fs_type": "string",
                          "partition": "number",
                          "read_only": "bool",
                          "volume_id": "string"
                        }
                      ]
                    ],
                    "azure_disk": [
                      "list",
                      [
                        "object",
                        {
                          "caching_mode": "string",
                          "data_disk_uri": "string",
                          "disk_name": "string",
                          "fs_type": "string",
                          "kind": "string",
                          "read_only": "bool"
                        }
                      ]
                    ],
                    "azure_file": [
                      "list",
                      [
                        "object",
                        {
                          "read_only": "bool",
                          "secret_name": "string",
                          "secret_namespace": "string",
                          "share_name": "string"
                        }
                      ]
                    ],
                    "ceph_fs": [
                      "list",
                      [
                        "object",
                        {
                          "monitors": [
                            "set",
                            "string"
                          ],
                          "path": "string",
                          "read_only": "bool",
                          "secret_file": "string",
                          "secret_ref": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "namespace": "string"
                              }
                            ]
                          ],
                          "user": "string"
                        }
                      ]
                    ],
                    "cinder": [
                      "list",
                      [
                        "object",
                        {
                          "fs_type": "string",
                          "read_only": "bool",
                          "volume_id": "string"
                        }
                      ]
                    ],
                    "config_map": [
                      "list",
                      [
                        "object",
                        {
                          "default_mode": "string",
                          "items": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "mode": "string",
                                "path": "string"
                              }
                            ]
                          ],
                          "name": "string",
                          "optional": "bool"
                        }
                      ]
                    ],
                    "csi": [
                      "list",
                      [
                        "object",
                        {
                          "driver": "string",
                          "fs_type": "string",
                          "node_publish_secret_ref": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string"
                              }
                            ]
                          ],
                          "read_only": "bool",
                          "volume_attributes": [
                            "map",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "downward_api": [
                      "list",
                      [
                        "object",
                        {
                          "default_mode": "string",
                          "items": [
                            "list",
                            [
                              "object",
                              {
                                "field_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "api_version": "string",
                                      "field_path": "string"
                                    }
                                  ]
                                ],
                                "mode": "string",
                                "path": "string",
                                "resource_field_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "container_name": "string",
                                      "divisor": "string",
                                      "resource": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "empty_dir": [
                      "list",
                      [
                        "object",
                        {
                          "medium": "string",
                          "size_limit": "string"
                        }
                      ]
                    ],
                    "ephemeral": [
                      "list",
                      [
                        "object",
                        {
                          "volume_claim_template": [
                            "list",
                            [
                              "object",
                              {
                                "metadata": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "annotations": [
                                        "map",
                                        "string"
                                      ],
                                      "labels": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "spec": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "access_modes": [
                                        "set",
                                        "string"
                                      ],
                                      "resources": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "limits": [
                                              "map",
                                              "string"
                                            ],
                                            "requests": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ],
                                      "selector": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "match_expressions": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "key": "string",
                                                  "operator": "string",
                                                  "values": [
                                                    "set",
                                                    "string"
                                                  ]
                                                }
                                              ]
                                            ],
                                            "match_labels": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ],
                                      "storage_class_name": "string",
                                      "volume_mode": "string",
                                      "volume_name": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "fc": [
                      "list",
                      [
                        "object",
                        {
                          "fs_type": "string",
                          "lun": "number",
                          "read_only": "bool",
                          "target_ww_ns": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "flex_volume": [
                      "list",
                      [
                        "object",
                        {
                          "driver": "string",
                          "fs_type": "string",
                          "options": [
                            "map",
                            "string"
                          ],
                          "read_only": "bool",
                          "secret_ref": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "namespace": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "flocker": [
                      "list",
                      [
                        "object",
                        {
                          "dataset_name": "string",
                          "dataset_uuid": "string"
                        }
                      ]
                    ],
                    "gce_persistent_disk": [
                      "list",
                      [
                        "object",
                        {
                          "fs_type": "string",
                          "partition": "number",
                          "pd_name": "string",
                          "read_only": "bool"
                        }
                      ]
                    ],
                    "git_repo": [
                      "list",
                      [
                        "object",
                        {
                          "directory": "string",
                          "repository": "string",
                          "revision": "string"
                        }
                      ]
                    ],
                    "glusterfs": [
                      "list",
                      [
                        "object",
                        {
                          "endpoints_name": "string",
                          "path": "string",
                          "read_only": "bool"
                        }
                      ]
                    ],
                    "host_path": [
                      "list",
                      [
                        "object",
                        {
                          "path": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "iscsi": [
                      "list",
                      [
                        "object",
                        {
                          "fs_type": "string",
                          "iqn": "string",
                          "iscsi_interface": "string",
                          "lun": "number",
                          "read_only": "bool",
                          "target_portal": "string"
                        }
                      ]
                    ],
                    "local": [
                      "list",
                      [
                        "object",
                        {
                          "path": "string"
                        }
                      ]
                    ],
                    "name": "string",
                    "nfs": [
                      "list",
                      [
                        "object",
                        {
                          "path": "string",
                          "read_only": "bool",
                          "server": "string"
                        }
                      ]
                    ],
                    "persistent_volume_claim": [
                      "list",
                      [
                        "object",
                        {
                          "claim_name": "string",
                          "read_only": "bool"
                        }
                      ]
                    ],
                    "photon_persistent_disk": [
                      "list",
                      [
                        "object",
                        {
                          "fs_type": "string",
                          "pd_id": "string"
                        }
                      ]
                    ],
                    "projected": [
                      "list",
                      [
                        "object",
                        {
                          "default_mode": "string",
                          "sources": [
                            "list",
                            [
                              "object",
                              {
                                "config_map": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "items": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "mode": "string",
                                            "path": "string"
                                          }
                                        ]
                                      ],
                                      "name": "string",
                                      "optional": "bool"
                                    }
                                  ]
                                ],
                                "downward_api": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "items": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "field_ref": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "api_version": "string",
                                                  "field_path": "string"
                                                }
                                              ]
                                            ],
                                            "mode": "string",
                                            "path": "string",
                                            "resource_field_ref": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "container_name": "string",
                                                  "divisor": "string",
                                                  "resource": "string"
                                                }
                                              ]
                                            ]
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "secret": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "items": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "mode": "string",
                                            "path": "string"
                                          }
                                        ]
                                      ],
                                      "name": "string",
                                      "optional": "bool"
                                    }
                                  ]
                                ],
                                "service_account_token": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "audience": "string",
                                      "expiration_seconds": "number",
                                      "path": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "quobyte": [
                      "list",
                      [
                        "object",
                        {
                          "group": "string",
                          "read_only": "bool",
                          "registry": "string",
                          "user": "string",
                          "volume": "string"
                        }
                      ]
                    ],
                    "rbd": [
                      "list",
                      [
                        "object",
                        {
                          "ceph_monitors": [
                            "set",
                            "string"
                          ],
                          "fs_type": "string",
                          "keyring": "string",
                          "rados_user": "string",
                          "rbd_image": "string",
                          "rbd_pool": "string",
                          "read_only": "bool",
                          "secret_ref": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "namespace": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "secret": [
                      "list",
                      [
                        "object",
                        {
                          "default_mode": "string",
                          "items": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "mode": "string",
                                "path": "string"
                              }
                            ]
                          ],
                          "optional": "bool",
                          "secret_name": "string"
                        }
                      ]
                    ],
                    "vsphere_volume": [
                      "list",
                      [
                        "object",
                        {
                          "fs_type": "string",
                          "volume_path": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "An unstructured key value map stored with the pod that may be used to store arbitrary metadata. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
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
              "description": "Map of string keys and values that can be used to organize and categorize (scope and select) the pod. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of the pod, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description": "Namespace defines the space within which name of the pod must be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this pod that can be used by clients to determine when pod has changed. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The unique in time and space value for this pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Standard pod's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "A pod is a group of one or more containers, the shared storage for those containers, and options about how to run the containers. Pods are always co-located and co-scheduled, and run in a shared context. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod/.",
    "description_kind": "plain"
  },
  "version": 0
}`

func KubernetesPodSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(kubernetesPod), &result)
	return &result
}
