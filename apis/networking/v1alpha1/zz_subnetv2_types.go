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

type AllocationPoolsInitParameters struct {

	// The ending address.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The starting address.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AllocationPoolsObservation struct {

	// The ending address.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The starting address.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AllocationPoolsParameters struct {

	// The ending address.
	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// The starting address.
	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type HostRoutesInitParameters struct {

	// The destination CIDR.
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	// The next hop in the route.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`
}

type HostRoutesObservation struct {

	// The destination CIDR.
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	// The next hop in the route.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`
}

type HostRoutesParameters struct {

	// The destination CIDR.
	// +kubebuilder:validation:Optional
	DestinationCidr *string `json:"destinationCidr" tf:"destination_cidr,omitempty"`

	// The next hop in the route.
	// +kubebuilder:validation:Optional
	NextHop *string `json:"nextHop" tf:"next_hop,omitempty"`
}

type SubnetV2InitParameters struct {

	// An array of sub-ranges of CIDR available for
	// dynamic allocation to ports. The allocation_pool object structure is
	// documented below. Changing this creates a new subnet.
	AllocationPools []AllocationPoolsInitParameters `json:"allocationPools,omitempty" tf:"allocation_pools,omitempty"`

	// CIDR representing IP range for this subnet, based on IP
	// version. Changing this creates a new subnet.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet. Default value is ["100.125.4.25", "100.125.129.199"]
	// +listType=set
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting no_gateway will cause a default
	// gateway of .1 to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	HostRoutes []HostRoutesInitParameters `json:"hostRoutes,omitempty" tf:"host_routes,omitempty"`

	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of the parent network. Changing this
	// creates a new subnet.
	// +crossplane:generate:reference:type=NetworkV2
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a NetworkV2 to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type SubnetV2Observation struct {

	// An array of sub-ranges of CIDR available for
	// dynamic allocation to ports. The allocation_pool object structure is
	// documented below. Changing this creates a new subnet.
	AllocationPools []AllocationPoolsObservation `json:"allocationPools,omitempty" tf:"allocation_pools,omitempty"`

	// CIDR representing IP range for this subnet, based on IP
	// version. Changing this creates a new subnet.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet. Default value is ["100.125.4.25", "100.125.129.199"]
	// +listType=set
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting no_gateway will cause a default
	// gateway of .1 to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	HostRoutes []HostRoutesObservation `json:"hostRoutes,omitempty" tf:"host_routes,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type SubnetV2Parameters struct {

	// An array of sub-ranges of CIDR available for
	// dynamic allocation to ports. The allocation_pool object structure is
	// documented below. Changing this creates a new subnet.
	// +kubebuilder:validation:Optional
	AllocationPools []AllocationPoolsParameters `json:"allocationPools,omitempty" tf:"allocation_pools,omitempty"`

	// CIDR representing IP range for this subnet, based on IP
	// version. Changing this creates a new subnet.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet. Default value is ["100.125.4.25", "100.125.129.199"]
	// +kubebuilder:validation:Optional
	// +listType=set
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting no_gateway will cause a default
	// gateway of .1 to be used. Changing this updates the gateway IP of the
	// existing subnet.
	// +kubebuilder:validation:Optional
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	// +kubebuilder:validation:Optional
	HostRoutes []HostRoutesParameters `json:"hostRoutes,omitempty" tf:"host_routes,omitempty"`

	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	// +kubebuilder:validation:Optional
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of the parent network. Changing this
	// creates a new subnet.
	// +crossplane:generate:reference:type=NetworkV2
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a NetworkV2 to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	// +kubebuilder:validation:Optional
	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// SubnetV2Spec defines the desired state of SubnetV2
type SubnetV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetV2Parameters `json:"forProvider"`
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
	InitProvider SubnetV2InitParameters `json:"initProvider,omitempty"`
}

// SubnetV2Status defines the observed state of SubnetV2.
type SubnetV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubnetV2 is the Schema for the SubnetV2s API. Manages a VPC Subnet resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type SubnetV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidr) || (has(self.initProvider) && has(self.initProvider.cidr))",message="spec.forProvider.cidr is a required parameter"
	Spec   SubnetV2Spec   `json:"spec"`
	Status SubnetV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetV2List contains a list of SubnetV2s
type SubnetV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetV2 `json:"items"`
}

// Repository type metadata.
var (
	SubnetV2_Kind             = "SubnetV2"
	SubnetV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetV2_Kind}.String()
	SubnetV2_KindAPIVersion   = SubnetV2_Kind + "." + CRDGroupVersion.String()
	SubnetV2_GroupVersionKind = CRDGroupVersion.WithKind(SubnetV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetV2{}, &SubnetV2List{})
}
