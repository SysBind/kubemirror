package copy

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

func PackSecrets(secrets *v1.SecretList) {
	for _, secret := range secrets.Items {
		fmt.Printf("Packing secret %s\n", secret.GetName())
	}
}
