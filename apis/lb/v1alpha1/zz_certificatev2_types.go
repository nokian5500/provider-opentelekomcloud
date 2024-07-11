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

type CertificateV2InitParameters struct {

	// The public encrypted key of the Certificate, PEM format.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Human-readable description for the Certificate.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain of the Certificate.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Human-readable name for the Certificate. Does not have
	// to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The private encrypted key of the Certificate, PEM format.
	// Required for certificates of type server.
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB certificate. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// LB certificate.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of certificate the container holds. Either server or client.
	// Defaults to server if not set. Changing this creates a new LB certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CertificateV2Observation struct {

	// The public encrypted key of the Certificate, PEM format.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Indicates the creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Human-readable description for the Certificate.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain of the Certificate.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Indicates certificate expiration time.
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Human-readable name for the Certificate. Does not have
	// to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The private encrypted key of the Certificate, PEM format.
	// Required for certificates of type server.
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB certificate. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// LB certificate.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of certificate the container holds. Either server or client.
	// Defaults to server if not set. Changing this creates a new LB certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Indicates the update time.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CertificateV2Parameters struct {

	// The public encrypted key of the Certificate, PEM format.
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Human-readable description for the Certificate.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain of the Certificate.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Human-readable name for the Certificate. Does not have
	// to be unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The private encrypted key of the Certificate, PEM format.
	// Required for certificates of type server.
	// +kubebuilder:validation:Optional
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB certificate. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// LB certificate.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of certificate the container holds. Either server or client.
	// Defaults to server if not set. Changing this creates a new LB certificate.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// CertificateV2Spec defines the desired state of CertificateV2
type CertificateV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateV2Parameters `json:"forProvider"`
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
	InitProvider CertificateV2InitParameters `json:"initProvider,omitempty"`
}

// CertificateV2Status defines the observed state of CertificateV2.
type CertificateV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateV2 is the Schema for the CertificateV2s API. Manages a ELB Certificate resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type CertificateV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificate) || (has(self.initProvider) && has(self.initProvider.certificate))",message="spec.forProvider.certificate is a required parameter"
	Spec   CertificateV2Spec   `json:"spec"`
	Status CertificateV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateV2List contains a list of CertificateV2s
type CertificateV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateV2 `json:"items"`
}

// Repository type metadata.
var (
	CertificateV2_Kind             = "CertificateV2"
	CertificateV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateV2_Kind}.String()
	CertificateV2_KindAPIVersion   = CertificateV2_Kind + "." + CRDGroupVersion.String()
	CertificateV2_GroupVersionKind = CRDGroupVersion.WithKind(CertificateV2_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateV2{}, &CertificateV2List{})
}
