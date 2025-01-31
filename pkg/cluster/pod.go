package cluster

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodExists checks if a Kubernetes pod exists in the specified namespace
func PodExists(podName string, namespace string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	return err == nil
}

// checkPodForContainer checks if a specific container exists within a given pod// checkPodForContainer checks if a specific container exists within a given pod
func CheckPodForContainer(podName string, containerName string, namespace string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	for _, container := range pod.Spec.Containers {
		if container.Name == containerName {
			return true
		}
	}

	return false
}

// CheckContainerForImage checks if a specific container in a pod is running a given image
func CheckContainerForImage(podName string, containerName string, namespace string, imageName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	for _, container := range pod.Spec.Containers {
		if container.Name == containerName && container.Image == imageName {
			return true
		}
	}

	return false
}
