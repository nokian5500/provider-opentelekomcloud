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

type AuthenticatingProxyInitParameters struct {

	// X509 CA certificate configured in authenticating_proxy mode. The maximum size of the certificate is 1 MB.
	CA *string `json:"ca,omitempty" tf:"ca,omitempty"`

	// Client certificate issued by the X509 CA certificate configured in authenticating_proxy mode.
	// This certificate is used for authentication from kube-apiserver to the extended API server.
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	// Private key of the client certificate issued by the X509 CA certificate configured in authenticating_proxy mode.
	// This key is used for authentication from kube-apiserver to the extended API server.
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`
}

type AuthenticatingProxyObservation struct {

	// X509 CA certificate configured in authenticating_proxy mode. The maximum size of the certificate is 1 MB.
	CA *string `json:"ca,omitempty" tf:"ca,omitempty"`

	// Client certificate issued by the X509 CA certificate configured in authenticating_proxy mode.
	// This certificate is used for authentication from kube-apiserver to the extended API server.
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	// Private key of the client certificate issued by the X509 CA certificate configured in authenticating_proxy mode.
	// This key is used for authentication from kube-apiserver to the extended API server.
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`
}

type AuthenticatingProxyParameters struct {

	// X509 CA certificate configured in authenticating_proxy mode. The maximum size of the certificate is 1 MB.
	// +kubebuilder:validation:Optional
	CA *string `json:"ca" tf:"ca,omitempty"`

	// Client certificate issued by the X509 CA certificate configured in authenticating_proxy mode.
	// This certificate is used for authentication from kube-apiserver to the extended API server.
	// +kubebuilder:validation:Optional
	Cert *string `json:"cert" tf:"cert,omitempty"`

	// Private key of the client certificate issued by the X509 CA certificate configured in authenticating_proxy mode.
	// This key is used for authentication from kube-apiserver to the extended API server.
	// +kubebuilder:validation:Optional
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`
}

type CertificateClustersInitParameters struct {
}

type CertificateClustersObservation struct {
	CertificateAuthorityData *string `json:"certificateAuthorityData,omitempty" tf:"certificate_authority_data,omitempty"`

	// Cluster name. Changing this parameter will create a new cluster resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Server *string `json:"server,omitempty" tf:"server,omitempty"`
}

type CertificateClustersParameters struct {
}

type CertificateUsersInitParameters struct {
}

type CertificateUsersObservation struct {
	ClientCertificateData *string `json:"clientCertificateData,omitempty" tf:"client_certificate_data,omitempty"`

	ClientKeyData *string `json:"clientKeyData,omitempty" tf:"client_key_data,omitempty"`

	// Cluster name. Changing this parameter will create a new cluster resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CertificateUsersParameters struct {
}

type ClusterV3InitParameters struct {

	// Cluster annotation, key/value pair format. Changing this parameter will create a new cluster resource.
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Authenticating proxy configuration. Required if authentication_mode is set to authenticating_proxy.
	AuthenticatingProxy []AuthenticatingProxyInitParameters `json:"authenticatingProxy,omitempty" tf:"authenticating_proxy,omitempty"`

	// CA root certificate provided in the authenticating_proxy mode.
	// Deprecated, use authenticating_proxy instead.
	AuthenticatingProxyCA *string `json:"authenticatingProxyCa,omitempty" tf:"authenticating_proxy_ca,omitempty"`

	// Authentication mode of the cluster, possible values are rbac and authenticating_proxy.
	// Defaults to rbac. Changing this parameter will create a new cluster resource.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// Charging mode of the cluster, which is 0 (on demand). Changing this parameter will create a new cluster resource.
	BillingMode *float64 `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

	// Cluster Type, possible values are VirtualMachine and BareMetal. Changing this parameter will create a new cluster resource.
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// For the cluster version, possible values are v1.27, v1.25, v1.23, v1.21.
	// If this parameter is not set, the cluster of the latest version is created by default.
	// Changing this parameter will create a new cluster resource. OTC-API
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// Container network segment. Changing this parameter will create a new cluster resource.
	ContainerNetworkCidr *string `json:"containerNetworkCidr,omitempty" tf:"container_network_cidr,omitempty"`

	// Container network type.
	ContainerNetworkType *string `json:"containerNetworkType,omitempty" tf:"container_network_type,omitempty"`

	// Specified whether to delete all associated network resources when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteAllNetwork *string `json:"deleteAllNetwork,omitempty" tf:"delete_all_network,omitempty"`

	// Specified whether to delete all associated storage resources when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteAllStorage *string `json:"deleteAllStorage,omitempty" tf:"delete_all_storage,omitempty"`

	// Specified whether to unbind associated SFS Turbo file systems when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteEFS *string `json:"deleteEfs,omitempty" tf:"delete_efs,omitempty"`

	// Specified whether to delete ENI ports when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteEni *string `json:"deleteEni,omitempty" tf:"delete_eni,omitempty"`

	// Specified whether to delete associated EVS disks when deleting the CCE cluster.
	// valid values are true, try and false. Default is false.
	DeleteEvs *string `json:"deleteEvs,omitempty" tf:"delete_evs,omitempty"`

	// Specified whether to delete cluster Service/ingress-related resources, such as ELB when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteNet *string `json:"deleteNet,omitempty" tf:"delete_net,omitempty"`

	// Specified whether to delete associated OBS buckets when deleting the CCE cluster.
	// valid values are true, try and false. Default is false.
	DeleteObs *string `json:"deleteObs,omitempty" tf:"delete_obs,omitempty"`

	// Specified whether to delete associated SFS file systems when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteSfs *string `json:"deleteSfs,omitempty" tf:"delete_sfs,omitempty"`

	// Cluster description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// EIP address of the cluster.
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// System and data disks encryption of master nodes. Changing this parameter will create a new cluster resource.
	EnableVolumeEncryption *bool `json:"enableVolumeEncryption,omitempty" tf:"enable_volume_encryption,omitempty"`

	// Specifies the ENI network segment. Specified when creating a CCE Turbo cluster. Changing this parameter will create a new cluster resource.
	EniSubnetCidr *string `json:"eniSubnetCidr,omitempty" tf:"eni_subnet_cidr,omitempty"`

	// -  Specifies the ENI subnet ID. Specified when creating a CCE Turbo cluster. Changing this parameter will create a new cluster resource.
	EniSubnetID *string `json:"eniSubnetId,omitempty" tf:"eni_subnet_id,omitempty"`

	// Extended parameter. Changing this parameter will create a new cluster resource.
	// List of cluster extended params.
	ExtendParam map[string]*string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Cluster specifications. Changing this parameter will create a new cluster resource.
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// The ID of the high speed network used to create bare metal nodes. Changing this parameter will create a new cluster resource.
	HighwaySubnetID *string `json:"highwaySubnetId,omitempty" tf:"highway_subnet_id,omitempty"`

	// Skip all cluster addons operations.
	IgnoreAddons *bool `json:"ignoreAddons,omitempty" tf:"ignore_addons,omitempty"`

	// Skip sensitive cluster data.
	IgnoreCertificateClustersData *bool `json:"ignoreCertificateClustersData,omitempty" tf:"ignore_certificate_clusters_data,omitempty"`

	// Skip sensitive user data.
	IgnoreCertificateUsersData *bool `json:"ignoreCertificateUsersData,omitempty" tf:"ignore_certificate_users_data,omitempty"`

	// Service forwarding mode. Two modes are available:
	KubeProxyMode *string `json:"kubeProxyMode,omitempty" tf:"kube_proxy_mode,omitempty"`

	// Service CIDR block, or the IP address range which the kubernetes
	// clusterIp must fall within. This parameter is available only for clusters of v1.11.7 and later.
	KubernetesSvcIPRange *string `json:"kubernetesSvcIpRange,omitempty" tf:"kubernetes_svc_ip_range,omitempty"`

	// Cluster tag, key/value pair format. Changing this parameter will create a new cluster resource.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Enable multiple AZs for the cluster, only when using HA flavors. Changing this parameter will create a new cluster resource.
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// Cluster name. Changing this parameter will create a new cluster resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Remove addons installed by the default after the cluster creation.
	NoAddons *bool `json:"noAddons,omitempty" tf:"no_addons,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Network ID of the subnet used to create the node. Changing this parameter will create a new cluster resource.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The ID of the VPC used to create the node. Changing this parameter will create a new cluster resource.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ClusterV3Observation struct {

	// Cluster annotation, key/value pair format. Changing this parameter will create a new cluster resource.
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Authenticating proxy configuration. Required if authentication_mode is set to authenticating_proxy.
	AuthenticatingProxy []AuthenticatingProxyObservation `json:"authenticatingProxy,omitempty" tf:"authenticating_proxy,omitempty"`

	// CA root certificate provided in the authenticating_proxy mode.
	// Deprecated, use authenticating_proxy instead.
	AuthenticatingProxyCA *string `json:"authenticatingProxyCa,omitempty" tf:"authenticating_proxy_ca,omitempty"`

	// Authentication mode of the cluster, possible values are rbac and authenticating_proxy.
	// Defaults to rbac. Changing this parameter will create a new cluster resource.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// Charging mode of the cluster, which is 0 (on demand). Changing this parameter will create a new cluster resource.
	BillingMode *float64 `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

	CertificateClusters []CertificateClustersObservation `json:"certificateClusters,omitempty" tf:"certificate_clusters,omitempty"`

	CertificateUsers []CertificateUsersObservation `json:"certificateUsers,omitempty" tf:"certificate_users,omitempty"`

	// Cluster Type, possible values are VirtualMachine and BareMetal. Changing this parameter will create a new cluster resource.
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// For the cluster version, possible values are v1.27, v1.25, v1.23, v1.21.
	// If this parameter is not set, the cluster of the latest version is created by default.
	// Changing this parameter will create a new cluster resource. OTC-API
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// Container network segment. Changing this parameter will create a new cluster resource.
	ContainerNetworkCidr *string `json:"containerNetworkCidr,omitempty" tf:"container_network_cidr,omitempty"`

	// Container network type.
	ContainerNetworkType *string `json:"containerNetworkType,omitempty" tf:"container_network_type,omitempty"`

	// Specified whether to delete all associated network resources when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteAllNetwork *string `json:"deleteAllNetwork,omitempty" tf:"delete_all_network,omitempty"`

	// Specified whether to delete all associated storage resources when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteAllStorage *string `json:"deleteAllStorage,omitempty" tf:"delete_all_storage,omitempty"`

	// Specified whether to unbind associated SFS Turbo file systems when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteEFS *string `json:"deleteEfs,omitempty" tf:"delete_efs,omitempty"`

	// Specified whether to delete ENI ports when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteEni *string `json:"deleteEni,omitempty" tf:"delete_eni,omitempty"`

	// Specified whether to delete associated EVS disks when deleting the CCE cluster.
	// valid values are true, try and false. Default is false.
	DeleteEvs *string `json:"deleteEvs,omitempty" tf:"delete_evs,omitempty"`

	// Specified whether to delete cluster Service/ingress-related resources, such as ELB when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteNet *string `json:"deleteNet,omitempty" tf:"delete_net,omitempty"`

	// Specified whether to delete associated OBS buckets when deleting the CCE cluster.
	// valid values are true, try and false. Default is false.
	DeleteObs *string `json:"deleteObs,omitempty" tf:"delete_obs,omitempty"`

	// Specified whether to delete associated SFS file systems when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	DeleteSfs *string `json:"deleteSfs,omitempty" tf:"delete_sfs,omitempty"`

	// Cluster description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// EIP address of the cluster.
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// System and data disks encryption of master nodes. Changing this parameter will create a new cluster resource.
	EnableVolumeEncryption *bool `json:"enableVolumeEncryption,omitempty" tf:"enable_volume_encryption,omitempty"`

	// Specifies the ENI network segment. Specified when creating a CCE Turbo cluster. Changing this parameter will create a new cluster resource.
	EniSubnetCidr *string `json:"eniSubnetCidr,omitempty" tf:"eni_subnet_cidr,omitempty"`

	// -  Specifies the ENI subnet ID. Specified when creating a CCE Turbo cluster. Changing this parameter will create a new cluster resource.
	EniSubnetID *string `json:"eniSubnetId,omitempty" tf:"eni_subnet_id,omitempty"`

	// Extended parameter. Changing this parameter will create a new cluster resource.
	// List of cluster extended params.
	ExtendParam map[string]*string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// The external network address.
	External *string `json:"external,omitempty" tf:"external,omitempty"`

	// The endpoint of the cluster to be accessed through API Gateway.
	ExternalOtc *string `json:"externalOtc,omitempty" tf:"external_otc,omitempty"`

	// Cluster specifications. Changing this parameter will create a new cluster resource.
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// The ID of the high speed network used to create bare metal nodes. Changing this parameter will create a new cluster resource.
	HighwaySubnetID *string `json:"highwaySubnetId,omitempty" tf:"highway_subnet_id,omitempty"`

	// ID of the cluster resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Skip all cluster addons operations.
	IgnoreAddons *bool `json:"ignoreAddons,omitempty" tf:"ignore_addons,omitempty"`

	// Skip sensitive cluster data.
	IgnoreCertificateClustersData *bool `json:"ignoreCertificateClustersData,omitempty" tf:"ignore_certificate_clusters_data,omitempty"`

	// Skip sensitive user data.
	IgnoreCertificateUsersData *bool `json:"ignoreCertificateUsersData,omitempty" tf:"ignore_certificate_users_data,omitempty"`

	// List of installed addon IDs. Empty if ignore_addons is true.
	InstalledAddons []*string `json:"installedAddons,omitempty" tf:"installed_addons,omitempty"`

	// The internal network address.
	Internal *string `json:"internal,omitempty" tf:"internal,omitempty"`

	// Service forwarding mode. Two modes are available:
	KubeProxyMode *string `json:"kubeProxyMode,omitempty" tf:"kube_proxy_mode,omitempty"`

	// Service CIDR block, or the IP address range which the kubernetes
	// clusterIp must fall within. This parameter is available only for clusters of v1.11.7 and later.
	KubernetesSvcIPRange *string `json:"kubernetesSvcIpRange,omitempty" tf:"kubernetes_svc_ip_range,omitempty"`

	// Cluster tag, key/value pair format. Changing this parameter will create a new cluster resource.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Enable multiple AZs for the cluster, only when using HA flavors. Changing this parameter will create a new cluster resource.
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// Cluster name. Changing this parameter will create a new cluster resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Remove addons installed by the default after the cluster creation.
	NoAddons *bool `json:"noAddons,omitempty" tf:"no_addons,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the autogenerated security group for the CCE master port.
	SecurityGroupControl *string `json:"securityGroupControl,omitempty" tf:"security_group_control,omitempty"`

	// ID of the autogenerated security group for the CCE nodes.
	SecurityGroupNode *string `json:"securityGroupNode,omitempty" tf:"security_group_node,omitempty"`

	// Cluster status information.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The Network ID of the subnet used to create the node. Changing this parameter will create a new cluster resource.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The ID of the VPC used to create the node. Changing this parameter will create a new cluster resource.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ClusterV3Parameters struct {

	// Cluster annotation, key/value pair format. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Authenticating proxy configuration. Required if authentication_mode is set to authenticating_proxy.
	// +kubebuilder:validation:Optional
	AuthenticatingProxy []AuthenticatingProxyParameters `json:"authenticatingProxy,omitempty" tf:"authenticating_proxy,omitempty"`

	// CA root certificate provided in the authenticating_proxy mode.
	// Deprecated, use authenticating_proxy instead.
	// +kubebuilder:validation:Optional
	AuthenticatingProxyCA *string `json:"authenticatingProxyCa,omitempty" tf:"authenticating_proxy_ca,omitempty"`

	// Authentication mode of the cluster, possible values are rbac and authenticating_proxy.
	// Defaults to rbac. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// Charging mode of the cluster, which is 0 (on demand). Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	BillingMode *float64 `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

	// Cluster Type, possible values are VirtualMachine and BareMetal. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// For the cluster version, possible values are v1.27, v1.25, v1.23, v1.21.
	// If this parameter is not set, the cluster of the latest version is created by default.
	// Changing this parameter will create a new cluster resource. OTC-API
	// +kubebuilder:validation:Optional
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// Container network segment. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	ContainerNetworkCidr *string `json:"containerNetworkCidr,omitempty" tf:"container_network_cidr,omitempty"`

	// Container network type.
	// +kubebuilder:validation:Optional
	ContainerNetworkType *string `json:"containerNetworkType,omitempty" tf:"container_network_type,omitempty"`

	// Specified whether to delete all associated network resources when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	// +kubebuilder:validation:Optional
	DeleteAllNetwork *string `json:"deleteAllNetwork,omitempty" tf:"delete_all_network,omitempty"`

	// Specified whether to delete all associated storage resources when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	// +kubebuilder:validation:Optional
	DeleteAllStorage *string `json:"deleteAllStorage,omitempty" tf:"delete_all_storage,omitempty"`

	// Specified whether to unbind associated SFS Turbo file systems when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	// +kubebuilder:validation:Optional
	DeleteEFS *string `json:"deleteEfs,omitempty" tf:"delete_efs,omitempty"`

	// Specified whether to delete ENI ports when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	// +kubebuilder:validation:Optional
	DeleteEni *string `json:"deleteEni,omitempty" tf:"delete_eni,omitempty"`

	// Specified whether to delete associated EVS disks when deleting the CCE cluster.
	// valid values are true, try and false. Default is false.
	// +kubebuilder:validation:Optional
	DeleteEvs *string `json:"deleteEvs,omitempty" tf:"delete_evs,omitempty"`

	// Specified whether to delete cluster Service/ingress-related resources, such as ELB when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	// +kubebuilder:validation:Optional
	DeleteNet *string `json:"deleteNet,omitempty" tf:"delete_net,omitempty"`

	// Specified whether to delete associated OBS buckets when deleting the CCE cluster.
	// valid values are true, try and false. Default is false.
	// +kubebuilder:validation:Optional
	DeleteObs *string `json:"deleteObs,omitempty" tf:"delete_obs,omitempty"`

	// Specified whether to delete associated SFS file systems when deleting the CCE
	// cluster. valid values are true, try and false. Default is false.
	// +kubebuilder:validation:Optional
	DeleteSfs *string `json:"deleteSfs,omitempty" tf:"delete_sfs,omitempty"`

	// Cluster description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// EIP address of the cluster.
	// +kubebuilder:validation:Optional
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// System and data disks encryption of master nodes. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	EnableVolumeEncryption *bool `json:"enableVolumeEncryption,omitempty" tf:"enable_volume_encryption,omitempty"`

	// Specifies the ENI network segment. Specified when creating a CCE Turbo cluster. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	EniSubnetCidr *string `json:"eniSubnetCidr,omitempty" tf:"eni_subnet_cidr,omitempty"`

	// -  Specifies the ENI subnet ID. Specified when creating a CCE Turbo cluster. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	EniSubnetID *string `json:"eniSubnetId,omitempty" tf:"eni_subnet_id,omitempty"`

	// Extended parameter. Changing this parameter will create a new cluster resource.
	// List of cluster extended params.
	// +kubebuilder:validation:Optional
	ExtendParam map[string]*string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Cluster specifications. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// The ID of the high speed network used to create bare metal nodes. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	HighwaySubnetID *string `json:"highwaySubnetId,omitempty" tf:"highway_subnet_id,omitempty"`

	// Skip all cluster addons operations.
	// +kubebuilder:validation:Optional
	IgnoreAddons *bool `json:"ignoreAddons,omitempty" tf:"ignore_addons,omitempty"`

	// Skip sensitive cluster data.
	// +kubebuilder:validation:Optional
	IgnoreCertificateClustersData *bool `json:"ignoreCertificateClustersData,omitempty" tf:"ignore_certificate_clusters_data,omitempty"`

	// Skip sensitive user data.
	// +kubebuilder:validation:Optional
	IgnoreCertificateUsersData *bool `json:"ignoreCertificateUsersData,omitempty" tf:"ignore_certificate_users_data,omitempty"`

	// Service forwarding mode. Two modes are available:
	// +kubebuilder:validation:Optional
	KubeProxyMode *string `json:"kubeProxyMode,omitempty" tf:"kube_proxy_mode,omitempty"`

	// Service CIDR block, or the IP address range which the kubernetes
	// clusterIp must fall within. This parameter is available only for clusters of v1.11.7 and later.
	// +kubebuilder:validation:Optional
	KubernetesSvcIPRange *string `json:"kubernetesSvcIpRange,omitempty" tf:"kubernetes_svc_ip_range,omitempty"`

	// Cluster tag, key/value pair format. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Enable multiple AZs for the cluster, only when using HA flavors. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// Cluster name. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Remove addons installed by the default after the cluster creation.
	// +kubebuilder:validation:Optional
	NoAddons *bool `json:"noAddons,omitempty" tf:"no_addons,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Network ID of the subnet used to create the node. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The ID of the VPC used to create the node. Changing this parameter will create a new cluster resource.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

// ClusterV3Spec defines the desired state of ClusterV3
type ClusterV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterV3Parameters `json:"forProvider"`
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
	InitProvider ClusterV3InitParameters `json:"initProvider,omitempty"`
}

// ClusterV3Status defines the observed state of ClusterV3.
type ClusterV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterV3 is the Schema for the ClusterV3s API. Manages a CCE Cluster resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type ClusterV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterType) || (has(self.initProvider) && has(self.initProvider.clusterType))",message="spec.forProvider.clusterType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.containerNetworkType) || (has(self.initProvider) && has(self.initProvider.containerNetworkType))",message="spec.forProvider.containerNetworkType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flavorId) || (has(self.initProvider) && has(self.initProvider.flavorId))",message="spec.forProvider.flavorId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnetId) || (has(self.initProvider) && has(self.initProvider.subnetId))",message="spec.forProvider.subnetId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcId) || (has(self.initProvider) && has(self.initProvider.vpcId))",message="spec.forProvider.vpcId is a required parameter"
	Spec   ClusterV3Spec   `json:"spec"`
	Status ClusterV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterV3List contains a list of ClusterV3s
type ClusterV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterV3 `json:"items"`
}

// Repository type metadata.
var (
	ClusterV3_Kind             = "ClusterV3"
	ClusterV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterV3_Kind}.String()
	ClusterV3_KindAPIVersion   = ClusterV3_Kind + "." + CRDGroupVersion.String()
	ClusterV3_GroupVersionKind = CRDGroupVersion.WithKind(ClusterV3_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterV3{}, &ClusterV3List{})
}
