//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootDiskObservation) DeepCopyInto(out *BootDiskObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootDiskObservation.
func (in *BootDiskObservation) DeepCopy() *BootDiskObservation {
	if in == nil {
		return nil
	}
	out := new(BootDiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootDiskParameters) DeepCopyInto(out *BootDiskParameters) {
	*out = *in
	if in.AutoDelete != nil {
		in, out := &in.AutoDelete, &out.AutoDelete
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(string)
		**out = **in
	}
	if in.InitializeParams != nil {
		in, out := &in.InitializeParams, &out.InitializeParams
		*out = make([]InitializeParamsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootDiskParameters.
func (in *BootDiskParameters) DeepCopy() *BootDiskParameters {
	if in == nil {
		return nil
	}
	out := new(BootDiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordObservation) DeepCopyInto(out *DNSRecordObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordObservation.
func (in *DNSRecordObservation) DeepCopy() *DNSRecordObservation {
	if in == nil {
		return nil
	}
	out := new(DNSRecordObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordParameters) DeepCopyInto(out *DNSRecordParameters) {
	*out = *in
	if in.DNSZoneID != nil {
		in, out := &in.DNSZoneID, &out.DNSZoneID
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.Ptr != nil {
		in, out := &in.Ptr, &out.Ptr
		*out = new(bool)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordParameters.
func (in *DNSRecordParameters) DeepCopy() *DNSRecordParameters {
	if in == nil {
		return nil
	}
	out := new(DNSRecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostAffinityRulesObservation) DeepCopyInto(out *HostAffinityRulesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostAffinityRulesObservation.
func (in *HostAffinityRulesObservation) DeepCopy() *HostAffinityRulesObservation {
	if in == nil {
		return nil
	}
	out := new(HostAffinityRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostAffinityRulesParameters) DeepCopyInto(out *HostAffinityRulesParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Op != nil {
		in, out := &in.Op, &out.Op
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostAffinityRulesParameters.
func (in *HostAffinityRulesParameters) DeepCopy() *HostAffinityRulesParameters {
	if in == nil {
		return nil
	}
	out := new(HostAffinityRulesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6DNSRecordObservation) DeepCopyInto(out *IPv6DNSRecordObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6DNSRecordObservation.
func (in *IPv6DNSRecordObservation) DeepCopy() *IPv6DNSRecordObservation {
	if in == nil {
		return nil
	}
	out := new(IPv6DNSRecordObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6DNSRecordParameters) DeepCopyInto(out *IPv6DNSRecordParameters) {
	*out = *in
	if in.DNSZoneID != nil {
		in, out := &in.DNSZoneID, &out.DNSZoneID
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.Ptr != nil {
		in, out := &in.Ptr, &out.Ptr
		*out = new(bool)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6DNSRecordParameters.
func (in *IPv6DNSRecordParameters) DeepCopy() *IPv6DNSRecordParameters {
	if in == nil {
		return nil
	}
	out := new(IPv6DNSRecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializeParamsObservation) DeepCopyInto(out *InitializeParamsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializeParamsObservation.
func (in *InitializeParamsObservation) DeepCopy() *InitializeParamsObservation {
	if in == nil {
		return nil
	}
	out := new(InitializeParamsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializeParamsParameters) DeepCopyInto(out *InitializeParamsParameters) {
	*out = *in
	if in.BlockSize != nil {
		in, out := &in.BlockSize, &out.BlockSize
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializeParamsParameters.
func (in *InitializeParamsParameters) DeepCopy() *InitializeParamsParameters {
	if in == nil {
		return nil
	}
	out := new(InitializeParamsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceObservation) DeepCopyInto(out *InstanceObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LocalDisk != nil {
		in, out := &in.LocalDisk, &out.LocalDisk
		*out = make([]LocalDiskObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]NetworkInterfaceObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceObservation.
func (in *InstanceObservation) DeepCopy() *InstanceObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceParameters) DeepCopyInto(out *InstanceParameters) {
	*out = *in
	if in.AllowRecreate != nil {
		in, out := &in.AllowRecreate, &out.AllowRecreate
		*out = new(bool)
		**out = **in
	}
	if in.AllowStoppingForUpdate != nil {
		in, out := &in.AllowStoppingForUpdate, &out.AllowStoppingForUpdate
		*out = new(bool)
		**out = **in
	}
	if in.BootDisk != nil {
		in, out := &in.BootDisk, &out.BootDisk
		*out = make([]BootDiskParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.FolderIDRef != nil {
		in, out := &in.FolderIDRef, &out.FolderIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FolderIDSelector != nil {
		in, out := &in.FolderIDSelector, &out.FolderIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LocalDisk != nil {
		in, out := &in.LocalDisk, &out.LocalDisk
		*out = make([]LocalDiskParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkAccelerationType != nil {
		in, out := &in.NetworkAccelerationType, &out.NetworkAccelerationType
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]NetworkInterfaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlacementPolicy != nil {
		in, out := &in.PlacementPolicy, &out.PlacementPolicy
		*out = make([]PlacementPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourcesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SchedulingPolicy != nil {
		in, out := &in.SchedulingPolicy, &out.SchedulingPolicy
		*out = make([]SchedulingPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecondaryDisk != nil {
		in, out := &in.SecondaryDisk, &out.SecondaryDisk
		*out = make([]SecondaryDiskParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountIDRef != nil {
		in, out := &in.ServiceAccountIDRef, &out.ServiceAccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountIDSelector != nil {
		in, out := &in.ServiceAccountIDSelector, &out.ServiceAccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceParameters.
func (in *InstanceParameters) DeepCopy() *InstanceParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskObservation) DeepCopyInto(out *LocalDiskObservation) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskObservation.
func (in *LocalDiskObservation) DeepCopy() *LocalDiskObservation {
	if in == nil {
		return nil
	}
	out := new(LocalDiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskParameters) DeepCopyInto(out *LocalDiskParameters) {
	*out = *in
	if in.SizeBytes != nil {
		in, out := &in.SizeBytes, &out.SizeBytes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskParameters.
func (in *LocalDiskParameters) DeepCopy() *LocalDiskParameters {
	if in == nil {
		return nil
	}
	out := new(LocalDiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATDNSRecordObservation) DeepCopyInto(out *NATDNSRecordObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATDNSRecordObservation.
func (in *NATDNSRecordObservation) DeepCopy() *NATDNSRecordObservation {
	if in == nil {
		return nil
	}
	out := new(NATDNSRecordObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATDNSRecordParameters) DeepCopyInto(out *NATDNSRecordParameters) {
	*out = *in
	if in.DNSZoneID != nil {
		in, out := &in.DNSZoneID, &out.DNSZoneID
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.Ptr != nil {
		in, out := &in.Ptr, &out.Ptr
		*out = new(bool)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATDNSRecordParameters.
func (in *NATDNSRecordParameters) DeepCopy() *NATDNSRecordParameters {
	if in == nil {
		return nil
	}
	out := new(NATDNSRecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceObservation) DeepCopyInto(out *NetworkInterfaceObservation) {
	*out = *in
	if in.Index != nil {
		in, out := &in.Index, &out.Index
		*out = new(float64)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.NATIPVersion != nil {
		in, out := &in.NATIPVersion, &out.NATIPVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceObservation.
func (in *NetworkInterfaceObservation) DeepCopy() *NetworkInterfaceObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceParameters) DeepCopyInto(out *NetworkInterfaceParameters) {
	*out = *in
	if in.DNSRecord != nil {
		in, out := &in.DNSRecord, &out.DNSRecord
		*out = make([]DNSRecordParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.IPv6DNSRecord != nil {
		in, out := &in.IPv6DNSRecord, &out.IPv6DNSRecord
		*out = make([]IPv6DNSRecordParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NAT != nil {
		in, out := &in.NAT, &out.NAT
		*out = new(bool)
		**out = **in
	}
	if in.NATDNSRecord != nil {
		in, out := &in.NATDNSRecord, &out.NATDNSRecord
		*out = make([]NATDNSRecordParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NATIPAddress != nil {
		in, out := &in.NATIPAddress, &out.NATIPAddress
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIdsRefs != nil {
		in, out := &in.SecurityGroupIdsRefs, &out.SecurityGroupIdsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroupIdsSelector != nil {
		in, out := &in.SecurityGroupIdsSelector, &out.SecurityGroupIdsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceParameters.
func (in *NetworkInterfaceParameters) DeepCopy() *NetworkInterfaceParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementPolicyObservation) DeepCopyInto(out *PlacementPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementPolicyObservation.
func (in *PlacementPolicyObservation) DeepCopy() *PlacementPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(PlacementPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementPolicyParameters) DeepCopyInto(out *PlacementPolicyParameters) {
	*out = *in
	if in.HostAffinityRules != nil {
		in, out := &in.HostAffinityRules, &out.HostAffinityRules
		*out = make([]HostAffinityRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlacementGroupID != nil {
		in, out := &in.PlacementGroupID, &out.PlacementGroupID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementPolicyParameters.
func (in *PlacementPolicyParameters) DeepCopy() *PlacementPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(PlacementPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesObservation) DeepCopyInto(out *ResourcesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesObservation.
func (in *ResourcesObservation) DeepCopy() *ResourcesObservation {
	if in == nil {
		return nil
	}
	out := new(ResourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesParameters) DeepCopyInto(out *ResourcesParameters) {
	*out = *in
	if in.CoreFraction != nil {
		in, out := &in.CoreFraction, &out.CoreFraction
		*out = new(float64)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.Gpus != nil {
		in, out := &in.Gpus, &out.Gpus
		*out = new(float64)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesParameters.
func (in *ResourcesParameters) DeepCopy() *ResourcesParameters {
	if in == nil {
		return nil
	}
	out := new(ResourcesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyObservation) DeepCopyInto(out *SchedulingPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyObservation.
func (in *SchedulingPolicyObservation) DeepCopy() *SchedulingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyParameters) DeepCopyInto(out *SchedulingPolicyParameters) {
	*out = *in
	if in.Preemptible != nil {
		in, out := &in.Preemptible, &out.Preemptible
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyParameters.
func (in *SchedulingPolicyParameters) DeepCopy() *SchedulingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryDiskObservation) DeepCopyInto(out *SecondaryDiskObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryDiskObservation.
func (in *SecondaryDiskObservation) DeepCopy() *SecondaryDiskObservation {
	if in == nil {
		return nil
	}
	out := new(SecondaryDiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryDiskParameters) DeepCopyInto(out *SecondaryDiskParameters) {
	*out = *in
	if in.AutoDelete != nil {
		in, out := &in.AutoDelete, &out.AutoDelete
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryDiskParameters.
func (in *SecondaryDiskParameters) DeepCopy() *SecondaryDiskParameters {
	if in == nil {
		return nil
	}
	out := new(SecondaryDiskParameters)
	in.DeepCopyInto(out)
	return out
}
