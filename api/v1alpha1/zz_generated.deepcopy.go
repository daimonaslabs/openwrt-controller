//go:build !ignore_autogenerated

/*
Copyright 2025 Daimonas Labs.

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
	go_ubus_rpc "github.com/daimonaslabs/go-ubus-rpc"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootSection) DeepCopyInto(out *BootSection) {
	*out = *in
	out.UCIConfigOptionsStatic = in.UCIConfigOptionsStatic
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootSection.
func (in *BootSection) DeepCopy() *BootSection {
	if in == nil {
		return nil
	}
	out := new(BootSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSection) DeepCopyInto(out *ClientSection) {
	*out = *in
	out.UCIConfigOptionsStatic = in.UCIConfigOptionsStatic
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSection.
func (in *ClientSection) DeepCopy() *ClientSection {
	if in == nil {
		return nil
	}
	out := new(ClientSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DHCPConfig) DeepCopyInto(out *DHCPConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DHCPConfig.
func (in *DHCPConfig) DeepCopy() *DHCPConfig {
	if in == nil {
		return nil
	}
	out := new(DHCPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DHCPConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DHCPConfigList) DeepCopyInto(out *DHCPConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DHCPConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DHCPConfigList.
func (in *DHCPConfigList) DeepCopy() *DHCPConfigList {
	if in == nil {
		return nil
	}
	out := new(DHCPConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DHCPConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DHCPConfigSpec) DeepCopyInto(out *DHCPConfigSpec) {
	*out = *in
	if in.Dnsmasq != nil {
		in, out := &in.Dnsmasq, &out.Dnsmasq
		*out = make([]DnsmasqSection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DHCPConfigSpec.
func (in *DHCPConfigSpec) DeepCopy() *DHCPConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DHCPConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DHCPConfigStatus) DeepCopyInto(out *DHCPConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DHCPConfigStatus.
func (in *DHCPConfigStatus) DeepCopy() *DHCPConfigStatus {
	if in == nil {
		return nil
	}
	out := new(DHCPConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DHCPSection) DeepCopyInto(out *DHCPSection) {
	*out = *in
	out.UCIConfigOptionsStatic = in.UCIConfigOptionsStatic
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DHCPSection.
func (in *DHCPSection) DeepCopy() *DHCPSection {
	if in == nil {
		return nil
	}
	out := new(DHCPSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultsSection) DeepCopyInto(out *DefaultsSection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultsSection.
func (in *DefaultsSection) DeepCopy() *DefaultsSection {
	if in == nil {
		return nil
	}
	out := new(DefaultsSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnsmasqSection) DeepCopyInto(out *DnsmasqSection) {
	*out = *in
	out.UCIConfigOptionsStatic = in.UCIConfigOptionsStatic
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AddnHosts != nil {
		in, out := &in.AddnHosts, &out.AddnHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AddnMount != nil {
		in, out := &in.AddnMount, &out.AddnMount
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BogusNXDOMAIN != nil {
		in, out := &in.BogusNXDOMAIN, &out.BogusNXDOMAIN
		*out = make([]go_ubus_rpc.IP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPSet != nil {
		in, out := &in.IPSet, &out.IPSet
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ListenAddress != nil {
		in, out := &in.ListenAddress, &out.ListenAddress
		*out = make([]go_ubus_rpc.IP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NotInterface != nil {
		in, out := &in.NotInterface, &out.NotInterface
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RebindDomain != nil {
		in, out := &in.RebindDomain, &out.RebindDomain
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RevServer != nil {
		in, out := &in.RevServer, &out.RevServer
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnsmasqSection.
func (in *DnsmasqSection) DeepCopy() *DnsmasqSection {
	if in == nil {
		return nil
	}
	out := new(DnsmasqSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallConfig) DeepCopyInto(out *FirewallConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallConfig.
func (in *FirewallConfig) DeepCopy() *FirewallConfig {
	if in == nil {
		return nil
	}
	out := new(FirewallConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirewallConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallConfigList) DeepCopyInto(out *FirewallConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FirewallConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallConfigList.
func (in *FirewallConfigList) DeepCopy() *FirewallConfigList {
	if in == nil {
		return nil
	}
	out := new(FirewallConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirewallConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallConfigSpec) DeepCopyInto(out *FirewallConfigSpec) {
	*out = *in
	out.Defaults = in.Defaults
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallConfigSpec.
func (in *FirewallConfigSpec) DeepCopy() *FirewallConfigSpec {
	if in == nil {
		return nil
	}
	out := new(FirewallConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallConfigStatus) DeepCopyInto(out *FirewallConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallConfigStatus.
func (in *FirewallConfigStatus) DeepCopy() *FirewallConfigStatus {
	if in == nil {
		return nil
	}
	out := new(FirewallConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingSection) DeepCopyInto(out *ForwardingSection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingSection.
func (in *ForwardingSection) DeepCopy() *ForwardingSection {
	if in == nil {
		return nil
	}
	out := new(ForwardingSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostSection) DeepCopyInto(out *HostSection) {
	*out = *in
	out.UCIConfigOptionsStatic = in.UCIConfigOptionsStatic
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostSection.
func (in *HostSection) DeepCopy() *HostSection {
	if in == nil {
		return nil
	}
	out := new(HostSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPSetSection) DeepCopyInto(out *IPSetSection) {
	*out = *in
	out.UCIConfigOptionsStatic = in.UCIConfigOptionsStatic
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPSetSection.
func (in *IPSetSection) DeepCopy() *IPSetSection {
	if in == nil {
		return nil
	}
	out := new(IPSetSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedirectSection) DeepCopyInto(out *RedirectSection) {
	*out = *in
	in.Monthdays.DeepCopyInto(&out.Monthdays)
	if in.ReflectionZone != nil {
		in, out := &in.ReflectionZone, &out.ReflectionZone
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.SrcDIP.DeepCopyInto(&out.SrcDIP)
	in.SrcIP.DeepCopyInto(&out.SrcIP)
	in.SrcMAC.DeepCopyInto(&out.SrcMAC)
	in.StartDate.DeepCopyInto(&out.StartDate)
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.StopDate.DeepCopyInto(&out.StopDate)
	in.StopTime.DeepCopyInto(&out.StopTime)
	in.Weekdays.DeepCopyInto(&out.Weekdays)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedirectSection.
func (in *RedirectSection) DeepCopy() *RedirectSection {
	if in == nil {
		return nil
	}
	out := new(RedirectSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelaySection) DeepCopyInto(out *RelaySection) {
	*out = *in
	out.UCIConfigOptionsStatic = in.UCIConfigOptionsStatic
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelaySection.
func (in *RelaySection) DeepCopy() *RelaySection {
	if in == nil {
		return nil
	}
	out := new(RelaySection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSection) DeepCopyInto(out *RuleSection) {
	*out = *in
	if in.ICMPType != nil {
		in, out := &in.ICMPType, &out.ICMPType
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Monthdays.DeepCopyInto(&out.Monthdays)
	in.SrcIP.DeepCopyInto(&out.SrcIP)
	in.SrcMAC.DeepCopyInto(&out.SrcMAC)
	in.StartDate.DeepCopyInto(&out.StartDate)
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.StopDate.DeepCopyInto(&out.StopDate)
	in.StopTime.DeepCopyInto(&out.StopTime)
	in.Weekdays.DeepCopyInto(&out.Weekdays)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSection.
func (in *RuleSection) DeepCopy() *RuleSection {
	if in == nil {
		return nil
	}
	out := new(RuleSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WifiIfaceSection) DeepCopyInto(out *WifiIfaceSection) {
	*out = *in
	out.UCIConfigOptionsStatic = in.UCIConfigOptionsStatic
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WifiIfaceSection.
func (in *WifiIfaceSection) DeepCopy() *WifiIfaceSection {
	if in == nil {
		return nil
	}
	out := new(WifiIfaceSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WirelessConfig) DeepCopyInto(out *WirelessConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WirelessConfig.
func (in *WirelessConfig) DeepCopy() *WirelessConfig {
	if in == nil {
		return nil
	}
	out := new(WirelessConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WirelessConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WirelessConfigList) DeepCopyInto(out *WirelessConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WirelessConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WirelessConfigList.
func (in *WirelessConfigList) DeepCopy() *WirelessConfigList {
	if in == nil {
		return nil
	}
	out := new(WirelessConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WirelessConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WirelessConfigSpec) DeepCopyInto(out *WirelessConfigSpec) {
	*out = *in
	if in.Dnsmasq != nil {
		in, out := &in.Dnsmasq, &out.Dnsmasq
		*out = make([]DnsmasqSection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WirelessConfigSpec.
func (in *WirelessConfigSpec) DeepCopy() *WirelessConfigSpec {
	if in == nil {
		return nil
	}
	out := new(WirelessConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WirelessConfigStatus) DeepCopyInto(out *WirelessConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WirelessConfigStatus.
func (in *WirelessConfigStatus) DeepCopy() *WirelessConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WirelessConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSection) DeepCopyInto(out *ZoneSection) {
	*out = *in
	if in.Device != nil {
		in, out := &in.Device, &out.Device
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Helper != nil {
		in, out := &in.Helper, &out.Helper
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MasqDest != nil {
		in, out := &in.MasqDest, &out.MasqDest
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MasqSrc != nil {
		in, out := &in.MasqSrc, &out.MasqSrc
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSection.
func (in *ZoneSection) DeepCopy() *ZoneSection {
	if in == nil {
		return nil
	}
	out := new(ZoneSection)
	in.DeepCopyInto(out)
	return out
}
