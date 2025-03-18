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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DHCPConfigSpec defines the desired state of DHCPConfig
type DHCPConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DHCPConfig. Edit dhcpconfig_types.go to remove/update
	// generic XConfigSection types probably won't work, need separate structs for each config type
	// and each option to be its own field in the struct
	Dnsmasq []DnsmasqSection
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

// DHCPConfigList contains a list of DHCPConfig
type DHCPConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"me
// +kubebuilder:object:root=true
tadata,omitempty"`
	Items []DHCPConfig `json:"items"`
}

// TODO add all options for dnsmasq sections here
type DnsmasqSection struct {
	DomainNeeded bool `json:"domainNeeded,omitempty" ubus:"domainneeded,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DHCPConfig{}, &DHCPConfigList{})
}
