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

type StreamV2InitParameters struct {

	// Maximum number of partitions for automatic scale-up when auto-scaling is enabled.
	AutoScaleMaxPartitionCount *float64 `json:"autoScaleMaxPartitionCount,omitempty" tf:"auto_scale_max_partition_count,omitempty"`

	// Minimum number of partitions for automatic scale-down
	// when auto-scaling is enabled. Minimum: 1.
	AutoScaleMinPartitionCount *float64 `json:"autoScaleMinPartitionCount,omitempty" tf:"auto_scale_min_partition_count,omitempty"`

	// Data compression type. The following types are available:
	// snappy, gzip, zip. Data is not compressed by default.
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// Source data type.
	// BLOB: a collection of binary data stored as a single entity in a database management system.
	// Default value: BLOB.
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Name of the stream. The stream name can contain 1 to 64 characters,
	// including letters, digits, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of partitions. Partitions are the base throughput unit of a DIS stream.
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Period of time for which data is retained in the stream.
	// Value range: 24-72 Unit: hour. If this parameter is left blank, the default value is used.
	// Maximum: 72
	// Default: 24
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// Stream type.
	StreamType *string `json:"streamType,omitempty" tf:"stream_type,omitempty"`

	// Tags key/value pairs to associate with the instance.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StreamV2Observation struct {

	// Maximum number of partitions for automatic scale-up when auto-scaling is enabled.
	AutoScaleMaxPartitionCount *float64 `json:"autoScaleMaxPartitionCount,omitempty" tf:"auto_scale_max_partition_count,omitempty"`

	// Minimum number of partitions for automatic scale-down
	// when auto-scaling is enabled. Minimum: 1.
	AutoScaleMinPartitionCount *float64 `json:"autoScaleMinPartitionCount,omitempty" tf:"auto_scale_min_partition_count,omitempty"`

	// Data compression type. The following types are available:
	// snappy, gzip, zip. Data is not compressed by default.
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// Time when the stream is created. The value is a timestamp.
	Created *float64 `json:"created,omitempty" tf:"created,omitempty"`

	// Source data type.
	// BLOB: a collection of binary data stored as a single entity in a database management system.
	// Default value: BLOB.
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// : Unique identifier of the partition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the stream. The stream name can contain 1 to 64 characters,
	// including letters, digits, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of partitions. Partitions are the base throughput unit of a DIS stream.
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Stream partitions details.
	Partitions []StreamV2PartitionsObservation `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// Total number of readable partitions (including partitions in ACTIVE and DELETED state).
	ReadablePartitionCount *float64 `json:"readablePartitionCount,omitempty" tf:"readable_partition_count,omitempty"`

	// Period of time for which data is retained in the stream.
	// Value range: 24-72 Unit: hour. If this parameter is left blank, the default value is used.
	// Maximum: 72
	// Default: 24
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// Current status of the stream, can be:
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Unique identifier of the stream.
	StreamID *string `json:"streamId,omitempty" tf:"stream_id,omitempty"`

	// Stream type.
	StreamType *string `json:"streamType,omitempty" tf:"stream_type,omitempty"`

	// Tags key/value pairs to associate with the instance.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Total number of writable partitions (including partitions in ACTIVE state only).
	WritablePartitionCount *float64 `json:"writablePartitionCount,omitempty" tf:"writable_partition_count,omitempty"`
}

type StreamV2Parameters struct {

	// Maximum number of partitions for automatic scale-up when auto-scaling is enabled.
	// +kubebuilder:validation:Optional
	AutoScaleMaxPartitionCount *float64 `json:"autoScaleMaxPartitionCount,omitempty" tf:"auto_scale_max_partition_count,omitempty"`

	// Minimum number of partitions for automatic scale-down
	// when auto-scaling is enabled. Minimum: 1.
	// +kubebuilder:validation:Optional
	AutoScaleMinPartitionCount *float64 `json:"autoScaleMinPartitionCount,omitempty" tf:"auto_scale_min_partition_count,omitempty"`

	// Data compression type. The following types are available:
	// snappy, gzip, zip. Data is not compressed by default.
	// +kubebuilder:validation:Optional
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// Source data type.
	// BLOB: a collection of binary data stored as a single entity in a database management system.
	// Default value: BLOB.
	// +kubebuilder:validation:Optional
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Name of the stream. The stream name can contain 1 to 64 characters,
	// including letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of partitions. Partitions are the base throughput unit of a DIS stream.
	// +kubebuilder:validation:Optional
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Period of time for which data is retained in the stream.
	// Value range: 24-72 Unit: hour. If this parameter is left blank, the default value is used.
	// Maximum: 72
	// Default: 24
	// +kubebuilder:validation:Optional
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// Stream type.
	// +kubebuilder:validation:Optional
	StreamType *string `json:"streamType,omitempty" tf:"stream_type,omitempty"`

	// Tags key/value pairs to associate with the instance.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StreamV2PartitionsInitParameters struct {
}

type StreamV2PartitionsObservation struct {

	// : Possible value range of the hash key used by the partition.
	HashRange *string `json:"hashRange,omitempty" tf:"hash_range,omitempty"`

	// : Unique identifier of the partition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// : Parent partition.
	ParentPartitions *string `json:"parentPartitions,omitempty" tf:"parent_partitions,omitempty"`

	// : Sequence number range of the partition.
	SequenceNumberRange *string `json:"sequenceNumberRange,omitempty" tf:"sequence_number_range,omitempty"`

	// Current status of the stream, can be:
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type StreamV2PartitionsParameters struct {
}

// StreamV2Spec defines the desired state of StreamV2
type StreamV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamV2Parameters `json:"forProvider"`
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
	InitProvider StreamV2InitParameters `json:"initProvider,omitempty"`
}

// StreamV2Status defines the observed state of StreamV2.
type StreamV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StreamV2 is the Schema for the StreamV2s API. Manages a DIS Stream resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type StreamV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partitionCount) || (has(self.initProvider) && has(self.initProvider.partitionCount))",message="spec.forProvider.partitionCount is a required parameter"
	Spec   StreamV2Spec   `json:"spec"`
	Status StreamV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamV2List contains a list of StreamV2s
type StreamV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamV2 `json:"items"`
}

// Repository type metadata.
var (
	StreamV2_Kind             = "StreamV2"
	StreamV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamV2_Kind}.String()
	StreamV2_KindAPIVersion   = StreamV2_Kind + "." + CRDGroupVersion.String()
	StreamV2_GroupVersionKind = CRDGroupVersion.WithKind(StreamV2_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamV2{}, &StreamV2List{})
}
