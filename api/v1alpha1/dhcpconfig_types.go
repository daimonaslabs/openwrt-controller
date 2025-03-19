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
	"net/netip"

	ubus "github.com/daimonaslabs/go-ubus-rpc"
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
	AddLocalDomain      ubus.UbusBool `json:"addLocalDomain,omitempty" ubus:"add_local_domain,omitempty"`
	AddLocalHostname    ubus.UbusBool `json:"addLocalHostname,omitempty" ubus:"add_local_hostname,omitempty"`
	AddLocalFQDN        int           `json:"addLocalFQDN,omitempty" ubus:"add_local_fqdn,omitempty"`
	AddWANFQDN          int           `json:"addWANFQDN,omitempty" ubus:"add_wan_fqdn,omitempty"`
	AddnHosts           []string      `json:"addnHosts,omitempty" ubus:"addnhosts,omitempty"`
	AddnMount           []string      `json:"addnMount,omitempty" ubus:"addnmount,omitempty"`
	Authoritative       ubus.UbusBool `json:"authoritative,omitempty" ubus:"authoritative,omitempty"`
	BogusNXDOMAIN       []netip.Addr  `json:"bogusNXDOMAIN,omitempty" ubus:"bogusnxdomain,omitempty"`
	BogusPriv           ubus.UbusBool `json:"bogusPriv,omitempty" ubus:"boguspriv,omitempty"`
	CacheLocal          ubus.UbusBool `json:"cacheLocal,omitempty" ubus:"cachelocal,omitempty"`
	CacheSize           int           `json:"cacheSize,omitempty" ubus:"cachesize,omitempty"`
	DBus                ubus.UbusBool `json:"dbus,omitempty" ubus:"dbus,omitempty"`
	DHCPBoot            string        `json:"dhcpBoot,omitempty" ubus:"dhcp_boot,omitempty"`
	DHCPHostsFile       string        `json:"dhcpHostsFile,omitempty" ubus:"dhcphostsfile,omitempty"`
	DHCPLeaseMax        int           `json:"dhcpLeaseMax,omitempty" ubus:"dhcpleasemax,omitempty"`
	DNSForwardMax       int           `json:"dnsForwardMax,omitempty" ubus:"dnsforwardmax,omitempty"`
	Domain              string        `json:"domain,omitempty" ubus:"domain,omitempty"`
	DomainNeeded        ubus.UbusBool `json:"domainNeeded,omitempty" ubus:"domainneeded,omitempty"`
	DNSSEC              ubus.UbusBool `json:"dnssec,omitempty" ubus:"dnssec,omitempty"`
	DNSSECCheckUnsigned ubus.UbusBool `json:"dnssecCheckUnsigned,omitempty" ubus:"dnsseccheckunsigned,omitempty"`
	EDNSPacketMax       int           `json:"ednsPacketMax,omitempty" ubus:"ednspacket_max,omitempty"`
	EnableTFTP          ubus.UbusBool `json:"enableTFTP,omitempty" ubus:"enable_tftp,omitempty"`
	ExpandHosts         ubus.UbusBool `json:"expandHosts,omitempty" ubus:"expandhosts,omitempty"`
	FilterWin2k         ubus.UbusBool `json:"filterWin2k,omitempty" ubus:"filterwin2k,omitempty"`
	FQDN                ubus.UbusBool `json:"fqdn,omitempty" ubus:"fqdn,omitempty"`
	ListenAddress       []netip.Addr  `json:"listenAddress,omitempty" ubus:"listen_address,omitempty"`
	Interface           []string      `json:"interface,omitempty" ubus:"interface,omitempty"`
	NotInterface        []string      `json:"notInterface,omitempty" ubus:"notinterface,omitempty"`
	IPSet               []string      `json:"ipSet,omitempty" ubus:"ipset,omitempty"`
	LeaseFile           string        `json:"leaseFile,omitempty" ubus:"leasefile,omitempty"`
	Local               string        `json:"local,omitempty" ubus:"local,omitempty"`
	LocaliseQueries     ubus.UbusBool `json:"localiseQueries,omitempty" ubus:"localise_queries,omitempty"`
	LocalService        ubus.UbusBool `json:"localService,omitempty" ubus:"localservice,omitempty"`
	LocalTTL            int           `json:"localTTL,omitempty" ubus:"local_ttl,omitempty"`
	LocalUse            ubus.UbusBool `json:"localUse,omitempty" ubus:"localuse,omitempty"`
	LogFacility         string        `json:"logFacility,omitempty" ubus:"logfacility,omitempty"`
	LogQueries          ubus.UbusBool `json:"logQueries,omitempty" ubus:"logqueries,omitempty"`
	NoDaemon            ubus.UbusBool `json:"noDaemon,omitempty" ubus:"nodaemon,omitempty"`
	NoHosts             ubus.UbusBool `json:"noHosts,omitempty" ubus:"nohosts,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DHCPConfig{}, &DHCPConfigList{})
}
