package copy

import (
	k8sbinary "github.com/sysbind/kubemirror/pkg/binary"
)

type Packer struct {
	namespace string
	kubectl   k8sbinary.Kubectl
}
