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
	ubus "github.com/daimonaslabs/go-ubus-rpc/pkg/encoding"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BootSections []ubus.BootSection
type ClientSections []ubus.ClientSection
type DHCPSections []ubus.DHCPSection
type DnsmasqSections []ubus.DnsmasqSection
type HostSections []ubus.HostSection
type IPSetSections []ubus.IPSetSection
type RelaySections []ubus.RelaySection

// Important: Run "make generate" to regenerate code after modifying this file

// DHCPConfigSpec defines the desired state of DHCPConfig
type DHCPConfigSpec struct {
	DnsmasqSections
}

// DHCPConfigStatus defines the observed state of DHCPConfig
type DHCPConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DHCPConfig is the Schema for the dhcpconfigs API
type DHCPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DHCPConfigSpec   `json:"spec,omitempty"`
	Status DHCPConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DHCPConfigList contains a list of DHCPConfig
type DHCPConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DHCPConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DHCPConfig{}, &DHCPConfigList{})
}
