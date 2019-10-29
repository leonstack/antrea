// +build !ignore_autogenerated

/*
Copyright 2019 Antrea Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressGroup) DeepCopyInto(out *AddressGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]IPAddress, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(IPAddress, len(*in))
				copy(*out, *in)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressGroup.
func (in *AddressGroup) DeepCopy() *AddressGroup {
	if in == nil {
		return nil
	}
	out := new(AddressGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressGroupList) DeepCopyInto(out *AddressGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddressGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressGroupList.
func (in *AddressGroupList) DeepCopy() *AddressGroupList {
	if in == nil {
		return nil
	}
	out := new(AddressGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressGroupPatch) DeepCopyInto(out *AddressGroupPatch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AddedIPAddresses != nil {
		in, out := &in.AddedIPAddresses, &out.AddedIPAddresses
		*out = make([]IPAddress, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(IPAddress, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.RemovedIPAddresses != nil {
		in, out := &in.RemovedIPAddresses, &out.RemovedIPAddresses
		*out = make([]IPAddress, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(IPAddress, len(*in))
				copy(*out, *in)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressGroupPatch.
func (in *AddressGroupPatch) DeepCopy() *AddressGroupPatch {
	if in == nil {
		return nil
	}
	out := new(AddressGroupPatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressGroupPatch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedToGroup) DeepCopyInto(out *AppliedToGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]PodReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedToGroup.
func (in *AppliedToGroup) DeepCopy() *AppliedToGroup {
	if in == nil {
		return nil
	}
	out := new(AppliedToGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedToGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedToGroupList) DeepCopyInto(out *AppliedToGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppliedToGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedToGroupList.
func (in *AppliedToGroupList) DeepCopy() *AppliedToGroupList {
	if in == nil {
		return nil
	}
	out := new(AppliedToGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedToGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedToGroupPatch) DeepCopyInto(out *AppliedToGroupPatch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AddedPods != nil {
		in, out := &in.AddedPods, &out.AddedPods
		*out = make([]PodReference, len(*in))
		copy(*out, *in)
	}
	if in.RemovedPods != nil {
		in, out := &in.RemovedPods, &out.RemovedPods
		*out = make([]PodReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedToGroupPatch.
func (in *AppliedToGroupPatch) DeepCopy() *AppliedToGroupPatch {
	if in == nil {
		return nil
	}
	out := new(AppliedToGroupPatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedToGroupPatch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in IPAddress) DeepCopyInto(out *IPAddress) {
	{
		in := &in
		*out = make(IPAddress, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddress.
func (in IPAddress) DeepCopy() IPAddress {
	if in == nil {
		return nil
	}
	out := new(IPAddress)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlock) DeepCopyInto(out *IPBlock) {
	*out = *in
	in.CIDR.DeepCopyInto(&out.CIDR)
	if in.Except != nil {
		in, out := &in.Except, &out.Except
		*out = make([]IPNet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlock.
func (in *IPBlock) DeepCopy() *IPBlock {
	if in == nil {
		return nil
	}
	out := new(IPBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPNet) DeepCopyInto(out *IPNet) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = make(IPAddress, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPNet.
func (in *IPNet) DeepCopy() *IPNet {
	if in == nil {
		return nil
	}
	out := new(IPNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicy) DeepCopyInto(out *NetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]NetworkPolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AppliedToGroups != nil {
		in, out := &in.AppliedToGroups, &out.AppliedToGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicy.
func (in *NetworkPolicy) DeepCopy() *NetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyList) DeepCopyInto(out *NetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyList.
func (in *NetworkPolicyList) DeepCopy() *NetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyPeer) DeepCopyInto(out *NetworkPolicyPeer) {
	*out = *in
	if in.AddressGroups != nil {
		in, out := &in.AddressGroups, &out.AddressGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPBlocks != nil {
		in, out := &in.IPBlocks, &out.IPBlocks
		*out = make([]IPBlock, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyPeer.
func (in *NetworkPolicyPeer) DeepCopy() *NetworkPolicyPeer {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyRule) DeepCopyInto(out *NetworkPolicyRule) {
	*out = *in
	in.From.DeepCopyInto(&out.From)
	in.To.DeepCopyInto(&out.To)
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyRule.
func (in *NetworkPolicyRule) DeepCopy() *NetworkPolicyRule {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodReference) DeepCopyInto(out *PodReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodReference.
func (in *PodReference) DeepCopy() *PodReference {
	if in == nil {
		return nil
	}
	out := new(PodReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(Protocol)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}