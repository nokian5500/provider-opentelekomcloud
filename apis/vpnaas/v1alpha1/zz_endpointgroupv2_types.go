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

type EndpointGroupV2InitParameters struct {

	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints []*string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// The name of the group. Changing this updates the name of
	// the existing group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// region argument of the provider is used. Changing this creates a new group.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type EndpointGroupV2Observation struct {

	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints []*string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the group. Changing this updates the name of
	// the existing group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// region argument of the provider is used. Changing this creates a new group.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type EndpointGroupV2Parameters struct {

	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	// +kubebuilder:validation:Optional
	Endpoints []*string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// The name of the group. Changing this updates the name of
	// the existing group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// region argument of the provider is used. Changing this creates a new group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// EndpointGroupV2Spec defines the desired state of EndpointGroupV2
type EndpointGroupV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointGroupV2Parameters `json:"forProvider"`
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
	InitProvider EndpointGroupV2InitParameters `json:"initProvider,omitempty"`
}

// EndpointGroupV2Status defines the observed state of EndpointGroupV2.
type EndpointGroupV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointGroupV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointGroupV2 is the Schema for the EndpointGroupV2s API. Manages a VPNAAS Endpoint Group resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type EndpointGroupV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointGroupV2Spec   `json:"spec"`
	Status            EndpointGroupV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointGroupV2List contains a list of EndpointGroupV2s
type EndpointGroupV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointGroupV2 `json:"items"`
}

// Repository type metadata.
var (
	EndpointGroupV2_Kind             = "EndpointGroupV2"
	EndpointGroupV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointGroupV2_Kind}.String()
	EndpointGroupV2_KindAPIVersion   = EndpointGroupV2_Kind + "." + CRDGroupVersion.String()
	EndpointGroupV2_GroupVersionKind = CRDGroupVersion.WithKind(EndpointGroupV2_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointGroupV2{}, &EndpointGroupV2List{})
}
