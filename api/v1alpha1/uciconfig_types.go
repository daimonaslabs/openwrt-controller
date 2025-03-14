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

type UCIConfigType string
var UCIConfigTypes = []string{'dhcp', 'dropbear', 'firewall', 'luci', 'network', 'rpcd', 'system', 'ubootenv', 'ucitrack', 'uhttpd', 'wireless'}
type UCIConfigSection struct {
	ConfigOptions map[string]string `json:"options,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UCIConfigSpec defines the desired state of UCIConfig
type UCIConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// maps to .type in actual ubus uci call
	ConfigType UCIConfigType string `json:"configType,omitempty"`

	// +optional
	// maps to .name in actual ubus uci call
	Section UCIConfigSection `json:"section,omitempty"`
}

// UCIConfigStatus defines the observed state of UCIConfig
type UCIConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// UCIConfig is the Schema for the uciconfigs API
type UCIConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UCIConfigSpec   `json:"spec,omitempty"`
	Status UCIConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UCIConfigList contains a list of UCIConfig
type UCIConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UCIConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UCIConfig{}, &UCIConfigList{})
}
