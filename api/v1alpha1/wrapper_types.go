package v1alpha1

import "net/netip"

// +k8s:deepcopy-gen=false
type Addr struct {
	netip.Addr
}

// This DeepCopyInto is a manually created deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addr) DeepCopyInto(out *Addr) {
	*out = *in
}

// This DeepCopy is a manually created deepcopy function, copying the receiver, creating a new Addr.
func (in *Addr) DeepCopy() *Addr {
	if in == nil {
		return nil
	}
	out := new(Addr)
	in.DeepCopyInto(out)
	return out
}
