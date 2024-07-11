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

type RouterInitParameters struct {

	// The Router(VPC) ID. which VPC network will assicate with.
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// The Region name for this private zone.
	RouterRegion *string `json:"routerRegion,omitempty" tf:"router_region,omitempty"`
}

type RouterObservation struct {

	// The Router(VPC) ID. which VPC network will assicate with.
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// The Region name for this private zone.
	RouterRegion *string `json:"routerRegion,omitempty" tf:"router_region,omitempty"`
}

type RouterParameters struct {

	// The Router(VPC) ID. which VPC network will assicate with.
	// +kubebuilder:validation:Optional
	RouterID *string `json:"routerId" tf:"router_id,omitempty"`

	// The Region name for this private zone.
	// +kubebuilder:validation:Optional
	RouterRegion *string `json:"routerRegion" tf:"router_region,omitempty"`
}

type ZoneV2InitParameters struct {

	// A description of the zone.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The email contact for the zone record.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The name of the zone.   Changing this creates a new DNS zone.
	// -> Note: The . at the end of the name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Routers(VPCs) configuration for the private zone.
	// it is required when type is private.
	Router []RouterInitParameters `json:"router,omitempty" tf:"router,omitempty"`

	// The time to live (TTL) of the zone.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The key/value pairs to associate with the zone.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of zone. Can either be public or private.
	// Changing this creates a new zone.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options. Changing this creates a new zone.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type ZoneV2Observation struct {

	// A description of the zone.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The email contact for the zone record.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of master DNS servers.
	Masters []*string `json:"masters,omitempty" tf:"masters,omitempty"`

	// The name of the zone.   Changing this creates a new DNS zone.
	// -> Note: The . at the end of the name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Routers(VPCs) configuration for the private zone.
	// it is required when type is private.
	Router []RouterObservation `json:"router,omitempty" tf:"router,omitempty"`

	// The time to live (TTL) of the zone.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The key/value pairs to associate with the zone.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of zone. Can either be public or private.
	// Changing this creates a new zone.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options. Changing this creates a new zone.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type ZoneV2Parameters struct {

	// A description of the zone.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The email contact for the zone record.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The name of the zone.   Changing this creates a new DNS zone.
	// -> Note: The . at the end of the name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Routers(VPCs) configuration for the private zone.
	// it is required when type is private.
	// +kubebuilder:validation:Optional
	Router []RouterParameters `json:"router,omitempty" tf:"router,omitempty"`

	// The time to live (TTL) of the zone.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The key/value pairs to associate with the zone.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of zone. Can either be public or private.
	// Changing this creates a new zone.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options. Changing this creates a new zone.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// ZoneV2Spec defines the desired state of ZoneV2
type ZoneV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneV2Parameters `json:"forProvider"`
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
	InitProvider ZoneV2InitParameters `json:"initProvider,omitempty"`
}

// ZoneV2Status defines the observed state of ZoneV2.
type ZoneV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneV2 is the Schema for the ZoneV2s API. Manages a DNS Zones resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type ZoneV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ZoneV2Spec   `json:"spec"`
	Status ZoneV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneV2List contains a list of ZoneV2s
type ZoneV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneV2 `json:"items"`
}

// Repository type metadata.
var (
	ZoneV2_Kind             = "ZoneV2"
	ZoneV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ZoneV2_Kind}.String()
	ZoneV2_KindAPIVersion   = ZoneV2_Kind + "." + CRDGroupVersion.String()
	ZoneV2_GroupVersionKind = CRDGroupVersion.WithKind(ZoneV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ZoneV2{}, &ZoneV2List{})
}
