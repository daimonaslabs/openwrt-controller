# openwrt-controller
A Kubernetes controller to manage a router running OpenWRT

### Bootstrapping
```
$ operator-sdk init --domain daimonaslabs.io --repo github.com/daimonaslabs/openwrt-controller --owner "Daimonas Labs"
$ operator-sdk create api --group openwrt --version v1alpha1 --kind DHCPConfig --resource --controller
```