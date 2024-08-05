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

type FirewallGroupV2InitParameters struct {

	// Administrative up/down status for the firewall group
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the admin_state_up of an existing firewall group.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// A description for the firewall group. Changing this
	// updates the description of an existing firewall group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The egress policy resource id for the firewall group. Changing
	// this updates the egress_policy_id of an existing firewall group.
	EgressPolicyID *string `json:"egressPolicyId,omitempty" tf:"egress_policy_id,omitempty"`

	// The ingress policy resource id for the firewall group. Changing
	// this updates the ingress_policy_id of an existing firewall group.
	IngressPolicyID *string `json:"ingressPolicyId,omitempty" tf:"ingress_policy_id,omitempty"`

	// A name for the firewall group. Changing this
	// updates the name of an existing firewall group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Port(s) to associate this firewall group instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall group.
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the floating IP. Required if admin wants
	// to create a firewall group for another tenant. Changing this creates a new
	// firewall group.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type FirewallGroupV2Observation struct {

	// Administrative up/down status for the firewall group
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the admin_state_up of an existing firewall group.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// A description for the firewall group. Changing this
	// updates the description of an existing firewall group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The egress policy resource id for the firewall group. Changing
	// this updates the egress_policy_id of an existing firewall group.
	EgressPolicyID *string `json:"egressPolicyId,omitempty" tf:"egress_policy_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ingress policy resource id for the firewall group. Changing
	// this updates the ingress_policy_id of an existing firewall group.
	IngressPolicyID *string `json:"ingressPolicyId,omitempty" tf:"ingress_policy_id,omitempty"`

	// A name for the firewall group. Changing this
	// updates the name of an existing firewall group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Port(s) to associate this firewall group instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall group.
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the floating IP. Required if admin wants
	// to create a firewall group for another tenant. Changing this creates a new
	// firewall group.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type FirewallGroupV2Parameters struct {

	// Administrative up/down status for the firewall group
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the admin_state_up of an existing firewall group.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// A description for the firewall group. Changing this
	// updates the description of an existing firewall group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The egress policy resource id for the firewall group. Changing
	// this updates the egress_policy_id of an existing firewall group.
	// +kubebuilder:validation:Optional
	EgressPolicyID *string `json:"egressPolicyId,omitempty" tf:"egress_policy_id,omitempty"`

	// The ingress policy resource id for the firewall group. Changing
	// this updates the ingress_policy_id of an existing firewall group.
	// +kubebuilder:validation:Optional
	IngressPolicyID *string `json:"ingressPolicyId,omitempty" tf:"ingress_policy_id,omitempty"`

	// A name for the firewall group. Changing this
	// updates the name of an existing firewall group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Port(s) to associate this firewall group instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall group.
	// +kubebuilder:validation:Optional
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the floating IP. Required if admin wants
	// to create a firewall group for another tenant. Changing this creates a new
	// firewall group.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// FirewallGroupV2Spec defines the desired state of FirewallGroupV2
type FirewallGroupV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallGroupV2Parameters `json:"forProvider"`
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
	InitProvider FirewallGroupV2InitParameters `json:"initProvider,omitempty"`
}

// FirewallGroupV2Status defines the observed state of FirewallGroupV2.
type FirewallGroupV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallGroupV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FirewallGroupV2 is the Schema for the FirewallGroupV2s API. Manages a VPC Firewall Group resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type FirewallGroupV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallGroupV2Spec   `json:"spec"`
	Status            FirewallGroupV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallGroupV2List contains a list of FirewallGroupV2s
type FirewallGroupV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallGroupV2 `json:"items"`
}

// Repository type metadata.
var (
	FirewallGroupV2_Kind             = "FirewallGroupV2"
	FirewallGroupV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallGroupV2_Kind}.String()
	FirewallGroupV2_KindAPIVersion   = FirewallGroupV2_Kind + "." + CRDGroupVersion.String()
	FirewallGroupV2_GroupVersionKind = CRDGroupVersion.WithKind(FirewallGroupV2_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallGroupV2{}, &FirewallGroupV2List{})
}
