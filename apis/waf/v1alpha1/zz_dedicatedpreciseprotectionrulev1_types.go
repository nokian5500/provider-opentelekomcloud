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

type DedicatedPreciseProtectionRuleV1ActionInitParameters struct {

	// Action type. Changing this creates a new rule.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// ID of a known attack source rule. This parameter can be configured only when category is set to block. Changing this creates a new rule.
	FollowedActionID *string `json:"followedActionId,omitempty" tf:"followed_action_id,omitempty"`
}

type DedicatedPreciseProtectionRuleV1ActionObservation struct {

	// Action type. Changing this creates a new rule.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// ID of a known attack source rule. This parameter can be configured only when category is set to block. Changing this creates a new rule.
	FollowedActionID *string `json:"followedActionId,omitempty" tf:"followed_action_id,omitempty"`
}

type DedicatedPreciseProtectionRuleV1ActionParameters struct {

	// Action type. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	Category *string `json:"category" tf:"category,omitempty"`

	// ID of a known attack source rule. This parameter can be configured only when category is set to block. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	FollowedActionID *string `json:"followedActionId,omitempty" tf:"followed_action_id,omitempty"`
}

type DedicatedPreciseProtectionRuleV1ConditionsInitParameters struct {

	// Field type. The options are url, user-agent, ip, params, cookie, referer, header, request_line, method, and request.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Content of the conditions. This parameter is mandatory when the suffix of logic_operation is not any or all. This parameter is mandatory when the suffix of logic_operation is not any or all. Changing this creates a new rule.
	Contents []*string `json:"contents,omitempty" tf:"contents,omitempty"`

	// Subfield. Changing this creates a new rule.
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// Logic for matching the condition. Changing this creates a new rule.
	LogicOperation *string `json:"logicOperation,omitempty" tf:"logic_operation,omitempty"`

	// Reference table ID. This parameter is mandatory when the suffix of logic_operation is any or all. The reference table type must be the same as the category type. Changing this creates a new rule.
	ValueListID *string `json:"valueListId,omitempty" tf:"value_list_id,omitempty"`
}

type DedicatedPreciseProtectionRuleV1ConditionsObservation struct {

	// Field type. The options are url, user-agent, ip, params, cookie, referer, header, request_line, method, and request.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Content of the conditions. This parameter is mandatory when the suffix of logic_operation is not any or all. This parameter is mandatory when the suffix of logic_operation is not any or all. Changing this creates a new rule.
	Contents []*string `json:"contents,omitempty" tf:"contents,omitempty"`

	// Subfield. Changing this creates a new rule.
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// Logic for matching the condition. Changing this creates a new rule.
	LogicOperation *string `json:"logicOperation,omitempty" tf:"logic_operation,omitempty"`

	// Reference table ID. This parameter is mandatory when the suffix of logic_operation is any or all. The reference table type must be the same as the category type. Changing this creates a new rule.
	ValueListID *string `json:"valueListId,omitempty" tf:"value_list_id,omitempty"`
}

type DedicatedPreciseProtectionRuleV1ConditionsParameters struct {

	// Field type. The options are url, user-agent, ip, params, cookie, referer, header, request_line, method, and request.
	// +kubebuilder:validation:Optional
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Content of the conditions. This parameter is mandatory when the suffix of logic_operation is not any or all. This parameter is mandatory when the suffix of logic_operation is not any or all. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	Contents []*string `json:"contents,omitempty" tf:"contents,omitempty"`

	// Subfield. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// Logic for matching the condition. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	LogicOperation *string `json:"logicOperation,omitempty" tf:"logic_operation,omitempty"`

	// Reference table ID. This parameter is mandatory when the suffix of logic_operation is any or all. The reference table type must be the same as the category type. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	ValueListID *string `json:"valueListId,omitempty" tf:"value_list_id,omitempty"`
}

