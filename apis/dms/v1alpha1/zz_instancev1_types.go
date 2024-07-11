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

type InstanceV1InitParameters struct {

	// Indicates a username. A username consists of 4 to 64 characters
	// and supports only letters, digits, and hyphens (-).
	AccessUser *string `json:"accessUser,omitempty" tf:"access_user,omitempty"`

	// Indicates the ID of an AZ. The parameter value can not be
	// left blank or an empty array. For details, see section
	// Querying AZ Information.
	AvailableZones []*string `json:"availableZones,omitempty" tf:"available_zones,omitempty"`

	// Indicates the description of an instance. It is a character
	// string containing not more than 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates a message engine. Only kafka is supported now.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Indicates the version of a message engine.
	// Only 2.3.0 is supported now.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Indicates the time at which a maintenance time window starts.
	// Format: HH:mm.
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Indicates the time at which a maintenance time window ends.
	// Format: HH:mm.
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Indicates the name of an instance. An instance name starts with a letter,
	// consists of 4 to 64 characters, and supports only letters, digits, and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This parameter is mandatory when a kafka instance is created.
	// Indicates the maximum number of topics in a Kafka instance.
	PartitionNum *float64 `json:"partitionNum,omitempty" tf:"partition_num,omitempty"`

	// Indicates a product ID.
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Indicates the action to be taken when the memory usage reaches
	// the disk capacity threshold. The possible values are:
	RetentionPolicy *string `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Indicates the ID of a security group.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// This parameter is mandatory if the engine is kafka.
	// Indicates the baseline bandwidth of a Kafka instance, that is, the maximum amount
	// of data transferred per unit time. Unit: byte/s. Options: 100MB, 300MB,
	// 600MB, 1200MB.
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// Indicates the message storage space. Value range:
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Indicates the storage I/O specification. Options for a Kafka instance:
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// Indicates the ID of the subnet (OpenStack network ID).
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Indicates the ID of a VPC (OpenStack router ID).
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type InstanceV1Observation struct {

	// Indicates a username. A username consists of 4 to 64 characters
	// and supports only letters, digits, and hyphens (-).
	AccessUser *string `json:"accessUser,omitempty" tf:"access_user,omitempty"`

	// Indicates the ID of an AZ. The parameter value can not be
	// left blank or an empty array. For details, see section
	// Querying AZ Information.
	AvailableZones []*string `json:"availableZones,omitempty" tf:"available_zones,omitempty"`

	// Indicates the IP address of an instance.
	ConnectAddress *string `json:"connectAddress,omitempty" tf:"connect_address,omitempty"`

	// Indicates the time when an instance is created. The time is in the format
	// of timestamp, that is, the offset milliseconds from 1970-01-01 00:00:00 UTC to the specified time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Indicates the description of an instance. It is a character
	// string containing not more than 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates a message engine. Only kafka is supported now.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Indicates the version of a message engine.
	// Only 2.3.0 is supported now.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the time at which a maintenance time window starts.
	// Format: HH:mm.
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Indicates the time at which a maintenance time window ends.
	// Format: HH:mm.
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Indicates the name of an instance. An instance name starts with a letter,
	// consists of 4 to 64 characters, and supports only letters, digits, and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrderID *string `json:"orderId,omitempty" tf:"order_id,omitempty"`

	// This parameter is mandatory when a kafka instance is created.
	// Indicates the maximum number of topics in a Kafka instance.
	PartitionNum *float64 `json:"partitionNum,omitempty" tf:"partition_num,omitempty"`

	// Indicates the port number of an instance.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Indicates a product ID.
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Indicates a resource specifications identifier.
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty" tf:"resource_spec_code,omitempty"`

	// Indicates the action to be taken when the memory usage reaches
	// the disk capacity threshold. The possible values are:
	RetentionPolicy *string `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Indicates the ID of a security group.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Indicates the name of a security group.
	SecurityGroupName *string `json:"securityGroupName,omitempty" tf:"security_group_name,omitempty"`

	// This parameter is mandatory if the engine is kafka.
	// Indicates the baseline bandwidth of a Kafka instance, that is, the maximum amount
	// of data transferred per unit time. Unit: byte/s. Options: 100MB, 300MB,
	// 600MB, 1200MB.
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// Indicates the status of an instance. For details, see section Instance Status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Indicates the message storage space. Value range:
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Indicates the storage I/O specification. Options for a Kafka instance:
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// Indicates the ID of the subnet (OpenStack network ID).
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Indicates the name of a subnet.
	SubnetName *string `json:"subnetName,omitempty" tf:"subnet_name,omitempty"`

	// Indicates an instance type. Options: single and cluster.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Indicates the used message storage space. Unit: GB
	UsedStorageSpace *float64 `json:"usedStorageSpace,omitempty" tf:"used_storage_space,omitempty"`

	// Indicates a user ID.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Indicates a username.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Indicates the ID of a VPC (OpenStack router ID).
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Indicates the name of an instance. An instance name starts with a letter,
	// consists of 4 to 64 characters, and supports only letters, digits, and hyphens (-).
	VPCName *string `json:"vpcName,omitempty" tf:"vpc_name,omitempty"`
}

type InstanceV1Parameters struct {

	// Indicates a username. A username consists of 4 to 64 characters
	// and supports only letters, digits, and hyphens (-).
	// +kubebuilder:validation:Optional
	AccessUser *string `json:"accessUser,omitempty" tf:"access_user,omitempty"`

	// Indicates the ID of an AZ. The parameter value can not be
	// left blank or an empty array. For details, see section
	// Querying AZ Information.
	// +kubebuilder:validation:Optional
	AvailableZones []*string `json:"availableZones,omitempty" tf:"available_zones,omitempty"`

	// Indicates the description of an instance. It is a character
	// string containing not more than 1024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates a message engine. Only kafka is supported now.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Indicates the version of a message engine.
	// Only 2.3.0 is supported now.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Indicates the time at which a maintenance time window starts.
	// Format: HH:mm.
	// +kubebuilder:validation:Optional
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Indicates the time at which a maintenance time window ends.
	// Format: HH:mm.
	// +kubebuilder:validation:Optional
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Indicates the name of an instance. An instance name starts with a letter,
	// consists of 4 to 64 characters, and supports only letters, digits, and hyphens (-).
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This parameter is mandatory when a kafka instance is created.
	// Indicates the maximum number of topics in a Kafka instance.
	// +kubebuilder:validation:Optional
	PartitionNum *float64 `json:"partitionNum,omitempty" tf:"partition_num,omitempty"`

	// Indicates the password of an instance. An instance password
	// must meet the following complexity requirements: Must be 8 to 32 characters long.
	// Must contain at least 2 of the following character types: lowercase letters, uppercase
	// letters, digits, and special characters (~!@#$%^&*()-_=+\|[{}]:'",<.>/?).
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Indicates a product ID.
	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Indicates the action to be taken when the memory usage reaches
	// the disk capacity threshold. The possible values are:
	// +kubebuilder:validation:Optional
	RetentionPolicy *string `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Indicates the ID of a security group.
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// This parameter is mandatory if the engine is kafka.
	// Indicates the baseline bandwidth of a Kafka instance, that is, the maximum amount
	// of data transferred per unit time. Unit: byte/s. Options: 100MB, 300MB,
	// 600MB, 1200MB.
	// +kubebuilder:validation:Optional
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// Indicates the message storage space. Value range:
	// +kubebuilder:validation:Optional
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Indicates the storage I/O specification. Options for a Kafka instance:
	// +kubebuilder:validation:Optional
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// Indicates the ID of the subnet (OpenStack network ID).
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Indicates the ID of a VPC (OpenStack router ID).
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

// InstanceV1Spec defines the desired state of InstanceV1
type InstanceV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceV1Parameters `json:"forProvider"`
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
	InitProvider InstanceV1InitParameters `json:"initProvider,omitempty"`
}

