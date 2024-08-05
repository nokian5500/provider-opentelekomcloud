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

type DataImageV2InitParameters struct {

	// The master key used for encrypting an image.
	// Changing this creates a new image.
	CmkID *string `json:"cmkId,omitempty" tf:"cmk_id,omitempty"`

	// A description of the image. Changing this creates a new image.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The URL of the external image file in the OBS bucket.
	// This parameter is mandatory when you create a private image from an external file
	// uploaded to an OBS bucket. The format is OBS bucket name:Image file name.
	// Changing this creates a new image.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// The minimum size of the system disk in the unit of GB.
	// This parameter is mandatory when you create a private image from an external file
	// uploaded to an OBS bucket. The value ranges from 1 GB to 1024 GB.
	// Changing this creates a new image.
	MinDisk *float64 `json:"minDisk,omitempty" tf:"min_disk,omitempty"`

	// The name of the image.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The OS type. It can only be Windows or Linux.
	// This parameter is valid when you create a private image from an external file
	// uploaded to an OBS bucket. Changing this creates a new image.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The tags of the image.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the ECS atatched volume that needs to be converted into an image.
	// This parameter is mandatory when you create a privete image from an ECS.
	// Changing this creates a new image.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type DataImageV2Observation struct {

	// The master key used for encrypting an image.
	// Changing this creates a new image.
	CmkID *string `json:"cmkId,omitempty" tf:"cmk_id,omitempty"`

	// The image resource. The pattern can be 'instance,instance_id' or 'file,image_url'.
	DataOrigin *string `json:"dataOrigin,omitempty" tf:"data_origin,omitempty"`

	// A description of the image. Changing this creates a new image.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The image file format. The value can be vhd, zvhd, raw, zvhd2, or qcow2.
	DiskFormat *string `json:"diskFormat,omitempty" tf:"disk_format,omitempty"`

	// A unique ID assigned by IMS.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The size(bytes) of the image file format.
	ImageSize *string `json:"imageSize,omitempty" tf:"image_size,omitempty"`

	// The URL of the external image file in the OBS bucket.
	// This parameter is mandatory when you create a private image from an external file
	// uploaded to an OBS bucket. The format is OBS bucket name:Image file name.
	// Changing this creates a new image.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// The minimum size of the system disk in the unit of GB.
	// This parameter is mandatory when you create a private image from an external file
	// uploaded to an OBS bucket. The value ranges from 1 GB to 1024 GB.
	// Changing this creates a new image.
	MinDisk *float64 `json:"minDisk,omitempty" tf:"min_disk,omitempty"`

	// The name of the image.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The OS type. It can only be Windows or Linux.
	// This parameter is valid when you create a private image from an external file
	// uploaded to an OBS bucket. Changing this creates a new image.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The tags of the image.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether the image is visible to other tenants.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// The ID of the ECS atatched volume that needs to be converted into an image.
	// This parameter is mandatory when you create a privete image from an ECS.
	// Changing this creates a new image.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type DataImageV2Parameters struct {

	// The master key used for encrypting an image.
	// Changing this creates a new image.
	// +kubebuilder:validation:Optional
	CmkID *string `json:"cmkId,omitempty" tf:"cmk_id,omitempty"`

	// A description of the image. Changing this creates a new image.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The URL of the external image file in the OBS bucket.
	// This parameter is mandatory when you create a private image from an external file
	// uploaded to an OBS bucket. The format is OBS bucket name:Image file name.
	// Changing this creates a new image.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// The minimum size of the system disk in the unit of GB.
	// This parameter is mandatory when you create a private image from an external file
	// uploaded to an OBS bucket. The value ranges from 1 GB to 1024 GB.
	// Changing this creates a new image.
	// +kubebuilder:validation:Optional
	MinDisk *float64 `json:"minDisk,omitempty" tf:"min_disk,omitempty"`

	// The name of the image.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The OS type. It can only be Windows or Linux.
	// This parameter is valid when you create a private image from an external file
	// uploaded to an OBS bucket. Changing this creates a new image.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The tags of the image.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the ECS atatched volume that needs to be converted into an image.
	// This parameter is mandatory when you create a privete image from an ECS.
	// Changing this creates a new image.
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

// DataImageV2Spec defines the desired state of DataImageV2
type DataImageV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataImageV2Parameters `json:"forProvider"`
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
	InitProvider DataImageV2InitParameters `json:"initProvider,omitempty"`
}

// DataImageV2Status defines the observed state of DataImageV2.
type DataImageV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataImageV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DataImageV2 is the Schema for the DataImageV2s API. Manages a IMS Data Image resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type DataImageV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DataImageV2Spec   `json:"spec"`
	Status DataImageV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataImageV2List contains a list of DataImageV2s
type DataImageV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataImageV2 `json:"items"`
}

// Repository type metadata.
var (
	DataImageV2_Kind             = "DataImageV2"
	DataImageV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataImageV2_Kind}.String()
	DataImageV2_KindAPIVersion   = DataImageV2_Kind + "." + CRDGroupVersion.String()
	DataImageV2_GroupVersionKind = CRDGroupVersion.WithKind(DataImageV2_Kind)
)

func init() {
	SchemeBuilder.Register(&DataImageV2{}, &DataImageV2List{})
}