type DedicatedPreciseProtectionRuleV1InitParameters struct {

	// Protection action to take if the number of requests reaches the upper limit. Changing this creates a new rule.
	// The conditions block supports:
	Action []DedicatedPreciseProtectionRuleV1ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Match condition List. Changing this creates a new rule.
	// The conditions block supports:
	Conditions []DedicatedPreciseProtectionRuleV1ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Rule description. Changing this creates a new rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The WAF policy ID. Changing this creates a new rule.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Priority of a rule. A small value indicates a high priority. If two rules are assigned with the same priority, the rule added earlier has higher priority. Value range: 0 to 1000. Changing this creates a new rule.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Timestamp (ms) when the precise protection rule takes effect. This parameter is returned only when time is true. Changing this creates a new rule.
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`

	// Timestamp (ms) when the precise protection rule expires. This parameter is returned only when time is true. Changing this creates a new rule.
	Terminal *float64 `json:"terminal,omitempty" tf:"terminal,omitempty"`

	// Time the precise protection rule takes effect. Changing this creates a new rule.
	// Values:
	Time *bool `json:"time,omitempty" tf:"time,omitempty"`
}

type DedicatedPreciseProtectionRuleV1Observation struct {

	// Protection action to take if the number of requests reaches the upper limit. Changing this creates a new rule.
	// The conditions block supports:
	Action []DedicatedPreciseProtectionRuleV1ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// Match condition List. Changing this creates a new rule.
	// The conditions block supports:
	Conditions []DedicatedPreciseProtectionRuleV1ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Timestamp the rule is created.
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Rule description. Changing this creates a new rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The WAF policy ID. Changing this creates a new rule.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Priority of a rule. A small value indicates a high priority. If two rules are assigned with the same priority, the rule added earlier has higher priority. Value range: 0 to 1000. Changing this creates a new rule.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Timestamp (ms) when the precise protection rule takes effect. This parameter is returned only when time is true. Changing this creates a new rule.
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`

	// Rule status. The value can be:
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// Timestamp (ms) when the precise protection rule expires. This parameter is returned only when time is true. Changing this creates a new rule.
	Terminal *float64 `json:"terminal,omitempty" tf:"terminal,omitempty"`

	// Time the precise protection rule takes effect. Changing this creates a new rule.
	// Values:
	Time *bool `json:"time,omitempty" tf:"time,omitempty"`
}

type DedicatedPreciseProtectionRuleV1Parameters struct {

	// Protection action to take if the number of requests reaches the upper limit. Changing this creates a new rule.
	// The conditions block supports:
	// +kubebuilder:validation:Optional
	Action []DedicatedPreciseProtectionRuleV1ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Match condition List. Changing this creates a new rule.
	// The conditions block supports:
	// +kubebuilder:validation:Optional
	Conditions []DedicatedPreciseProtectionRuleV1ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Rule description. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The WAF policy ID. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Priority of a rule. A small value indicates a high priority. If two rules are assigned with the same priority, the rule added earlier has higher priority. Value range: 0 to 1000. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Timestamp (ms) when the precise protection rule takes effect. This parameter is returned only when time is true. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`

	// Timestamp (ms) when the precise protection rule expires. This parameter is returned only when time is true. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	Terminal *float64 `json:"terminal,omitempty" tf:"terminal,omitempty"`

	// Time the precise protection rule takes effect. Changing this creates a new rule.
	// Values:
	// +kubebuilder:validation:Optional
	Time *bool `json:"time,omitempty" tf:"time,omitempty"`
}

// DedicatedPreciseProtectionRuleV1Spec defines the desired state of DedicatedPreciseProtectionRuleV1
type DedicatedPreciseProtectionRuleV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedPreciseProtectionRuleV1Parameters `json:"forProvider"`
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
	InitProvider DedicatedPreciseProtectionRuleV1InitParameters `json:"initProvider,omitempty"`
}

// DedicatedPreciseProtectionRuleV1Status defines the observed state of DedicatedPreciseProtectionRuleV1.
type DedicatedPreciseProtectionRuleV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedPreciseProtectionRuleV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedPreciseProtectionRuleV1 is the Schema for the DedicatedPreciseProtectionRuleV1s API. Manages a WAF Dedicated Precise Protection Rule resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type DedicatedPreciseProtectionRuleV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyId) || (has(self.initProvider) && has(self.initProvider.policyId))",message="spec.forProvider.policyId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.priority) || (has(self.initProvider) && has(self.initProvider.priority))",message="spec.forProvider.priority is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.time) || (has(self.initProvider) && has(self.initProvider.time))",message="spec.forProvider.time is a required parameter"
	Spec   DedicatedPreciseProtectionRuleV1Spec   `json:"spec"`
	Status DedicatedPreciseProtectionRuleV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedPreciseProtectionRuleV1List contains a list of DedicatedPreciseProtectionRuleV1s
type DedicatedPreciseProtectionRuleV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedPreciseProtectionRuleV1 `json:"items"`
}

// Repository type metadata.
var (
	DedicatedPreciseProtectionRuleV1_Kind             = "DedicatedPreciseProtectionRuleV1"
	DedicatedPreciseProtectionRuleV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedPreciseProtectionRuleV1_Kind}.String()
	DedicatedPreciseProtectionRuleV1_KindAPIVersion   = DedicatedPreciseProtectionRuleV1_Kind + "." + CRDGroupVersion.String()
	DedicatedPreciseProtectionRuleV1_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedPreciseProtectionRuleV1_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedPreciseProtectionRuleV1{}, &DedicatedPreciseProtectionRuleV1List{})
}
