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

type IPListInitParameters struct {

	// Provides remarks about the IP address group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the IP addresses in the IP address group.
	// IPv6 is unsupported. The value cannot be an IPv6 address.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`
}

type IPListObservation struct {

	// Provides remarks about the IP address group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the IP addresses in the IP address group.
	// IPv6 is unsupported. The value cannot be an IPv6 address.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`
}

type IPListParameters struct {

	// Provides remarks about the IP address group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the IP addresses in the IP address group.
	// IPv6 is unsupported. The value cannot be an IPv6 address.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip" tf:"ip,omitempty"`
}

type IpgroupV3InitParameters struct {

	// Provides supplementary information about the IP address group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the IP addresses or CIDR blocks in the IP address group.
	// Any IP address can be used if this block isn't specified.
	IPList []IPListInitParameters `json:"ipList,omitempty" tf:"ip_list,omitempty"`

	// Specifies the IP address group name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the project ID of the IP address group.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type IpgroupV3Observation struct {

	// Indicates the creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Provides supplementary information about the IP address group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the IP addresses or CIDR blocks in the IP address group.
	// Any IP address can be used if this block isn't specified.
	IPList []IPListObservation `json:"ipList,omitempty" tf:"ip_list,omitempty"`

	// Lists the IDs of listeners with which the IP address group is associated.
	// +listType=set
	Listeners []*string `json:"listeners,omitempty" tf:"listeners,omitempty"`

	// Specifies the IP address group name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the project ID of the IP address group.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Indicates the update time.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type IpgroupV3Parameters struct {

	// Provides supplementary information about the IP address group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the IP addresses or CIDR blocks in the IP address group.
	// Any IP address can be used if this block isn't specified.
	// +kubebuilder:validation:Optional
	IPList []IPListParameters `json:"ipList,omitempty" tf:"ip_list,omitempty"`

	// Specifies the IP address group name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the project ID of the IP address group.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// IpgroupV3Spec defines the desired state of IpgroupV3
type IpgroupV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpgroupV3Parameters `json:"forProvider"`
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
	InitProvider IpgroupV3InitParameters `json:"initProvider,omitempty"`
}

// IpgroupV3Status defines the observed state of IpgroupV3.
type IpgroupV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpgroupV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IpgroupV3 is the Schema for the IpgroupV3s API. Manages a LB IpGroup resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type IpgroupV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpgroupV3Spec   `json:"spec"`
	Status            IpgroupV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpgroupV3List contains a list of IpgroupV3s
type IpgroupV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpgroupV3 `json:"items"`
}

// Repository type metadata.
var (
	IpgroupV3_Kind             = "IpgroupV3"
	IpgroupV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpgroupV3_Kind}.String()
	IpgroupV3_KindAPIVersion   = IpgroupV3_Kind + "." + CRDGroupVersion.String()
	IpgroupV3_GroupVersionKind = CRDGroupVersion.WithKind(IpgroupV3_Kind)
)

func init() {
	SchemeBuilder.Register(&IpgroupV3{}, &IpgroupV3List{})
}
