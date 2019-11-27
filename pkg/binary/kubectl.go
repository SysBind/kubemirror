// package binary provides wrappers around k8s related binaries (e.g: kubectl, docker, helm)
// handles executing and parsing outputs
package binary

import (
	"os"
	"os/exec"
)

type K8sObjectType int

const (
	Pod K8sObjectType = iota
	ReplicaSet
	DeamonSet
	Deployment
	StatefulSet
	Job
	CronJob
	Entpoint
	Service
	Ingress
	ConfigMap
	Sercret
	PersistentVolumeClaim
	PersistentVolume
)

type Kubectl struct {
	path       string
	kubeconfig string
}

func (k *Kubectl) Get(object K8sObjectType) error {
	cmd := exec.Command(k.path, "get")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (obj K8sObjectType) String() string {

	names := [...]string{
		"pod",
		"replicaset",
		"deamonset",
		"deployment",
		"statefulset",
		"job",
		"cronjob",
		"entpoint",
		"service",
		"ingress",
		"configmap",
		"secret",
		"persistentvolumeclaim",
		"persistentvolume"}

	if obj < Pod || obj > PersistentVolume {
		return "Unknown"
	}
	return names[obj]
}
