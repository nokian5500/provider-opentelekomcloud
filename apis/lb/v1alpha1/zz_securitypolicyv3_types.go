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

type SecurityPolicyV3InitParameters struct {

	// Lists the cipher suites supported by the custom security policy.
	Ciphers []*string `json:"ciphers,omitempty" tf:"ciphers,omitempty"`

	// Provides supplementary information about the security policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the security policy name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Lists the TLS protocols supported by the custom security policy.
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`
}

type SecurityPolicyV3Observation struct {

	// Lists the cipher suites supported by the custom security policy.
	Ciphers []*string `json:"ciphers,omitempty" tf:"ciphers,omitempty"`

	// The time when the custom security policy was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Provides supplementary information about the security policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique ID for the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ListenerIds []*string `json:"listenerIds,omitempty" tf:"listener_ids,omitempty"`

	// Specifies the security policy name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project ID of the custom security policy.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Lists the TLS protocols supported by the custom security policy.
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// The time when the custom security policy was updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type SecurityPolicyV3Parameters struct {

	// Lists the cipher suites supported by the custom security policy.
	// +kubebuilder:validation:Optional
	Ciphers []*string `json:"ciphers,omitempty" tf:"ciphers,omitempty"`

	// Provides supplementary information about the security policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the security policy name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Lists the TLS protocols supported by the custom security policy.
	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`
}

// SecurityPolicyV3Spec defines the desired state of SecurityPolicyV3
type SecurityPolicyV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityPolicyV3Parameters `json:"forProvider"`
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
	InitProvider SecurityPolicyV3InitParameters `json:"initProvider,omitempty"`
}

// SecurityPolicyV3Status defines the observed state of SecurityPolicyV3.
type SecurityPolicyV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityPolicyV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityPolicyV3 is the Schema for the SecurityPolicyV3s API. Manages a LB Security Policy resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type SecurityPolicyV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ciphers) || (has(self.initProvider) && has(self.initProvider.ciphers))",message="spec.forProvider.ciphers is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocols) || (has(self.initProvider) && has(self.initProvider.protocols))",message="spec.forProvider.protocols is a required parameter"
	Spec   SecurityPolicyV3Spec   `json:"spec"`
	Status SecurityPolicyV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityPolicyV3List contains a list of SecurityPolicyV3s
type SecurityPolicyV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityPolicyV3 `json:"items"`
}

// Repository type metadata.
var (
	SecurityPolicyV3_Kind             = "SecurityPolicyV3"
	SecurityPolicyV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityPolicyV3_Kind}.String()
	SecurityPolicyV3_KindAPIVersion   = SecurityPolicyV3_Kind + "." + CRDGroupVersion.String()
	SecurityPolicyV3_GroupVersionKind = CRDGroupVersion.WithKind(SecurityPolicyV3_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityPolicyV3{}, &SecurityPolicyV3List{})
}
