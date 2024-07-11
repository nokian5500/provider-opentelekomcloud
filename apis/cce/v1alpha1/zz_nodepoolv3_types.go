// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DataVolumesInitParameters struct {

	// Disk expansion parameters.
	// Please use alternative parameter extend_params.
	ExtendParam *string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Disk expansion parameters. A list of strings which describes additional disk parameters.
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// The Encryption KMS ID of the system volume. By default, it tries to get from env by OS_KMS_ID.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Disk size in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Disk type.
	Volumetype *string `json:"volumetype,omitempty" tf:"volumetype,omitempty"`
}

type DataVolumesObservation struct {

	// Disk expansion parameters.
	// Please use alternative parameter extend_params.
	ExtendParam *string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Disk expansion parameters. A list of strings which describes additional disk parameters.
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// The Encryption KMS ID of the system volume. By default, it tries to get from env by OS_KMS_ID.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Disk size in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Disk type.
	Volumetype *string `json:"volumetype,omitempty" tf:"volumetype,omitempty"`
}

type DataVolumesParameters struct {

	// Disk expansion parameters.
	// Please use alternative parameter extend_params.
	// +kubebuilder:validation:Optional
	ExtendParam *string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Disk expansion parameters. A list of strings which describes additional disk parameters.
	// +kubebuilder:validation:Optional
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// The Encryption KMS ID of the system volume. By default, it tries to get from env by OS_KMS_ID.
	// +kubebuilder:validation:Optional
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Disk size in GB.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Disk type.
	// +kubebuilder:validation:Optional
	Volumetype *string `json:"volumetype" tf:"volumetype,omitempty"`
}

