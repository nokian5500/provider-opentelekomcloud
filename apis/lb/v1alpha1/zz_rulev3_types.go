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

type ConditionsInitParameters struct {

	// Specifies the key of match item.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionsObservation struct {

	// Specifies the key of match item.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionsParameters struct {

	// Specifies the key of match item.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type RuleV3InitParameters struct {

	// - Specifies how requests are matched with the domain name or URL.
	// The values can be: EQUAL_TO, REGEX, STARTS_WITH.
	CompareType *string `json:"compareType,omitempty" tf:"compare_type,omitempty"`

	// Specifies the matching conditions of the forwarding rule.
	// This parameter is available only when advanced_forwarding is set to true.
	// Not available in eu-nl.
	Conditions []ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// ID of the policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the Policy. Only administrative users can specify a tenant UUID other than
	// their own. Changing this creates a new Policy.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Specifies the match content. The value can be one of the following: HOST_NAME, PATH.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RuleV3Observation struct {

	// - Specifies how requests are matched with the domain name or URL.
	// The values can be: EQUAL_TO, REGEX, STARTS_WITH.
	CompareType *string `json:"compareType,omitempty" tf:"compare_type,omitempty"`

	// Specifies the matching conditions of the forwarding rule.
	// This parameter is available only when advanced_forwarding is set to true.
	// Not available in eu-nl.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the Policy. Only administrative users can specify a tenant UUID other than
	// their own. Changing this creates a new Policy.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the policy rule.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Specifies the match content. The value can be one of the following: HOST_NAME, PATH.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RuleV3Parameters struct {

	// - Specifies how requests are matched with the domain name or URL.
	// The values can be: EQUAL_TO, REGEX, STARTS_WITH.
	// +kubebuilder:validation:Optional
	CompareType *string `json:"compareType,omitempty" tf:"compare_type,omitempty"`

	// Specifies the matching conditions of the forwarding rule.
	// This parameter is available only when advanced_forwarding is set to true.
	// Not available in eu-nl.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// ID of the policy.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the Policy. Only administrative users can specify a tenant UUID other than
	// their own. Changing this creates a new Policy.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Specifies the match content. The value can be one of the following: HOST_NAME, PATH.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// RuleV3Spec defines the desired state of RuleV3
type RuleV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleV3Parameters `json:"forProvider"`
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
	InitProvider RuleV3InitParameters `json:"initProvider,omitempty"`
}

// RuleV3Status defines the observed state of RuleV3.
type RuleV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RuleV3 is the Schema for the RuleV3s API. Manages a LB Rule resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type RuleV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.compareType) || (has(self.initProvider) && has(self.initProvider.compareType))",message="spec.forProvider.compareType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyId) || (has(self.initProvider) && has(self.initProvider.policyId))",message="spec.forProvider.policyId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   RuleV3Spec   `json:"spec"`
	Status RuleV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleV3List contains a list of RuleV3s
type RuleV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleV3 `json:"items"`
}

// Repository type metadata.
var (
	RuleV3_Kind             = "RuleV3"
	RuleV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleV3_Kind}.String()
	RuleV3_KindAPIVersion   = RuleV3_Kind + "." + CRDGroupVersion.String()
	RuleV3_GroupVersionKind = CRDGroupVersion.WithKind(RuleV3_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleV3{}, &RuleV3List{})
}
