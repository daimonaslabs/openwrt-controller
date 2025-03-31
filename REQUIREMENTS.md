# Requirements

Technical requirements for the project.

## Scaffolding
#### [operator-sdk](https://pkg.go.dev/github.com/operator-framework/operator-sdk)
Operator SDK is one of the most common scaffolding tools used for k8s operators and has built-in support 
for lots of useful things like automated testing and publishing on Artifact Hub.

## Libraries
#### [controller-runtime](https://pkg.go.dev/sigs.k8s.io/controller-runtime)
Most widely used controller framework, supported by the upstream Kubernetes project itself.

#### [go-ubus-rpc](https://github.com/daimonaslabs/go-ubus-rpc)
This will be used to make calls to the running instance of OpenWRT and execute [uci](https://openwrt.org/docs/techref/uci)
system calls via [ubus](https://openwrt.org/docs/techref/ubus) to modify the router's configuration. If any additional
RPC servers are needed beyond the built-in ones, they will be written in Go and added to ubus via [rpcd](https://openwrt.org/docs/techref/rpcd).

## Implementation
#### API Guidelines
There should be one CustomResource for each section type available in each config file in `/etc/config` (can also be shown with `ubus call uci configs`).
By default, these are:
- dhcp
    - boot
    - dnsmasq
    - dhcp
    - host
    - ipset
    - relay
- dropbear
- firewall
    - defaults
    - forwardings
    - include
    - ipsets
    - redirects
    - rules
    - zones
- luci
- network
- rpcd
- system
- ubootenv
- ucitrack
- uhttpd
    - uhttpd
- wireless
    - wifi-device
    - wifi-iface
    - wifi-station
    - wifi-vlan

Additionally, there should be one CustomResource for every service listed by `ubus call service list`,
one for the ubus `session` path, one for the ubus `hostapd` path, and a generic one for the ubus 
`file` path,

These objects should be constructed of embedded structs mirroring the ubus uci config `.type` field that they are, e.g:.

```
$ ubus call uci get '{"config":"network"}'                                                                                                                      
{
    "values": {
        "loopback": {
            ".anonymous": false,
            ".type": "interface", # one subtype
            ".name": "loopback",
            ".index": 0,
            "device": "lo",
            "proto": "static",
            "ipaddr": "127.0.0.1",
            "netmask": "255.0.0.0"
        },
        "globals": {
            ".anonymous": false,
            ".type": "globals", # another subtype
            ".name": "globals",
            ".index": 1,
            "ula_prefix": "fdcc:72bd:c25c::/48"
        },
        ...
    }
}
```

#### Controller Logic
The controller will require an IP address and a username/password to be passed to it as configuration values in order to connect
to the router. At first boot, a uci session must be established which will be used for all queries and updates. The configs should
be read from the board and used to create the CustomResource instances. From there, controller-runtime will manage state.


## Additional Resources
- [OpenWRT Docs](https://openwrt.org/docs/start)