// InstanceV1Status defines the observed state of InstanceV1.
type InstanceV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceV1 is the Schema for the InstanceV1s API. Manages a DMS Instance v1 resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type InstanceV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availableZones) || (has(self.initProvider) && has(self.initProvider.availableZones))",message="spec.forProvider.availableZones is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || (has(self.initProvider) && has(self.initProvider.engine))",message="spec.forProvider.engine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engineVersion) || (has(self.initProvider) && has(self.initProvider.engineVersion))",message="spec.forProvider.engineVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.productId) || (has(self.initProvider) && has(self.initProvider.productId))",message="spec.forProvider.productId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.securityGroupId) || (has(self.initProvider) && has(self.initProvider.securityGroupId))",message="spec.forProvider.securityGroupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageSpace) || (has(self.initProvider) && has(self.initProvider.storageSpace))",message="spec.forProvider.storageSpace is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageSpecCode) || (has(self.initProvider) && has(self.initProvider.storageSpecCode))",message="spec.forProvider.storageSpecCode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnetId) || (has(self.initProvider) && has(self.initProvider.subnetId))",message="spec.forProvider.subnetId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcId) || (has(self.initProvider) && has(self.initProvider.vpcId))",message="spec.forProvider.vpcId is a required parameter"
	Spec   InstanceV1Spec   `json:"spec"`
	Status InstanceV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceV1List contains a list of InstanceV1s
type InstanceV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceV1 `json:"items"`
}

// Repository type metadata.
var (
	InstanceV1_Kind             = "InstanceV1"
	InstanceV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceV1_Kind}.String()
	InstanceV1_KindAPIVersion   = InstanceV1_Kind + "." + CRDGroupVersion.String()
	InstanceV1_GroupVersionKind = CRDGroupVersion.WithKind(InstanceV1_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceV1{}, &InstanceV1List{})
}
