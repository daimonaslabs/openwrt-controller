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

// Important: Run "make generate" to regenerate code after modifying this file

type DefaultsSections []ubus.DefaultsSection
type ForwardingSections []ubus.ForwardingSection
type RedirectSections []ubus.RedirectSection
type RuleSections []ubus.RuleSection
type ZoneSections []ubus.ZoneSection

// FirewallConfigSpec defines the desired state of FirewallConfig
type FirewallConfigSpec struct {
	DefaultsSections
	ForwardingSections
	RedirectSections
	RuleSections
	ZoneSections
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

func init() {
	SchemeBuilder.Register(&FirewallConfig{}, &FirewallConfigList{})
}