type NodePoolV3InitParameters struct {

	// IAM agency name. Changing this parameter will create a new resource.
	AgencyName *string `json:"agencyName,omitempty" tf:"agency_name,omitempty"`

	// Specify the name of the available partition (AZ). If zone is not
	// specified than node_pool will be in randomly selected AZ. The default value is random. Changing
	// this parameter will create a new resource.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// ID of the cluster. Changing this parameter will create a new resource.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Represents the data disk to be created. Changing this parameter will create a new resource.
	DataVolumes []DataVolumesInitParameters `json:"dataVolumes,omitempty" tf:"data_volumes,omitempty"`

	// Available disk space of a single Docker container on the node using the device mapper.
	// Changing this parameter will create a new node pool.
	DockerBaseSize *float64 `json:"dockerBaseSize,omitempty" tf:"docker_base_size,omitempty"`

	// ConfigMap of the Docker data disk.
	// Changing this parameter will create a new node.
	DockerLvmConfigOverride *string `json:"dockerLvmConfigOverride,omitempty" tf:"docker_lvm_config_override,omitempty"`

	// Specifies the flavor id. Changing this parameter will create a new resource.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Initial number of expected nodes in the node pool.
	InitialNodeCount *float64 `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`

	// Tags of a Kubernetes node, key/value pair format.
	K8STags map[string]*string `json:"k8sTags,omitempty" tf:"k8s_tags,omitempty"`

	// Key pair name when logging in to select the key pair mode.
	// This parameter and password are alternative. Changing this parameter will create a new resource.
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Maximum number of nodes allowed if auto scaling is enabled.
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// The maximum number of instances a node is allowed to create.
	// Changing this parameter will create a new node pool.
	MaxPods *float64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`

	// Minimum number of nodes allowed if auto scaling is enabled.
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// Node Pool Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Node OS. Changing this parameter will create a new resource.
	// Supported OS depends on kubernetes version of the cluster.
	Os *string `json:"os,omitempty" tf:"os,omitempty"`

	// Script required after installation. The input value can be a Base64 encoded string or not.
	// Changing this parameter will create a new resource.
	Postinstall *string `json:"postinstall,omitempty" tf:"postinstall,omitempty"`

	// Script required before installation. The input value can be a Base64 encoded string or not.
	// Changing this parameter will create a new resource.
	Preinstall *string `json:"preinstall,omitempty" tf:"preinstall,omitempty"`

	// Weight of a node pool. A node pool with a higher weight has a higher priority during scaling.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// It corresponds to the system disk related configuration. Changing this parameter will create a new resource.
	RootVolume []RootVolumeInitParameters `json:"rootVolume,omitempty" tf:"root_volume,omitempty"`

	// Container runtime. Changing this parameter will create a new resource.
	// Use with high-caution, may trigger resource recreation. Options are:
	// docker - Docker
	// containerd - Containerd
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Interval between two scaling operations, in minutes.
	ScaleDownCooldownTime *float64 `json:"scaleDownCooldownTime,omitempty" tf:"scale_down_cooldown_time,omitempty"`

	// Whether to enable auto scaling. If Autoscaler is enabled, install the autoscaler add-on to use the auto scaling feature.
	ScaleEnable *bool `json:"scaleEnable,omitempty" tf:"scale_enable,omitempty"`

	// ECS group ID. If this parameter is specified, all nodes in the node pool will be created in this ECS group.
	ServerGroupReference *string `json:"serverGroupReference,omitempty" tf:"server_group_reference,omitempty"`

	// Specifies the json string vary depending on CCE node pools storage options.
	// -> Please refer to the documentation
	// for actual fields.
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// The ID of the subnet to which the NIC belongs. Changing this parameter will create a new resource.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Taints to created nodes to configure anti-affinity.
	Taints []TaintsInitParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// Tag of a VM, key/value pair format. Changing this parameter will create a new resource.
	UserTags map[string]*string `json:"userTags,omitempty" tf:"user_tags,omitempty"`
}

type NodePoolV3Observation struct {

	// IAM agency name. Changing this parameter will create a new resource.
	AgencyName *string `json:"agencyName,omitempty" tf:"agency_name,omitempty"`

	// Specify the name of the available partition (AZ). If zone is not
	// specified than node_pool will be in randomly selected AZ. The default value is random. Changing
	// this parameter will create a new resource.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// ID of the cluster. Changing this parameter will create a new resource.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Represents the data disk to be created. Changing this parameter will create a new resource.
	DataVolumes []DataVolumesObservation `json:"dataVolumes,omitempty" tf:"data_volumes,omitempty"`

	// Available disk space of a single Docker container on the node using the device mapper.
	// Changing this parameter will create a new node pool.
	DockerBaseSize *float64 `json:"dockerBaseSize,omitempty" tf:"docker_base_size,omitempty"`

	// ConfigMap of the Docker data disk.
	// Changing this parameter will create a new node.
	DockerLvmConfigOverride *string `json:"dockerLvmConfigOverride,omitempty" tf:"docker_lvm_config_override,omitempty"`

	// Specifies the flavor id. Changing this parameter will create a new resource.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Specifies a resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Initial number of expected nodes in the node pool.
	InitialNodeCount *float64 `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`

	// Tags of a Kubernetes node, key/value pair format.
	K8STags map[string]*string `json:"k8sTags,omitempty" tf:"k8s_tags,omitempty"`

	// Key pair name when logging in to select the key pair mode.
	// This parameter and password are alternative. Changing this parameter will create a new resource.
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Maximum number of nodes allowed if auto scaling is enabled.
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// The maximum number of instances a node is allowed to create.
	// Changing this parameter will create a new node pool.
	MaxPods *float64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`

	// Minimum number of nodes allowed if auto scaling is enabled.
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// Node Pool Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Node OS. Changing this parameter will create a new resource.
	// Supported OS depends on kubernetes version of the cluster.
	Os *string `json:"os,omitempty" tf:"os,omitempty"`

	// Script required after installation. The input value can be a Base64 encoded string or not.
	// Changing this parameter will create a new resource.
	Postinstall *string `json:"postinstall,omitempty" tf:"postinstall,omitempty"`

	// Script required before installation. The input value can be a Base64 encoded string or not.
	// Changing this parameter will create a new resource.
	Preinstall *string `json:"preinstall,omitempty" tf:"preinstall,omitempty"`

	// Weight of a node pool. A node pool with a higher weight has a higher priority during scaling.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// It corresponds to the system disk related configuration. Changing this parameter will create a new resource.
	RootVolume []RootVolumeObservation `json:"rootVolume,omitempty" tf:"root_volume,omitempty"`

	// Container runtime. Changing this parameter will create a new resource.
	// Use with high-caution, may trigger resource recreation. Options are:
	// docker - Docker
	// containerd - Containerd
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Interval between two scaling operations, in minutes.
	ScaleDownCooldownTime *float64 `json:"scaleDownCooldownTime,omitempty" tf:"scale_down_cooldown_time,omitempty"`

	// Whether to enable auto scaling. If Autoscaler is enabled, install the autoscaler add-on to use the auto scaling feature.
	ScaleEnable *bool `json:"scaleEnable,omitempty" tf:"scale_enable,omitempty"`

	// ECS group ID. If this parameter is specified, all nodes in the node pool will be created in this ECS group.
	ServerGroupReference *string `json:"serverGroupReference,omitempty" tf:"server_group_reference,omitempty"`

	// Node status information.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the json string vary depending on CCE node pools storage options.
	// -> Please refer to the documentation
	// for actual fields.
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// The ID of the subnet to which the NIC belongs. Changing this parameter will create a new resource.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Taints to created nodes to configure anti-affinity.
	Taints []TaintsObservation `json:"taints,omitempty" tf:"taints,omitempty"`

	// Tag of a VM, key/value pair format. Changing this parameter will create a new resource.
	UserTags map[string]*string `json:"userTags,omitempty" tf:"user_tags,omitempty"`
}

