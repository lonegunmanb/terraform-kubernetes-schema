package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-kubernetes-schema/v2/generated/data"
	"github.com/lonegunmanb/terraform-kubernetes-schema/v2/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["kubernetes_annotations"] = resource.KubernetesAnnotationsSchema()  
	resources["kubernetes_api_service"] = resource.KubernetesApiServiceSchema()  
	resources["kubernetes_api_service_v1"] = resource.KubernetesApiServiceV1Schema()  
	resources["kubernetes_certificate_signing_request"] = resource.KubernetesCertificateSigningRequestSchema()  
	resources["kubernetes_certificate_signing_request_v1"] = resource.KubernetesCertificateSigningRequestV1Schema()  
	resources["kubernetes_cluster_role"] = resource.KubernetesClusterRoleSchema()  
	resources["kubernetes_cluster_role_binding"] = resource.KubernetesClusterRoleBindingSchema()  
	resources["kubernetes_cluster_role_binding_v1"] = resource.KubernetesClusterRoleBindingV1Schema()  
	resources["kubernetes_cluster_role_v1"] = resource.KubernetesClusterRoleV1Schema()  
	resources["kubernetes_config_map"] = resource.KubernetesConfigMapSchema()  
	resources["kubernetes_config_map_v1"] = resource.KubernetesConfigMapV1Schema()  
	resources["kubernetes_config_map_v1_data"] = resource.KubernetesConfigMapV1DataSchema()  
	resources["kubernetes_cron_job"] = resource.KubernetesCronJobSchema()  
	resources["kubernetes_cron_job_v1"] = resource.KubernetesCronJobV1Schema()  
	resources["kubernetes_csi_driver"] = resource.KubernetesCsiDriverSchema()  
	resources["kubernetes_csi_driver_v1"] = resource.KubernetesCsiDriverV1Schema()  
	resources["kubernetes_daemon_set_v1"] = resource.KubernetesDaemonSetV1Schema()  
	resources["kubernetes_daemonset"] = resource.KubernetesDaemonsetSchema()  
	resources["kubernetes_default_service_account"] = resource.KubernetesDefaultServiceAccountSchema()  
	resources["kubernetes_default_service_account_v1"] = resource.KubernetesDefaultServiceAccountV1Schema()  
	resources["kubernetes_deployment"] = resource.KubernetesDeploymentSchema()  
	resources["kubernetes_deployment_v1"] = resource.KubernetesDeploymentV1Schema()  
	resources["kubernetes_endpoint_slice_v1"] = resource.KubernetesEndpointSliceV1Schema()  
	resources["kubernetes_endpoints"] = resource.KubernetesEndpointsSchema()  
	resources["kubernetes_endpoints_v1"] = resource.KubernetesEndpointsV1Schema()  
	resources["kubernetes_env"] = resource.KubernetesEnvSchema()  
	resources["kubernetes_horizontal_pod_autoscaler"] = resource.KubernetesHorizontalPodAutoscalerSchema()  
	resources["kubernetes_horizontal_pod_autoscaler_v1"] = resource.KubernetesHorizontalPodAutoscalerV1Schema()  
	resources["kubernetes_horizontal_pod_autoscaler_v2"] = resource.KubernetesHorizontalPodAutoscalerV2Schema()  
	resources["kubernetes_horizontal_pod_autoscaler_v2beta2"] = resource.KubernetesHorizontalPodAutoscalerV2Beta2Schema()  
	resources["kubernetes_ingress"] = resource.KubernetesIngressSchema()  
	resources["kubernetes_ingress_class"] = resource.KubernetesIngressClassSchema()  
	resources["kubernetes_ingress_class_v1"] = resource.KubernetesIngressClassV1Schema()  
	resources["kubernetes_ingress_v1"] = resource.KubernetesIngressV1Schema()  
	resources["kubernetes_job"] = resource.KubernetesJobSchema()  
	resources["kubernetes_job_v1"] = resource.KubernetesJobV1Schema()  
	resources["kubernetes_labels"] = resource.KubernetesLabelsSchema()  
	resources["kubernetes_limit_range"] = resource.KubernetesLimitRangeSchema()  
	resources["kubernetes_limit_range_v1"] = resource.KubernetesLimitRangeV1Schema()  
	resources["kubernetes_manifest"] = resource.KubernetesManifestSchema()  
	resources["kubernetes_mutating_webhook_configuration"] = resource.KubernetesMutatingWebhookConfigurationSchema()  
	resources["kubernetes_mutating_webhook_configuration_v1"] = resource.KubernetesMutatingWebhookConfigurationV1Schema()  
	resources["kubernetes_namespace"] = resource.KubernetesNamespaceSchema()  
	resources["kubernetes_namespace_v1"] = resource.KubernetesNamespaceV1Schema()  
	resources["kubernetes_network_policy"] = resource.KubernetesNetworkPolicySchema()  
	resources["kubernetes_network_policy_v1"] = resource.KubernetesNetworkPolicyV1Schema()  
	resources["kubernetes_node_taint"] = resource.KubernetesNodeTaintSchema()  
	resources["kubernetes_persistent_volume"] = resource.KubernetesPersistentVolumeSchema()  
	resources["kubernetes_persistent_volume_claim"] = resource.KubernetesPersistentVolumeClaimSchema()  
	resources["kubernetes_persistent_volume_claim_v1"] = resource.KubernetesPersistentVolumeClaimV1Schema()  
	resources["kubernetes_persistent_volume_v1"] = resource.KubernetesPersistentVolumeV1Schema()  
	resources["kubernetes_pod"] = resource.KubernetesPodSchema()  
	resources["kubernetes_pod_disruption_budget"] = resource.KubernetesPodDisruptionBudgetSchema()  
	resources["kubernetes_pod_disruption_budget_v1"] = resource.KubernetesPodDisruptionBudgetV1Schema()  
	resources["kubernetes_pod_security_policy"] = resource.KubernetesPodSecurityPolicySchema()  
	resources["kubernetes_pod_security_policy_v1beta1"] = resource.KubernetesPodSecurityPolicyV1Beta1Schema()  
	resources["kubernetes_pod_v1"] = resource.KubernetesPodV1Schema()  
	resources["kubernetes_priority_class"] = resource.KubernetesPriorityClassSchema()  
	resources["kubernetes_priority_class_v1"] = resource.KubernetesPriorityClassV1Schema()  
	resources["kubernetes_replication_controller"] = resource.KubernetesReplicationControllerSchema()  
	resources["kubernetes_replication_controller_v1"] = resource.KubernetesReplicationControllerV1Schema()  
	resources["kubernetes_resource_quota"] = resource.KubernetesResourceQuotaSchema()  
	resources["kubernetes_resource_quota_v1"] = resource.KubernetesResourceQuotaV1Schema()  
	resources["kubernetes_role"] = resource.KubernetesRoleSchema()  
	resources["kubernetes_role_binding"] = resource.KubernetesRoleBindingSchema()  
	resources["kubernetes_role_binding_v1"] = resource.KubernetesRoleBindingV1Schema()  
	resources["kubernetes_role_v1"] = resource.KubernetesRoleV1Schema()  
	resources["kubernetes_runtime_class_v1"] = resource.KubernetesRuntimeClassV1Schema()  
	resources["kubernetes_secret"] = resource.KubernetesSecretSchema()  
	resources["kubernetes_secret_v1"] = resource.KubernetesSecretV1Schema()  
	resources["kubernetes_secret_v1_data"] = resource.KubernetesSecretV1DataSchema()  
	resources["kubernetes_service"] = resource.KubernetesServiceSchema()  
	resources["kubernetes_service_account"] = resource.KubernetesServiceAccountSchema()  
	resources["kubernetes_service_account_v1"] = resource.KubernetesServiceAccountV1Schema()  
	resources["kubernetes_service_v1"] = resource.KubernetesServiceV1Schema()  
	resources["kubernetes_stateful_set"] = resource.KubernetesStatefulSetSchema()  
	resources["kubernetes_stateful_set_v1"] = resource.KubernetesStatefulSetV1Schema()  
	resources["kubernetes_storage_class"] = resource.KubernetesStorageClassSchema()  
	resources["kubernetes_storage_class_v1"] = resource.KubernetesStorageClassV1Schema()  
	resources["kubernetes_token_request_v1"] = resource.KubernetesTokenRequestV1Schema()  
	resources["kubernetes_validating_webhook_configuration"] = resource.KubernetesValidatingWebhookConfigurationSchema()  
	resources["kubernetes_validating_webhook_configuration_v1"] = resource.KubernetesValidatingWebhookConfigurationV1Schema()  
	dataSources["kubernetes_all_namespaces"] = data.KubernetesAllNamespacesSchema()  
	dataSources["kubernetes_config_map"] = data.KubernetesConfigMapSchema()  
	dataSources["kubernetes_config_map_v1"] = data.KubernetesConfigMapV1Schema()  
	dataSources["kubernetes_endpoints_v1"] = data.KubernetesEndpointsV1Schema()  
	dataSources["kubernetes_ingress"] = data.KubernetesIngressSchema()  
	dataSources["kubernetes_ingress_v1"] = data.KubernetesIngressV1Schema()  
	dataSources["kubernetes_mutating_webhook_configuration_v1"] = data.KubernetesMutatingWebhookConfigurationV1Schema()  
	dataSources["kubernetes_namespace"] = data.KubernetesNamespaceSchema()  
	dataSources["kubernetes_namespace_v1"] = data.KubernetesNamespaceV1Schema()  
	dataSources["kubernetes_nodes"] = data.KubernetesNodesSchema()  
	dataSources["kubernetes_persistent_volume_claim"] = data.KubernetesPersistentVolumeClaimSchema()  
	dataSources["kubernetes_persistent_volume_claim_v1"] = data.KubernetesPersistentVolumeClaimV1Schema()  
	dataSources["kubernetes_persistent_volume_v1"] = data.KubernetesPersistentVolumeV1Schema()  
	dataSources["kubernetes_pod"] = data.KubernetesPodSchema()  
	dataSources["kubernetes_pod_v1"] = data.KubernetesPodV1Schema()  
	dataSources["kubernetes_resource"] = data.KubernetesResourceSchema()  
	dataSources["kubernetes_resources"] = data.KubernetesResourcesSchema()  
	dataSources["kubernetes_secret"] = data.KubernetesSecretSchema()  
	dataSources["kubernetes_secret_v1"] = data.KubernetesSecretV1Schema()  
	dataSources["kubernetes_server_version"] = data.KubernetesServerVersionSchema()  
	dataSources["kubernetes_service"] = data.KubernetesServiceSchema()  
	dataSources["kubernetes_service_account"] = data.KubernetesServiceAccountSchema()  
	dataSources["kubernetes_service_account_v1"] = data.KubernetesServiceAccountV1Schema()  
	dataSources["kubernetes_service_v1"] = data.KubernetesServiceV1Schema()  
	dataSources["kubernetes_storage_class"] = data.KubernetesStorageClassSchema()  
	dataSources["kubernetes_storage_class_v1"] = data.KubernetesStorageClassV1Schema()  
	Resources = resources
	DataSources = dataSources
}