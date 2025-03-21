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

package v1alpha1

import (
	ubus "github.com/daimonaslabs/go-ubus-rpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make generate" to regenerate code after modifying this file

// FirewallConfigSpec defines the desired state of FirewallConfig
type FirewallConfigSpec struct {
	Defaults DefaultsSection
}

// FirewallConfigStatus defines the observed state of FirewallConfig
type FirewallConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FirewallConfig is the Schema for the firewallconfigs API
type FirewallConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirewallConfigSpec   `json:"spec,omitempty"`
	Status FirewallConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallConfigList contains a list of FirewallConfig
type FirewallConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallConfig `json:"items"`
}

type DefaultsSection struct {
	// Accepts redirects. Implemented upstream in Linux Kernel.
	AcceptRedirects ubus.UbusBool `json:"acceptRedirects,omitempty" ubus:"accept_redirects,omitempty"`
	// Implemented upstream in Linux Kernel.
	AcceptSourceRoute ubus.UbusBool `json:"acceptSourceRoute,omitempty" ubus:"accept_source_route,omitempty"`
	// Determines method of packet rejection.
	AnyRejectCode int `json:"anyRejectCode,omitempty" ubus:"any_reject_code,omitempty"`
	// Enable Conntrack helpers.
	AutoHelper ubus.UbusBool `json:"autoHelper,omitempty" ubus:"auth_helper,omitempty"`
	// (fw4 only, OpenWRT 22.03 and later) Enable automatic nftables includes under /usr/share/nftables.d/
	AutoIncludes ubus.UbusBool `json:"autoIncludes,omitempty" ubus:"auto_includes,omitempty"`
	// Enable generation of custom rule chain hooks for user generated rules. User rules would be typically
	// stored in firewall.user but some packages e.g. BCP38 also make use of these hooks.
	CustomChains ubus.UbusBool `json:"customChains,omitempty" ubus:"custom_chains,omitempty"`
	// Disable IPv6 firewall rules. (not supported by fw4).
	DisableIPv6 ubus.UbusBool `json:"disableIPv6,omitempty" ubus:"disable_ipv6,omitempty"`
	// Drop invalid packets (e.g. not matching any active connection).
	DropInvalid ubus.UbusBool `json:"dropInvalid,omitempty" ubus:"drop_invalid"`
	// Enable software flow offloading for connections (decrease cpu load / increase routing throughput).
	FlowOffloading ubus.UbusBool `json:"flowOffloading,omitempty" ubus:"flow_offloading,omitempty"`
	// Enable hardware flow offloading for connecions (depends on flow_offloading and hw capability).
	FlowOffloadingHW ubus.UbusBool `json:"flowOffloadingHW,omitempty" ubus:"flow_offloading_hw,omitempty"`
	// Set policy for the FORWARD chain of the filter table.
	Forward string `json:"forward,omitempty" ubus:"forward,omitempty"`
	// Set policy for the INPUT chain of the filter table.
	Input string `json:"input,omitempty" ubus:"input,omitempty"`
	// Set policy for the OUTPUT chain of the filter table.
	Output string `json:"output,omitempty" ubus:"output,omitempty"`
	// Enable SYN flood protection (obsoleted by synflood_protect setting).
	SynFlood ubus.UbusBool `json:"synFlood,omitempty" ubsu:"syn_flood,omitempty"`
	// Enable SYN flood protection.
	SynFloodProtect ubus.UbusBool `json:"synFloodProtect,omitempty" ubus:"synflood_protect,omitempty"`
	// Set rate limit (packets/second) for SYN packets above which the traffic is considered a flood.
	SynFloodRate string `json:"synFloodRate,omitempty" ubus:"synflood_rate,omitempty"`
	// Set burst limit for SYN packets above which the traffic is considered a flood if it exceeds the allowed rate.
	SynFloodBurst string `json:"synFloodBurst,omitempty" ubus:"synflood_burst,omitempty"`
	// 0 Disable, 1 Enable, 2 Enable when requested for ingress (but disable for egress) Explicit Congestion
	// Notification. Affects only traffic originating from the router itself. Implemented upstream in Linux Kernel.
	TCPECN int `json:"tcpECN,omitempty" ubus:"tcp_ecn,omitempty"`
	// Enable the use of SYN cookies.
	TCPSynCookies ubus.UbusBool `json:"tcpSynCookies,omitempty" ubus:"tcp_syncookies,omitempty"`
	// Determines method of packet rejection.
	TCPRejectCode int `json:"tcpRejectCode,omitempty" ubus:"tcp_reject_code,omitempty"`
	// Enable TCP window scaling.
	TCPWindowScaling ubus.UbusBool `json:"tcpWindowScaling,omitempty" ubus:"tcp_window_scaling,omitempty"`
}

type ZoneSection struct {
	Name string `json:"name,omitempty" ubus:"name,omitempty"`
}

func init() {
	SchemeBuilder.Register(&FirewallConfig{}, &FirewallConfigList{})
}
