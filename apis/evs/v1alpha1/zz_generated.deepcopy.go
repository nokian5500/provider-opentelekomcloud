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
func (in *AttachmentInitParameters) DeepCopyInto(out *AttachmentInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentInitParameters.
func (in *AttachmentInitParameters) DeepCopy() *AttachmentInitParameters {
	if in == nil {
		return nil
	}
	out := new(AttachmentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachmentObservation) DeepCopyInto(out *AttachmentObservation) {
	*out = *in
	if in.Device != nil {
		in, out := &in.Device, &out.Device
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentObservation.
func (in *AttachmentObservation) DeepCopy() *AttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(AttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachmentParameters) DeepCopyInto(out *AttachmentParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentParameters.
func (in *AttachmentParameters) DeepCopy() *AttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(AttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeV3) DeepCopyInto(out *VolumeV3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeV3.
func (in *VolumeV3) DeepCopy() *VolumeV3 {
	if in == nil {
		return nil
	}
	out := new(VolumeV3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeV3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeV3InitParameters) DeepCopyInto(out *VolumeV3InitParameters) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.BackupID != nil {
		in, out := &in.BackupID, &out.BackupID
		*out = new(string)
		**out = **in
	}
	if in.Cascade != nil {
		in, out := &in.Cascade, &out.Cascade
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DeviceType != nil {
		in, out := &in.DeviceType, &out.DeviceType
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.KMSID != nil {
		in, out := &in.KMSID, &out.KMSID
		*out = new(string)
		**out = **in
	}
	if in.Multiattach != nil {
		in, out := &in.Multiattach, &out.Multiattach
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
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
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeV3InitParameters.
func (in *VolumeV3InitParameters) DeepCopy() *VolumeV3InitParameters {
	if in == nil {
		return nil
	}
	out := new(VolumeV3InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeV3List) DeepCopyInto(out *VolumeV3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeV3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeV3List.
func (in *VolumeV3List) DeepCopy() *VolumeV3List {
	if in == nil {
		return nil
	}
	out := new(VolumeV3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeV3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeV3Observation) DeepCopyInto(out *VolumeV3Observation) {
	*out = *in
	if in.Attachment != nil {
		in, out := &in.Attachment, &out.Attachment
		*out = make([]AttachmentObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.BackupID != nil {
		in, out := &in.BackupID, &out.BackupID
		*out = new(string)
		**out = **in
	}
	if in.Cascade != nil {
		in, out := &in.Cascade, &out.Cascade
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DeviceType != nil {
		in, out := &in.DeviceType, &out.DeviceType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.KMSID != nil {
		in, out := &in.KMSID, &out.KMSID
		*out = new(string)
		**out = **in
	}
	if in.Multiattach != nil {
		in, out := &in.Multiattach, &out.Multiattach
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
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
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	if in.Wwn != nil {
		in, out := &in.Wwn, &out.Wwn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeV3Observation.
func (in *VolumeV3Observation) DeepCopy() *VolumeV3Observation {
	if in == nil {
		return nil
	}
	out := new(VolumeV3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeV3Parameters) DeepCopyInto(out *VolumeV3Parameters) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.BackupID != nil {
		in, out := &in.BackupID, &out.BackupID
		*out = new(string)
		**out = **in
	}
	if in.Cascade != nil {
		in, out := &in.Cascade, &out.Cascade
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DeviceType != nil {
		in, out := &in.DeviceType, &out.DeviceType
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.KMSID != nil {
		in, out := &in.KMSID, &out.KMSID
		*out = new(string)
		**out = **in
	}
	if in.Multiattach != nil {
		in, out := &in.Multiattach, &out.Multiattach
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
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
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeV3Parameters.
func (in *VolumeV3Parameters) DeepCopy() *VolumeV3Parameters {
	if in == nil {
		return nil
	}
	out := new(VolumeV3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeV3Spec) DeepCopyInto(out *VolumeV3Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeV3Spec.
func (in *VolumeV3Spec) DeepCopy() *VolumeV3Spec {
	if in == nil {
		return nil
	}
	out := new(VolumeV3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeV3Status) DeepCopyInto(out *VolumeV3Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeV3Status.
func (in *VolumeV3Status) DeepCopy() *VolumeV3Status {
	if in == nil {
		return nil
	}
	out := new(VolumeV3Status)
	in.DeepCopyInto(out)
	return out
}
