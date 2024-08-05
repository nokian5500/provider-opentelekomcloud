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

type DedicatedReferenceTableV1InitParameters struct {

	// The conditions of the reference table. The maximum length is 30. The maximum length of
	// condition is 2048 characters.
	// schema: Required
	Conditions []*string `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The description of the reference table. The maximum length is 128 characters.
	// Currently, could be set only on update.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
	// maximum length is 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the WAF reference table resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new reference table.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of the reference table, The options are url, user-agent, ip,
	// params, cookie, referer and header. Changing this setting will push a new reference table.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DedicatedReferenceTableV1Observation struct {

	// The conditions of the reference table. The maximum length is 30. The maximum length of
	// condition is 2048 characters.
	// schema: Required
	Conditions []*string `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The time when reference table was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The description of the reference table. The maximum length is 128 characters.
	// Currently, could be set only on update.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the reference table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
	// maximum length is 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the WAF reference table resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new reference table.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of the reference table, The options are url, user-agent, ip,
	// params, cookie, referer and header. Changing this setting will push a new reference table.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DedicatedReferenceTableV1Parameters struct {

	// The conditions of the reference table. The maximum length is 30. The maximum length of
	// condition is 2048 characters.
	// schema: Required
	// +kubebuilder:validation:Optional
	Conditions []*string `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The description of the reference table. The maximum length is 128 characters.
	// Currently, could be set only on update.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
	// maximum length is 64 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the WAF reference table resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new reference table.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of the reference table, The options are url, user-agent, ip,
	// params, cookie, referer and header. Changing this setting will push a new reference table.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DedicatedReferenceTableV1Spec defines the desired state of DedicatedReferenceTableV1
type DedicatedReferenceTableV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedReferenceTableV1Parameters `json:"forProvider"`
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
	InitProvider DedicatedReferenceTableV1InitParameters `json:"initProvider,omitempty"`
}

// DedicatedReferenceTableV1Status defines the observed state of DedicatedReferenceTableV1.
type DedicatedReferenceTableV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedReferenceTableV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedReferenceTableV1 is the Schema for the DedicatedReferenceTableV1s API. Manages a WAF Dedicated Reference Table resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type DedicatedReferenceTableV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   DedicatedReferenceTableV1Spec   `json:"spec"`
	Status DedicatedReferenceTableV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedReferenceTableV1List contains a list of DedicatedReferenceTableV1s
type DedicatedReferenceTableV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedReferenceTableV1 `json:"items"`
}

// Repository type metadata.
var (
	DedicatedReferenceTableV1_Kind             = "DedicatedReferenceTableV1"
	DedicatedReferenceTableV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedReferenceTableV1_Kind}.String()
	DedicatedReferenceTableV1_KindAPIVersion   = DedicatedReferenceTableV1_Kind + "." + CRDGroupVersion.String()
	DedicatedReferenceTableV1_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedReferenceTableV1_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedReferenceTableV1{}, &DedicatedReferenceTableV1List{})
}
