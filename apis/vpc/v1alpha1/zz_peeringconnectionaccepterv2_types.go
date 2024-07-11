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

type PeeringConnectionAccepterV2InitParameters struct {

	// - Whether or not to accept the peering request. Defaults to false.
	Accept *bool `json:"accept,omitempty" tf:"accept,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The VPC Peering Connection ID to manage. Changing this creates a new VPC peering connection accepter.
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

type PeeringConnectionAccepterV2Observation struct {

	// - Whether or not to accept the peering request. Defaults to false.
	Accept *bool `json:"accept,omitempty" tf:"accept,omitempty"`

	// The VPC peering connection ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The VPC peering connection name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Tenant Id of the accepter tenant.
	PeerTenantID *string `json:"peerTenantId,omitempty" tf:"peer_tenant_id,omitempty"`

	// The VPC ID of the accepter tenant.
	PeerVPCID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The VPC peering connection status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The ID of requester VPC involved in a VPC peering connection.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The VPC Peering Connection ID to manage. Changing this creates a new VPC peering connection accepter.
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

type PeeringConnectionAccepterV2Parameters struct {

	// - Whether or not to accept the peering request. Defaults to false.
	// +kubebuilder:validation:Optional
	Accept *bool `json:"accept,omitempty" tf:"accept,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The VPC Peering Connection ID to manage. Changing this creates a new VPC peering connection accepter.
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

// PeeringConnectionAccepterV2Spec defines the desired state of PeeringConnectionAccepterV2
type PeeringConnectionAccepterV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PeeringConnectionAccepterV2Parameters `json:"forProvider"`
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
	InitProvider PeeringConnectionAccepterV2InitParameters `json:"initProvider,omitempty"`
}

// PeeringConnectionAccepterV2Status defines the observed state of PeeringConnectionAccepterV2.
type PeeringConnectionAccepterV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PeeringConnectionAccepterV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnectionAccepterV2 is the Schema for the PeeringConnectionAccepterV2s API. Manages a VPC Peering Accepter resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type PeeringConnectionAccepterV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcPeeringConnectionId) || (has(self.initProvider) && has(self.initProvider.vpcPeeringConnectionId))",message="spec.forProvider.vpcPeeringConnectionId is a required parameter"
	Spec   PeeringConnectionAccepterV2Spec   `json:"spec"`
	Status PeeringConnectionAccepterV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnectionAccepterV2List contains a list of PeeringConnectionAccepterV2s
type PeeringConnectionAccepterV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PeeringConnectionAccepterV2 `json:"items"`
}

// Repository type metadata.
var (
	PeeringConnectionAccepterV2_Kind             = "PeeringConnectionAccepterV2"
	PeeringConnectionAccepterV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PeeringConnectionAccepterV2_Kind}.String()
	PeeringConnectionAccepterV2_KindAPIVersion   = PeeringConnectionAccepterV2_Kind + "." + CRDGroupVersion.String()
	PeeringConnectionAccepterV2_GroupVersionKind = CRDGroupVersion.WithKind(PeeringConnectionAccepterV2_Kind)
)

func init() {
	SchemeBuilder.Register(&PeeringConnectionAccepterV2{}, &PeeringConnectionAccepterV2List{})
}
