package cluster

import (
	"context"
	"fmt"

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

// CheckPodRunning checks if a Kubernetes pod is running
func CheckPodRunning(podName string, namespace string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	return pod.Status.Phase == "Running"
}

// CheckPodForContainer checks if a specific container exists within a given pod
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

// CheckPodsNode retrieves the node on which a pod is running
func CheckPodsNode(namespace string, podName string) (string, bool) {
	clientset, err := getKubernetesClient()
	if err != nil {
		return "", false
	}

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return "", false
	}

	return pod.Spec.NodeName, true
}

// CheckContainerEnv determines if a container has a specific environment variable set to a specified value
func CheckContainerEnv(namespace string, podName string, containerName string, envName string, envValue string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	for _, container := range pod.Spec.Containers {
		fmt.Println(container.Name)
		if container.Name == containerName {
			fmt.Println("env.Name", envName, "env.Value", envValue)
			for _, env := range container.Env {
				if env.Name == envName && env.Value == envValue {
					return true
				}
			}
		}
	}

	return false
}
