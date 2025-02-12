package cluster

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CheckContainerForVolume verifies that a given container, inside a given pod, is mounting a volume with a given name, within a given namespace
func CheckContainerForVolume(namespace string, pod string, container string, volume string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	podSpec, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	if err != nil {
		return false
	}

	for _, c := range podSpec.Spec.Containers {
		if c.Name == container {
			for _, vm := range c.VolumeMounts {
				if vm.Name == volume {
					return true
				}
			}
		}
	}

	return false
}
