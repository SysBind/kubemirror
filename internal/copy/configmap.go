package copy

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

func PackConfigMaps(cms *v1.ConfigMapList) {
	for _, cm := range cms.Items {
		fmt.Printf("Packing configmap %s\n", cm.GetName())
	}
}
