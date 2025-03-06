# Requirements

Technical requirements for the project.

## Scaffolding
#### [operator-sdk](https://pkg.go.dev/github.com/operator-framework/operator-sdk)
Operator SDK is one of the most common scaffolding tools used for k8s operators and has built-in support 
for lots of useful things like automated testing and publishing on Artifact Hub.

## Libraries
#### [controller-runtime](https://pkg.go.dev/sigs.k8s.io/controller-runtime)
Most widely used controller framework, supported by the upstream Kubernetes project itself.

#### [rpc](https://pkg.go.dev/github.com/ethereum/go-ethereum/rpc)
A basic implementation of JSON-RPC 2.0. This will be used to make calls to the running instance of OpenWRT and execute [uci](https://openwrt.org/docs/techref/uci)
system calls via [ubus](https://openwrt.org/docs/techref/ubus) to modify the router's configuration. If any additional
RPC servers are needed beyond the built-in ones, they will be written in Go and added to ubus via rpcd.