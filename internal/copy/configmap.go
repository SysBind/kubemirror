package copy

import (
	"fmt"
	k8sbinary "github.com/sysbind/kubemirror/pkg/binary"
	v1 "k8s.io/api/core/v1"
)

func (packer *Packer) PackConfigMaps(cms *v1.ConfigMapList) {
	for _, cm := range cms.Items {
		fmt.Printf("Packing configmap %s\n", cm.GetName())

		packer.kubectl.Get(k8sbinary.Pod)
	}
}
