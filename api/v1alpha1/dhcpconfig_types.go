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
	"github.com/daimonaslabs/go-ubus-rpc/pkg/ubus/uci"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make generate" to regenerate code after modifying this file

// DHCPConfigSpec defines the desired state of DHCPConfig
type DHCPConfigSpec struct {
	BootSections        []BootSection        `json:"bootSections,omitempty"`
	CircuitIDSections   []CircuitIDSection   `json:"circuitIDSections,omitempty"`
	DHCPSections        []DHCPSection        `json:"dhcpSections,omitempty"`
	DnsmasqSections     []DnsmasqSection     `json:"dnsmasqSections,omitempty"`
	HostSections        []HostSection        `json:"hostSections,omitempty"`
	HostRecordSections  []HostRecordSection  `json:"hostRecordSections,omitempty"`
	MACSections         []MACSection         `json:"macSections,omitempty"`
	OdhcpdSections      []OdhcpdSection      `json:"odhcpdSections,omitempty"`
	RelaySections       []RelaySection       `json:"relaySections,omitempty"`
	RemoteIDSections    []RemoteIDSection    `json:"remoteIDSections,omitempty"`
	SubscrIDSections    []SubscrIDSection    `json:"SubscrIDSections,omitempty"`
	TagSections         []TagSection         `json:"tagSections,omitempty"`
	UserClassSections   []UserClassSection   `json:"userClassSections,omitempty"`
	VendorClassSections []VendorClassSection `json:"vendorClassSections,omitempty"`
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

// TODO
type BootSection struct {
	UCIConfigOptionsStatic `json:",inline"`
}

// TODO
type ClientSection struct {
	UCIConfigOptionsStatic `json:",inline"`
}

// TODO
type DHCPSection struct {
	UCIConfigOptionsStatic `json:",inline"`
}

type DnsmasqSection struct {
	UCIConfigOptionsStatic `json:",inline"`
	// List of IP addresses for queried domains. See the dnsmasq man page for syntax details.
	Address []string `json:"address,omitempty"`
	// Add the local domain as search directive in resolv.conf.
	// +optional
	AddLocalDomain uci.StringBool `json:"addLocalDomain,omitempty"`
	// Add A, AAAA, and PTR records for this router only on DHCP served LAN.
	// Enhanced function available since OpenWrt 18.06 with option AddLocalFQDN
	AddLocalHostname uci.StringBool `json:"addLocalHostname,omitempty"`
	// Add A, AAAA, and PTR records for this router only on DHCP served LAN.
	// 0: Disable.
	// 1: Hostname on Primary Address.
	// 2: Hostname on All Addresses.
	// 3: FDQN on All Addresses.
	// 4: iface.host.domain on All Addresses.
	// Available since OpenWrt 18.06
	AddLocalFQDN int `json:"addLocalFQDN,omitempty"`
	// Add the MAC address of the requester to DNS queries which are forwarded upstream; this may be used to do
	// DNS filtering by the upstream server.
	// The MAC address can only be added if the requester is on the same subnet as the dnsmasq server. Note that
	// the mechanism used to achieve this (an EDNS0 option) is not yet standardised, so this should be considered
	// experimental. Also note that exposing MAC addresses in this way may have security and privacy implications.
	// The string value must be either "base64" or "text".
	AddMAC string `json:"addMAC,omitempty"`
	// Labels WAN interfaces like add_local_fqdn instead of your ISP assigned default which may be
	// obscure. WAN is inferred from config dhcp sections with option ignore 1 set, so they do not
	// need to be named WAN.
	// Available since OpenWrt 18.06
	AddWANFQDN int `json:"addWANFQDN,omitempty"`
	// Additional host files to read for serving DNS responses. Syntax in each file is the same as /etc/hosts.
	AddnHosts []string `json:"addnHosts,omitempty"`
	// Expose additional filesystem paths to the jailed dnsmasq process. This is useful in the case of manually
	// configured includes in the configuration file or symlinks pointing outside of the exposed paths as used,
	// for example, by an ad blocker or other name-banning package.
	AddnMount []string `json:"addnMount,omitempty"`
	// By default, when dnsmasq has more than one upstream server available, it will send queries to just one
	// server. Setting this parameter forces dnsmasq to send all queries to all available servers. The reply
	// from the server which answers first will be returned to the original requeser.
	AllServers uci.StringBool `json:"allServers,omitempty"`
	// Force dnsmasq into authoritative mode. This speeds up DHCP leasing. Used if this is the only server on
	// the network.
	Authoritative uci.StringBool `json:"authoritative,omitempty"`
	// IP addresses to convert into NXDOMAIN responses (to counteract “helpful” upstream DNS servers that never
	// return NXDOMAIN).
	BogusNXDOMAIN []uci.IP `json:"bogusNXDOMAIN,omitempty"`
	// Reject reverse lookups to private IP ranges where no corresponding entry exists in /etc/hosts.
	BogusPriv uci.StringBool `json:"bogusPriv,omitempty"`
	// When set to 0, use each network interface's DNS address in the local /etc/resolv.conf. Normally, only
	// the loopback address is used, and all queries go through dnsmasq.
	CacheLocal uci.StringBool `json:"cacheLocal,omitempty"`
	// Size of dnsmasq query cache.
	CacheSize int `json:"cacheSize,omitempty"`
	// Directory with additional configuration files.
	ConfDir string `json:"confDir,omitempty"`
	// Enable DBus messaging for dnsmasq.
	// Standard builds of dnsmasq on OpenWrt do not include DBus support.
	DBus uci.StringBool `json:"dbus,omitempty"`
	// Specifies BOOTP options, in most cases just the file name. You can also use:
	// “$FILENAME, $TFTP_SERVER_NAME, $TFTP_IP_ADDRESS”.
	DHCPBoot string `json:"dhcpBoot,omitempty"`
	// Maximum number of concurrent connections.
	DNSForwardMax int `json:"dnsForwardMax,omitempty"`
	// Specify an external file with per host DHCP options.
	DHCPHostsFile string `json:"dhcpHostsFile,omitempty"`
	// Maximum number of DHCP leases.
	DHCPLeaseMax int `json:"dhcpLeaseMax,omitempty"`
	// Run a custom script upon DHCP lease add / renew / remove actions.
	DHCPScript string `json:"dhcpScript,omitempty"`
	// Validate DNS replies and cache DNSSEC data.
	// Requires the dnsmasq-full package. Please note that many applications now require DNSSEC to work properly,
	// e.g. Google apps on iOS like Gmail and Google Maps, and Windows Update and Windows Account activation on
	// Windows PCs.
	DNSSEC uci.StringBool `json:"dnssec,omitempty"`
	// Check the zones of unsigned replies to ensure that unsigned replies are allowed in those zones. This
	// protects against an attacker forging unsigned replies for signed DNS zones, but is slower and requires that
	// the nameservers upstream of dnsmasq are DNSSEC-capable.
	// Requires the dnsmasq-full package.
	// Caution: If you use this option on a device that doesn't have a hardware clock, DNS resolution may break
	// after a reboot of the device due to an incorrect system time.
	DNSSECCheckUnsigned uci.StringBool `json:"dnssecCheckUnsigned,omitempty"`
	// DNS domain handed out to DHCP clients.
	Domain string `json:"domain,omitempty"`
	// Tells dnsmasq never to forward queries for plain names, without dots or domain parts, to upstream
	// nameservers. If the name is not known from /etc/hosts or DHCP then a “not found” answer is returned.
	DomainNeeded uci.StringBool `json:"domainNeeded,omitempty"`
	// Specify the largest EDNS.0 UDP packet which is supported by the DNS forwarder.
	EDNSPacketMax int `json:"ednsPacketMax,omitempty"`
	// Enable the builtin TFTP server.
	EnableTFTP uci.StringBool `json:"enableTFTP,omitempty"`
	// Add the local domain part to names found in /etc/hosts.
	ExpandHosts uci.StringBool `json:"expandHosts,omitempty"`
	// Do not forward requests that cannot be answered by public name servers.
	// Make sure it is disabled if you need to resolve SRV records or use SIP phones.
	FilterWin2k uci.StringBool `json:"filterWin2k,omitempty"`
	// Do not resolve unqualifed local hostnames. Needs Domain to be set.
	FQDN uci.StringBool `json:"fqdn,omitempty"`
	// List of interfaces to listen on. If unspecified, dnsmasq will listen to all interfaces except those listed
	// in NotInterface. Note that dnsmasq listens on loopback by default.
	Interface []string `json:"interface,omitempty"`
	// The syntax is: `list ipset '/example.com/example.org/example_ipv4,example_ipv6'``
	// Requires the dnsmasq-full package.
	IPSet []string `json:"ipSet,omitempty"`
	// Store DHCP leases in this file.
	LeaseFile string `json:"leaseFile,omitempty"`
	// Listen only on the specified IP addresses. If unspecified, listen on IP addresses from each interface.
	ListenAddress []uci.IP `json:"listenAddress,omitempty"`
	// Look up DNS entries for this domain from /etc/hosts. This follows the same syntax as Server entries.
	// See the dnsmasq man page for more details.
	Local string `json:"local,omitempty"`
	// Choose IP address to match the incoming interface if multiple addresses are assigned to a host name in
	// /etc/hosts. Initially disabled, but still enabled in the config by default.
	LocaliseQueries uci.StringBool `json:"localiseQueries,omitempty"`
	// Accept DNS queries only from hosts whose address is on a local subnet, ie a subnet for which an interface
	// exists on the server.
	LocalService uci.StringBool `json:"localService,omitempty"`
	// Default TTL for locally authoritative answers.
	LocalTTL int `json:"localTTL,omitempty"`
	// Use dnsmasq as a local system resolver. Depends on the NoResolv and ResolvFile options.
	LocalUse uci.StringBool `json:"localUse,omitempty"`
	// Enables extra DHCP logging; logs all the options sent to the DHCP clients and the tags used to determine
	// them.
	LogDHCP uci.StringBool `json:"logDHCP,omitempty"`
	// Set the facility to which dnsmasq will send syslog entries. See the dnsmasq man page for available
	// facilities.
	LogFacility string `json:"logFacility,omitempty"`
	// Log the results of DNS queries, dump cache on SIGUSR1, include requesting IP.
	LogQueries uci.StringBool `json:"logQueries,omitempty"`
	// Set the maximum TTL of DNS answers, even when the TTL in the answer is higher.
	MaxCacheTTL int `json:"maxCacheTTL,omitempty"`
	// Dnsmasq picks random ports as source for outbound queries. When this option is given, the ports used
	// will always be smaller than or equal to the specified MaxPort value (max valid value 65535). Useful for
	// systems behind firewalls.
	// See also MinPort.
	MaxPort int `json:"maxPort,omitempty"`
	// Limit the TTL in the DNS answer to this value.
	MaxTTL int `json:"maxTTL,omitempty"`
	// Set the minimum TTL of DNS answers, even when the TTL in the answer is lower.
	MinCacheTTL int `json:"minCacheTTL,omitempty"`
	// Dnsmasq picks random ports as source for outbound queries. When this option is given, the ports used
	// will always be larger than or equal to the specified MinPort value (min valid value 1024). Useful for
	// systems behind firewalls.
	// See also MaxPort.
	MinPort int `json:"minPort,omitempty"`
	// Don't daemonize the dnsmasq process.
	NoDaemon uci.StringBool `json:"noDaemon,omitempty"`
	// Don't read DNS names from /etc/hosts.
	NoHosts uci.StringBool `json:"noHosts,omitempty"`
	// Disable caching of negative “no such domain” responses.
	NoNegCache uci.StringBool `json:"noNegCache,omitempty"`
	// By default dnsmasq checks if an IPv4 address is in use before allocating it to a host by sending ICMP
	// echo request (aka ping) to the address in question. This parameter allows to disable this check.
	NoPing uci.StringBool `json:"noPing,omitempty"`
	// Don't read upstream servers from /etc/resolv.conf which is linked to resolvfile by default.
	NoResolv uci.StringBool `json:"noResolv,omitempty"`
	// Bind only configured interface addresses, instead of the wildcard address.
	NoWildcard uci.StringBool `json:"noWildcard,omitempty"`
	// Interfaces dnsmasq should not listen on.
	NotInterface []string `json:"notInterface,omitempty"`
	// Listening port for DNS queries, disables DNS server functionality if set to 0.
	Port int `json:"port,omitempty"`
	// Use a fixed port for outbound DNS queries.
	QueryPort int `json:"queryPort,omitempty"`
	// Suppress logging of the routine operation of DHCP. Errors and problems will still be logged.
	QuietDHCP uci.StringBool `json:"quietDHCP,omitempty"`
	// Enable DHCPv4 Rapid Commit (fast address assignment) See RFC 4039.
	RapidCommit uci.StringBool `json:"rapidCommit,omitempty"`
	// Read static lease entries from /etc/ethers, re-read on SIGHUP.
	ReadEthers uci.StringBool `json:"readEthers,omitempty"`
	// Enables DNS rebind attack protection by discarding upstream RFC1918 responses.
	RebindProtection uci.StringBool `json:"rebindProtection,omitempty"`
	// Allows upstream 127.0.0.0/8 responses, required for DNS based blacklist services, only takes effect if
	// rebind protection is enabled.
	RebindLocalhost uci.StringBool `json:"rebindLocalhost,omitempty"`
	// List of domains to allow RFC1918 responses for, only takes effect if rebind protection is enabled.
	// The correct syntax is: `list rebind_domain '/example.com/'`
	RebindDomain []string `json:"rebindDomain,omitempty"`
	// Specifies an alternative resolv file.
	ResolvFile string `json:"resolvFile,omitempty"`
	// List of network range with a DNS server to forward reverse DNS requests to. See the dnsmasq man page
	// for syntax details.
	RevServer []string `json:"revServer,omitempty"`
	// Dnsmasq is designed to choose IP addresses for DHCP clients using a hash of the client's MAC address.
	// This normally allows a client's address to remain stable long-term, even if the client sometimes allows
	// its DHCP lease to expire. In this default mode IP addresses are distributed pseudo-randomly over the
	// entire available address range. There are sometimes circumstances (typically server deployment) where
	// it is more convenient to have IP addresses allocated sequentially, starting from the lowest available
	// address, and setting this parameter enables this mode. Note that in the sequential mode, clients which
	// allow a lease to expire are much more likely to move IP address; for this reason it should not be
	// generally used.
	SequentialIP uci.StringBool `json:"sequentialIP,omitempty"`
	// List of DNS servers to forward requests to. See the dnsmasq man page for syntax details.
	Server []string `json:"server,omitempty"`
	// Specify upstream servers directly. If one or more optional domains are given, that server is used only
	// for those domains and they are queried only using the specified server.
	// Syntax is `server=/*.mydomain.tld/192.168.100.1` or see the dnsmasq man page for details.
	ServerList string `json:"serverList,omitempty"`
	// Obey order of DNS servers in /etc/resolv.conf.
	StrictOrder uci.StringBool `json:"strictOrder,omitempty"`
	// Specifies the TFTP root directory.
	TFTPRoot string `json:"tftpRoot,omitempty"`
}

// TODO
type HostSection struct {
	UCIConfigOptionsStatic `json:",inline"`
}

// TODO
type IPSetSection struct {
	UCIConfigOptionsStatic `json:",inline"`
}

// TODO
type RelaySection struct {
	UCIConfigOptionsStatic `json:",inline"`
}

func init() {
	SchemeBuilder.Register(&DHCPConfig{}, &DHCPConfigList{})
}
