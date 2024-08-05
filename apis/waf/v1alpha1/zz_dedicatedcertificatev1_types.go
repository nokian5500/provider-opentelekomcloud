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

type DedicatedCertificateV1InitParameters struct {

	// The certificate content. Changing this creates a new certificate.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The private key. Changing this creates a new certificate.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The certificate name. The value can contain a maximum of 64 characters.
	// Only digits, letters, underscores(_), and hyphens(-) are allowed. Changing this creates a new certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DedicatedCertificateV1Observation struct {

	// The certificate content. Changing this creates a new certificate.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Date when the certificate is uploaded.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Date when the certificate expires.
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// ID of the certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The private key. Changing this creates a new certificate.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The certificate name. The value can contain a maximum of 64 characters.
	// Only digits, letters, underscores(_), and hyphens(-) are allowed. Changing this creates a new certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DedicatedCertificateV1Parameters struct {

	// The certificate content. Changing this creates a new certificate.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The private key. Changing this creates a new certificate.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The certificate name. The value can contain a maximum of 64 characters.
	// Only digits, letters, underscores(_), and hyphens(-) are allowed. Changing this creates a new certificate.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// DedicatedCertificateV1Spec defines the desired state of DedicatedCertificateV1
type DedicatedCertificateV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedCertificateV1Parameters `json:"forProvider"`
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
	InitProvider DedicatedCertificateV1InitParameters `json:"initProvider,omitempty"`
}

// DedicatedCertificateV1Status defines the observed state of DedicatedCertificateV1.
type DedicatedCertificateV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedCertificateV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedCertificateV1 is the Schema for the DedicatedCertificateV1s API. Manages a WAF Dedicated Certificate resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type DedicatedCertificateV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || (has(self.initProvider) && has(self.initProvider.content))",message="spec.forProvider.content is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DedicatedCertificateV1Spec   `json:"spec"`
	Status DedicatedCertificateV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedCertificateV1List contains a list of DedicatedCertificateV1s
type DedicatedCertificateV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedCertificateV1 `json:"items"`
}

// Repository type metadata.
var (
	DedicatedCertificateV1_Kind             = "DedicatedCertificateV1"
	DedicatedCertificateV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedCertificateV1_Kind}.String()
	DedicatedCertificateV1_KindAPIVersion   = DedicatedCertificateV1_Kind + "." + CRDGroupVersion.String()
	DedicatedCertificateV1_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedCertificateV1_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedCertificateV1{}, &DedicatedCertificateV1List{})
}