type NodePoolV3Parameters struct {

	// IAM agency name. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	AgencyName *string `json:"agencyName,omitempty" tf:"agency_name,omitempty"`

	// Specify the name of the available partition (AZ). If zone is not
	// specified than node_pool will be in randomly selected AZ. The default value is random. Changing
	// this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// ID of the cluster. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Represents the data disk to be created. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	DataVolumes []DataVolumesParameters `json:"dataVolumes,omitempty" tf:"data_volumes,omitempty"`

	// Available disk space of a single Docker container on the node using the device mapper.
	// Changing this parameter will create a new node pool.
	// +kubebuilder:validation:Optional
	DockerBaseSize *float64 `json:"dockerBaseSize,omitempty" tf:"docker_base_size,omitempty"`

	// ConfigMap of the Docker data disk.
	// Changing this parameter will create a new node.
	// +kubebuilder:validation:Optional
	DockerLvmConfigOverride *string `json:"dockerLvmConfigOverride,omitempty" tf:"docker_lvm_config_override,omitempty"`

	// Specifies the flavor id. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Initial number of expected nodes in the node pool.
	// +kubebuilder:validation:Optional
	InitialNodeCount *float64 `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`

	// Tags of a Kubernetes node, key/value pair format.
	// +kubebuilder:validation:Optional
	K8STags map[string]*string `json:"k8sTags,omitempty" tf:"k8s_tags,omitempty"`

	// Key pair name when logging in to select the key pair mode.
	// This parameter and password are alternative. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Maximum number of nodes allowed if auto scaling is enabled.
	// +kubebuilder:validation:Optional
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// The maximum number of instances a node is allowed to create.
	// Changing this parameter will create a new node pool.
	// +kubebuilder:validation:Optional
	MaxPods *float64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`

	// Minimum number of nodes allowed if auto scaling is enabled.
	// +kubebuilder:validation:Optional
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// Node Pool Name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Node OS. Changing this parameter will create a new resource.
	// Supported OS depends on kubernetes version of the cluster.
	// +kubebuilder:validation:Optional
	Os *string `json:"os,omitempty" tf:"os,omitempty"`

	// Key pair name when logging in to select the key pair mode.
	// This parameter and password are alternative. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Script required after installation. The input value can be a Base64 encoded string or not.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Postinstall *string `json:"postinstall,omitempty" tf:"postinstall,omitempty"`

	// Script required before installation. The input value can be a Base64 encoded string or not.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Preinstall *string `json:"preinstall,omitempty" tf:"preinstall,omitempty"`

	// Weight of a node pool. A node pool with a higher weight has a higher priority during scaling.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// It corresponds to the system disk related configuration. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	RootVolume []RootVolumeParameters `json:"rootVolume,omitempty" tf:"root_volume,omitempty"`

	// Container runtime. Changing this parameter will create a new resource.
	// Use with high-caution, may trigger resource recreation. Options are:
	// docker - Docker
	// containerd - Containerd
	// +kubebuilder:validation:Optional
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Interval between two scaling operations, in minutes.
	// +kubebuilder:validation:Optional
	ScaleDownCooldownTime *float64 `json:"scaleDownCooldownTime,omitempty" tf:"scale_down_cooldown_time,omitempty"`

	// Whether to enable auto scaling. If Autoscaler is enabled, install the autoscaler add-on to use the auto scaling feature.
	// +kubebuilder:validation:Optional
	ScaleEnable *bool `json:"scaleEnable,omitempty" tf:"scale_enable,omitempty"`

	// ECS group ID. If this parameter is specified, all nodes in the node pool will be created in this ECS group.
	// +kubebuilder:validation:Optional
	ServerGroupReference *string `json:"serverGroupReference,omitempty" tf:"server_group_reference,omitempty"`

	// Specifies the json string vary depending on CCE node pools storage options.
	// -> Please refer to the documentation
	// for actual fields.
	// +kubebuilder:validation:Optional
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// The ID of the subnet to which the NIC belongs. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Taints to created nodes to configure anti-affinity.
	// +kubebuilder:validation:Optional
	Taints []TaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// Tag of a VM, key/value pair format. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	UserTags map[string]*string `json:"userTags,omitempty" tf:"user_tags,omitempty"`
}

type RootVolumeInitParameters struct {

	// Disk expansion parameters.
	// Please use alternative parameter extend_params.
	ExtendParam *string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Disk expansion parameters. A list of strings which describes additional disk parameters.
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// The Encryption KMS ID of the system volume. By default, it tries to get from env by OS_KMS_ID.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Disk size in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Disk type.
	Volumetype *string `json:"volumetype,omitempty" tf:"volumetype,omitempty"`
}

type RootVolumeObservation struct {

	// Disk expansion parameters.
	// Please use alternative parameter extend_params.
	ExtendParam *string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Disk expansion parameters. A list of strings which describes additional disk parameters.
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// The Encryption KMS ID of the system volume. By default, it tries to get from env by OS_KMS_ID.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Disk size in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Disk type.
	Volumetype *string `json:"volumetype,omitempty" tf:"volumetype,omitempty"`
}

type RootVolumeParameters struct {

	// Disk expansion parameters.
	// Please use alternative parameter extend_params.
	// +kubebuilder:validation:Optional
	ExtendParam *string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Disk expansion parameters. A list of strings which describes additional disk parameters.
	// +kubebuilder:validation:Optional
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// The Encryption KMS ID of the system volume. By default, it tries to get from env by OS_KMS_ID.
	// +kubebuilder:validation:Optional
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Disk size in GB.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Disk type.
	// +kubebuilder:validation:Optional
	Volumetype *string `json:"volumetype" tf:"volumetype,omitempty"`
}

type TaintsInitParameters struct {

	// Available options are NoSchedule, PreferNoSchedule, and NoExecute.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// A key must contain 1 to 63 characters starting with a letter or digit. Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed. A DNS subdomain name can be used as the prefix of a key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A value must start with a letter or digit and can contain a maximum of 63 characters, including letters, digits, hyphens (-), underscores (_), and periods (.).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TaintsObservation struct {

	// Available options are NoSchedule, PreferNoSchedule, and NoExecute.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// A key must contain 1 to 63 characters starting with a letter or digit. Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed. A DNS subdomain name can be used as the prefix of a key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A value must start with a letter or digit and can contain a maximum of 63 characters, including letters, digits, hyphens (-), underscores (_), and periods (.).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TaintsParameters struct {

	// Available options are NoSchedule, PreferNoSchedule, and NoExecute.
	// +kubebuilder:validation:Optional
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// A key must contain 1 to 63 characters starting with a letter or digit. Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed. A DNS subdomain name can be used as the prefix of a key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// A value must start with a letter or digit and can contain a maximum of 63 characters, including letters, digits, hyphens (-), underscores (_), and periods (.).
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// NodePoolV3Spec defines the desired state of NodePoolV3
type NodePoolV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodePoolV3Parameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NodePoolV3InitParameters `json:"initProvider,omitempty"`
}

// NodePoolV3Status defines the observed state of NodePoolV3.
type NodePoolV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodePoolV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolV3 is the Schema for the NodePoolV3s API. Manages a CCE Cluster Node Pool resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type NodePoolV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterId) || (has(self.initProvider) && has(self.initProvider.clusterId))",message="spec.forProvider.clusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataVolumes) || (has(self.initProvider) && has(self.initProvider.dataVolumes))",message="spec.forProvider.dataVolumes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flavor) || (has(self.initProvider) && has(self.initProvider.flavor))",message="spec.forProvider.flavor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.initialNodeCount) || (has(self.initProvider) && has(self.initProvider.initialNodeCount))",message="spec.forProvider.initialNodeCount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rootVolume) || (has(self.initProvider) && has(self.initProvider.rootVolume))",message="spec.forProvider.rootVolume is a required parameter"
	Spec   NodePoolV3Spec   `json:"spec"`
	Status NodePoolV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolV3List contains a list of NodePoolV3s
type NodePoolV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePoolV3 `json:"items"`
}

// Repository type metadata.
var (
	NodePoolV3_Kind             = "NodePoolV3"
	NodePoolV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodePoolV3_Kind}.String()
	NodePoolV3_KindAPIVersion   = NodePoolV3_Kind + "." + CRDGroupVersion.String()
	NodePoolV3_GroupVersionKind = CRDGroupVersion.WithKind(NodePoolV3_Kind)
)

func init() {
	SchemeBuilder.Register(&NodePoolV3{}, &NodePoolV3List{})
}
