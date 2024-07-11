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

type BucketPolicyInitParameters struct {

	// The name of the bucket to which to apply the policy.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The text of the policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type BucketPolicyObservation struct {

	// The name of the bucket to which to apply the policy.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The text of the policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type BucketPolicyParameters struct {

	// The name of the bucket to which to apply the policy.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The text of the policy.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

// BucketPolicySpec defines the desired state of BucketPolicy
type BucketPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketPolicyParameters `json:"forProvider"`
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
	InitProvider BucketPolicyInitParameters `json:"initProvider,omitempty"`
}

// BucketPolicyStatus defines the observed state of BucketPolicy.
type BucketPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketPolicy is the Schema for the BucketPolicys API. Manages a OBS Bucket Policy resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type BucketPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   BucketPolicySpec   `json:"spec"`
	Status BucketPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketPolicyList contains a list of BucketPolicys
type BucketPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketPolicy `json:"items"`
}

// Repository type metadata.
var (
	BucketPolicy_Kind             = "BucketPolicy"
	BucketPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketPolicy_Kind}.String()
	BucketPolicy_KindAPIVersion   = BucketPolicy_Kind + "." + CRDGroupVersion.String()
	BucketPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BucketPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketPolicy{}, &BucketPolicyList{})
}
