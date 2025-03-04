# Requirements

Technical requirements for the project.

## Libraries

#### [operator-sdk](https://sdk.operatorframework.io/docs/)
Since this is a Kubernetes controller, it makes sense to use a framework to help create it. Operator SDK is one of the
most common ones used and has built-in support for automated testing and publishing on Artifact Hub.

#### [gRPC](https://pkg.go.dev/google.golang.org/grpc)
This will be used to make calls to the running instance of OpenWRT and execute [uci](https://openwrt.org/docs/techref/uci)
system calls via [ubus](https://openwrt.org/docs/techref/ubus) to modify the router's configuration. If any additional
RPC servers are needed beyond the built-in ones, they will be written in Python and added to ubus via rpcd.