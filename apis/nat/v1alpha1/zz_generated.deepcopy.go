//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatRuleV2) DeepCopyInto(out *DnatRuleV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatRuleV2.
func (in *DnatRuleV2) DeepCopy() *DnatRuleV2 {
	if in == nil {
		return nil
	}
	out := new(DnatRuleV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DnatRuleV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatRuleV2InitParameters) DeepCopyInto(out *DnatRuleV2InitParameters) {
	*out = *in
	if in.ExternalServicePort != nil {
		in, out := &in.ExternalServicePort, &out.ExternalServicePort
		*out = new(float64)
		**out = **in
	}
	if in.FloatingIPID != nil {
		in, out := &in.FloatingIPID, &out.FloatingIPID
		*out = new(string)
		**out = **in
	}
	if in.InternalServicePort != nil {
		in, out := &in.InternalServicePort, &out.InternalServicePort
		*out = new(float64)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatRuleV2InitParameters.
func (in *DnatRuleV2InitParameters) DeepCopy() *DnatRuleV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(DnatRuleV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatRuleV2List) DeepCopyInto(out *DnatRuleV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DnatRuleV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatRuleV2List.
func (in *DnatRuleV2List) DeepCopy() *DnatRuleV2List {
	if in == nil {
		return nil
	}
	out := new(DnatRuleV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DnatRuleV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatRuleV2Observation) DeepCopyInto(out *DnatRuleV2Observation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ExternalServicePort != nil {
		in, out := &in.ExternalServicePort, &out.ExternalServicePort
		*out = new(float64)
		**out = **in
	}
	if in.FloatingIPAddress != nil {
		in, out := &in.FloatingIPAddress, &out.FloatingIPAddress
		*out = new(string)
		**out = **in
	}
	if in.FloatingIPID != nil {
		in, out := &in.FloatingIPID, &out.FloatingIPID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InternalServicePort != nil {
		in, out := &in.InternalServicePort, &out.InternalServicePort
		*out = new(float64)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatRuleV2Observation.
func (in *DnatRuleV2Observation) DeepCopy() *DnatRuleV2Observation {
	if in == nil {
		return nil
	}
	out := new(DnatRuleV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatRuleV2Parameters) DeepCopyInto(out *DnatRuleV2Parameters) {
	*out = *in
	if in.ExternalServicePort != nil {
		in, out := &in.ExternalServicePort, &out.ExternalServicePort
		*out = new(float64)
		**out = **in
	}
	if in.FloatingIPID != nil {
		in, out := &in.FloatingIPID, &out.FloatingIPID
		*out = new(string)
		**out = **in
	}
	if in.InternalServicePort != nil {
		in, out := &in.InternalServicePort, &out.InternalServicePort
		*out = new(float64)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatRuleV2Parameters.
func (in *DnatRuleV2Parameters) DeepCopy() *DnatRuleV2Parameters {
	if in == nil {
		return nil
	}
	out := new(DnatRuleV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatRuleV2Spec) DeepCopyInto(out *DnatRuleV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatRuleV2Spec.
func (in *DnatRuleV2Spec) DeepCopy() *DnatRuleV2Spec {
	if in == nil {
		return nil
	}
	out := new(DnatRuleV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatRuleV2Status) DeepCopyInto(out *DnatRuleV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatRuleV2Status.
func (in *DnatRuleV2Status) DeepCopy() *DnatRuleV2Status {
	if in == nil {
		return nil
	}
	out := new(DnatRuleV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayV2) DeepCopyInto(out *GatewayV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayV2.
func (in *GatewayV2) DeepCopy() *GatewayV2 {
	if in == nil {
		return nil
	}
	out := new(GatewayV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayV2InitParameters) DeepCopyInto(out *GatewayV2InitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InternalNetworkID != nil {
		in, out := &in.InternalNetworkID, &out.InternalNetworkID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RouterID != nil {
		in, out := &in.RouterID, &out.RouterID
		*out = new(string)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayV2InitParameters.
func (in *GatewayV2InitParameters) DeepCopy() *GatewayV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayV2List) DeepCopyInto(out *GatewayV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GatewayV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayV2List.
func (in *GatewayV2List) DeepCopy() *GatewayV2List {
	if in == nil {
		return nil
	}
	out := new(GatewayV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayV2Observation) DeepCopyInto(out *GatewayV2Observation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InternalNetworkID != nil {
		in, out := &in.InternalNetworkID, &out.InternalNetworkID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RouterID != nil {
		in, out := &in.RouterID, &out.RouterID
		*out = new(string)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayV2Observation.
func (in *GatewayV2Observation) DeepCopy() *GatewayV2Observation {
	if in == nil {
		return nil
	}
	out := new(GatewayV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayV2Parameters) DeepCopyInto(out *GatewayV2Parameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InternalNetworkID != nil {
		in, out := &in.InternalNetworkID, &out.InternalNetworkID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RouterID != nil {
		in, out := &in.RouterID, &out.RouterID
		*out = new(string)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayV2Parameters.
func (in *GatewayV2Parameters) DeepCopy() *GatewayV2Parameters {
	if in == nil {
		return nil
	}
	out := new(GatewayV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayV2Spec) DeepCopyInto(out *GatewayV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayV2Spec.
func (in *GatewayV2Spec) DeepCopy() *GatewayV2Spec {
	if in == nil {
		return nil
	}
	out := new(GatewayV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayV2Status) DeepCopyInto(out *GatewayV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayV2Status.
func (in *GatewayV2Status) DeepCopy() *GatewayV2Status {
	if in == nil {
		return nil
	}
	out := new(GatewayV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatRuleV2) DeepCopyInto(out *SnatRuleV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatRuleV2.
func (in *SnatRuleV2) DeepCopy() *SnatRuleV2 {
	if in == nil {
		return nil
	}
	out := new(SnatRuleV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatRuleV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatRuleV2InitParameters) DeepCopyInto(out *SnatRuleV2InitParameters) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.FloatingIPID != nil {
		in, out := &in.FloatingIPID, &out.FloatingIPID
		*out = new(string)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatRuleV2InitParameters.
func (in *SnatRuleV2InitParameters) DeepCopy() *SnatRuleV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(SnatRuleV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatRuleV2List) DeepCopyInto(out *SnatRuleV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnatRuleV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatRuleV2List.
func (in *SnatRuleV2List) DeepCopy() *SnatRuleV2List {
	if in == nil {
		return nil
	}
	out := new(SnatRuleV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatRuleV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatRuleV2Observation) DeepCopyInto(out *SnatRuleV2Observation) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.FloatingIPID != nil {
		in, out := &in.FloatingIPID, &out.FloatingIPID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatRuleV2Observation.
func (in *SnatRuleV2Observation) DeepCopy() *SnatRuleV2Observation {
	if in == nil {
		return nil
	}
	out := new(SnatRuleV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatRuleV2Parameters) DeepCopyInto(out *SnatRuleV2Parameters) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.FloatingIPID != nil {
		in, out := &in.FloatingIPID, &out.FloatingIPID
		*out = new(string)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatRuleV2Parameters.
func (in *SnatRuleV2Parameters) DeepCopy() *SnatRuleV2Parameters {
	if in == nil {
		return nil
	}
	out := new(SnatRuleV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatRuleV2Spec) DeepCopyInto(out *SnatRuleV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatRuleV2Spec.
func (in *SnatRuleV2Spec) DeepCopy() *SnatRuleV2Spec {
	if in == nil {
		return nil
	}
	out := new(SnatRuleV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatRuleV2Status) DeepCopyInto(out *SnatRuleV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatRuleV2Status.
func (in *SnatRuleV2Status) DeepCopy() *SnatRuleV2Status {
	if in == nil {
		return nil
	}
	out := new(SnatRuleV2Status)
	in.DeepCopyInto(out)
	return out
}
