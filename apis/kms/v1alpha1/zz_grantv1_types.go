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

type GrantV1InitParameters struct {

	// Indicates the ID of the authorized user.
	// Changing this creates new grant.
	GranteePrincipal *string `json:"granteePrincipal,omitempty" tf:"grantee_principal,omitempty"`

	// Indicates the ID of the KMS. Changing this creates new grant.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Name of a grant which can be 1 to 255 characters in length
	// and matches the regular expression ^[a-zA-Z0-9:/_-]{1,255}$.
	// Changing this creates new grant.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Permissions that can be granted.
	// The valid values are: create-datakey, create-datakey-without-plaintext,
	// encrypt-datakey, decrypt-datakey, describe-key, create-grant, retire-grant.
	// Changing this creates new grant.
	// +listType=set
	Operations []*string `json:"operations,omitempty" tf:"operations,omitempty"`

	// Indicates the ID of the retiring user.
	// Changing this creates new grant.
	RetiringPrincipal *string `json:"retiringPrincipal,omitempty" tf:"retiring_principal,omitempty"`
}

type GrantV1Observation struct {

	// Creation time. The value is a timestamp expressed in the number of
	// seconds since 00:00:00 UTC on January 1, 1970.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// Indicates the ID of the authorized user.
	// Changing this creates new grant.
	GranteePrincipal *string `json:"granteePrincipal,omitempty" tf:"grantee_principal,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the ID of the user who created the grant.
	IssuingPrincipal *string `json:"issuingPrincipal,omitempty" tf:"issuing_principal,omitempty"`

	// Indicates the ID of the KMS. Changing this creates new grant.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Name of a grant which can be 1 to 255 characters in length
	// and matches the regular expression ^[a-zA-Z0-9:/_-]{1,255}$.
	// Changing this creates new grant.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Permissions that can be granted.
	// The valid values are: create-datakey, create-datakey-without-plaintext,
	// encrypt-datakey, decrypt-datakey, describe-key, create-grant, retire-grant.
	// Changing this creates new grant.
	// +listType=set
	Operations []*string `json:"operations,omitempty" tf:"operations,omitempty"`

	// Indicates the ID of the retiring user.
	// Changing this creates new grant.
	RetiringPrincipal *string `json:"retiringPrincipal,omitempty" tf:"retiring_principal,omitempty"`
}

type GrantV1Parameters struct {

	// Indicates the ID of the authorized user.
	// Changing this creates new grant.
	// +kubebuilder:validation:Optional
	GranteePrincipal *string `json:"granteePrincipal,omitempty" tf:"grantee_principal,omitempty"`

	// Indicates the ID of the KMS. Changing this creates new grant.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Name of a grant which can be 1 to 255 characters in length
	// and matches the regular expression ^[a-zA-Z0-9:/_-]{1,255}$.
	// Changing this creates new grant.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Permissions that can be granted.
	// The valid values are: create-datakey, create-datakey-without-plaintext,
	// encrypt-datakey, decrypt-datakey, describe-key, create-grant, retire-grant.
	// Changing this creates new grant.
	// +kubebuilder:validation:Optional
	// +listType=set
	Operations []*string `json:"operations,omitempty" tf:"operations,omitempty"`

	// Indicates the ID of the retiring user.
	// Changing this creates new grant.
	// +kubebuilder:validation:Optional
	RetiringPrincipal *string `json:"retiringPrincipal,omitempty" tf:"retiring_principal,omitempty"`
}

// GrantV1Spec defines the desired state of GrantV1
type GrantV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrantV1Parameters `json:"forProvider"`
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
	InitProvider GrantV1InitParameters `json:"initProvider,omitempty"`
}

// GrantV1Status defines the observed state of GrantV1.
type GrantV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrantV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GrantV1 is the Schema for the GrantV1s API. Manages a KMS Grant resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type GrantV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.granteePrincipal) || (has(self.initProvider) && has(self.initProvider.granteePrincipal))",message="spec.forProvider.granteePrincipal is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keyId) || (has(self.initProvider) && has(self.initProvider.keyId))",message="spec.forProvider.keyId is a required parameter"
	Spec   GrantV1Spec   `json:"spec"`
	Status GrantV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrantV1List contains a list of GrantV1s
type GrantV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrantV1 `json:"items"`
}

// Repository type metadata.
var (
	GrantV1_Kind             = "GrantV1"
	GrantV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GrantV1_Kind}.String()
	GrantV1_KindAPIVersion   = GrantV1_Kind + "." + CRDGroupVersion.String()
	GrantV1_GroupVersionKind = CRDGroupVersion.WithKind(GrantV1_Kind)
)

func init() {
	SchemeBuilder.Register(&GrantV1{}, &GrantV1List{})
}
