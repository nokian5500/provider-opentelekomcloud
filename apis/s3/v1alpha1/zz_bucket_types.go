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

type BucketInitParameters struct {

	// The canned ACL
	// to apply. Defaults to private.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The ARN of the bucket. Will be of format arn:aws:s3:::bucketname.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Creates a unique bucket name beginning with the specified prefix.
	// Conflicts with bucket.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRule []CorsRuleInitParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// A boolean that indicates all objects should be deleted from the bucket
	// so that the bucket can be destroyed without error. These objects are not recoverable.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The Route 53 Hosted Zone ID
	// for this bucket's region.
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// A configuration of object lifecycle management
	// (documented below). The website object supports the following:
	LifecycleRule []LifecycleRuleInitParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A settings of bucket logging (documented below).
	Logging []LoggingInitParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// A valid bucket policy
	// JSON document.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A mapping of tags to assign to the bucket.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A state of versioning (documented below)
	Versioning []VersioningInitParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A website object (documented below).
	Website []WebsiteInitParameters `json:"website,omitempty" tf:"website,omitempty"`

	// The domain of the website endpoint, if the bucket is configured with a website. If not,
	// this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type BucketObservation struct {

	// The canned ACL
	// to apply. Defaults to private.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The ARN of the bucket. Will be of format arn:aws:s3:::bucketname.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The bucket domain name. Will be of format bucketname.s3.amazonaws.com.
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`

	// Creates a unique bucket name beginning with the specified prefix.
	// Conflicts with bucket.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRule []CorsRuleObservation `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// A boolean that indicates all objects should be deleted from the bucket
	// so that the bucket can be destroyed without error. These objects are not recoverable.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The Route 53 Hosted Zone ID
	// for this bucket's region.
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// Unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A configuration of object lifecycle management
	// (documented below). The website object supports the following:
	LifecycleRule []LifecycleRuleObservation `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A settings of bucket logging (documented below).
	Logging []LoggingObservation `json:"logging,omitempty" tf:"logging,omitempty"`

	// A valid bucket policy
	// JSON document.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A mapping of tags to assign to the bucket.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A state of versioning (documented below)
	Versioning []VersioningObservation `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A website object (documented below).
	Website []WebsiteObservation `json:"website,omitempty" tf:"website,omitempty"`

	// The domain of the website endpoint, if the bucket is configured with a website. If not,
	// this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type BucketParameters struct {

	// The canned ACL
	// to apply. Defaults to private.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The ARN of the bucket. Will be of format arn:aws:s3:::bucketname.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the bucket.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Creates a unique bucket name beginning with the specified prefix.
	// Conflicts with bucket.
	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// A boolean that indicates all objects should be deleted from the bucket
	// so that the bucket can be destroyed without error. These objects are not recoverable.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The Route 53 Hosted Zone ID
	// for this bucket's region.
	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// A configuration of object lifecycle management
	// (documented below). The website object supports the following:
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A settings of bucket logging (documented below).
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// A valid bucket policy
	// JSON document.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A mapping of tags to assign to the bucket.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A state of versioning (documented below)
	// +kubebuilder:validation:Optional
	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A website object (documented below).
	// +kubebuilder:validation:Optional
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`

	// The domain of the website endpoint, if the bucket is configured with a website. If not,
	// this will be an empty string. This is used to create Route 53 alias records.
	// +kubebuilder:validation:Optional
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	// +kubebuilder:validation:Optional
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type CorsRuleInitParameters struct {

	// Specifies which headers are allowed.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Specifies which origins are allowed.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// Specifies expose header in the response.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleObservation struct {

	// Specifies which headers are allowed.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Specifies which origins are allowed.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// Specifies expose header in the response.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleParameters struct {

	// Specifies which headers are allowed.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.
	// +kubebuilder:validation:Optional
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// Specifies which origins are allowed.
	// +kubebuilder:validation:Optional
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Specifies expose header in the response.
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type ExpirationInitParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket),
	// you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type ExpirationObservation struct {

	// Specifies the date after which you want the corresponding action to take effect.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket),
	// you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type ExpirationParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket),
	// you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers.
	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type LifecycleRuleInitParameters struct {

	// Specifies the number of days after initiating
	// a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state.
	// You can, however, suspend versioning on that bucket. If omitted, during bucket creation it will be in Disabled state.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire (documented below).
	Expiration []ExpirationInitParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when noncurrent object versions expire (documented below).
	NoncurrentVersionExpiration []NoncurrentVersionExpirationInitParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type LifecycleRuleObservation struct {

	// Specifies the number of days after initiating
	// a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state.
	// You can, however, suspend versioning on that bucket. If omitted, during bucket creation it will be in Disabled state.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire (documented below).
	Expiration []ExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when noncurrent object versions expire (documented below).
	NoncurrentVersionExpiration []NoncurrentVersionExpirationObservation `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type LifecycleRuleParameters struct {

	// Specifies the number of days after initiating
	// a multipart upload when the multipart upload must be completed.
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state.
	// You can, however, suspend versioning on that bucket. If omitted, during bucket creation it will be in Disabled state.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire (documented below).
	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when noncurrent object versions expire (documented below).
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type LoggingInitParameters struct {

	// The name of the bucket that will receive the log objects.
	TargetBucket *string `json:"targetBucket,omitempty" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type LoggingObservation struct {

	// The name of the bucket that will receive the log objects.
	TargetBucket *string `json:"targetBucket,omitempty" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type LoggingParameters struct {

	// The name of the bucket that will receive the log objects.
	// +kubebuilder:validation:Optional
	TargetBucket *string `json:"targetBucket" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	// +kubebuilder:validation:Optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type NoncurrentVersionExpirationInitParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionExpirationParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type VersioningInitParameters struct {

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state.
	// You can, however, suspend versioning on that bucket. If omitted, during bucket creation it will be in Disabled state.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable MFA delete for either Change the versioning state of your bucket or
	// Permanently delete an object version. Default is false.
	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`
}

type VersioningObservation struct {

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state.
	// You can, however, suspend versioning on that bucket. If omitted, during bucket creation it will be in Disabled state.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable MFA delete for either Change the versioning state of your bucket or
	// Permanently delete an object version. Default is false.
	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`
}

type VersioningParameters struct {

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state.
	// You can, however, suspend versioning on that bucket. If omitted, during bucket creation it will be in Disabled state.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable MFA delete for either Change the versioning state of your bucket or
	// Permanently delete an object version. Default is false.
	// +kubebuilder:validation:Optional
	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`
}

type WebsiteInitParameters struct {

	// An absolute path to the document to return in case of a 4XX error.
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Amazon S3 returns this index document when
	// requests are made to the root domain or any of the subfolders.
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to.
	// Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting
	// requests. The default is the protocol that is used in the original request.
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A json array containing routing rules
	// describing redirect behavior and when redirects are applied.
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type WebsiteObservation struct {

	// An absolute path to the document to return in case of a 4XX error.
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Amazon S3 returns this index document when
	// requests are made to the root domain or any of the subfolders.
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to.
	// Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting
	// requests. The default is the protocol that is used in the original request.
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A json array containing routing rules
	// describing redirect behavior and when redirects are applied.
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type WebsiteParameters struct {

	// An absolute path to the document to return in case of a 4XX error.
	// +kubebuilder:validation:Optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Amazon S3 returns this index document when
	// requests are made to the root domain or any of the subfolders.
	// +kubebuilder:validation:Optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to.
	// Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting
	// requests. The default is the protocol that is used in the original request.
	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A json array containing routing rules
	// describing redirect behavior and when redirects are applied.
	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
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
	InitProvider BucketInitParameters `json:"initProvider,omitempty"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API. Manages an S3 Bucket resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
