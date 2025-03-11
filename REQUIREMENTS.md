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
There should be one CustomResource for each config file in `/etc/config` (can also be shown with `ubus call uci configs`). By default, these are:
- dhcp
- dropbear
- firewall
- luci
- network
- rpcd
- system
- ubootenv
- ucitrack
- uhttpd
- wireless

There should also be one for `/etc/board.json`.

#### Controller Logic
At first boot, a uci session must be established which will be used for all queries and updates. The configs should be read from the board as they are
and used to create the CustomResource instances. From there, controller-runtime will be used to execute logic based on what changes
are detected.