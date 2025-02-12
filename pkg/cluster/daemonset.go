package cluster

import (
	"context"

	//appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//corev1 "k8s.io/api/core/v1"
)

// DaemonSetExists checks if a DaemonSet exists in the specified namespace
func DaemonSetExists(namespace string, dsName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), dsName, metav1.GetOptions{})
	return err == nil
}

// DaemonSetUsesImage verifies if a given DaemonSet is using a specified image
func DaemonSetUsesImage(namespace string, dsName string, imageName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	daemonSet, err := clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), dsName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	for _, container := range daemonSet.Spec.Template.Spec.Containers {
		if container.Image == imageName {
			return true
		}
	}

	return false
}

// DaemonSetHasLabel verifies if a given DaemonSet has a specified label
func DaemonSetHasLabel(namespace string, dsName string, labelKey string, labelValue string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	daemonSet, err := clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), dsName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	if val, exists := daemonSet.Labels[labelKey]; exists && val == labelValue {
		return true
	}

	return false
}

// DaemonSetHasResourceRequests verifies if a given DaemonSet has specified CPU and memory requests
func DaemonSetHasResourceRequests(namespace string, dsName string, cpuRequest string, memoryRequest string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	daemonSet, err := clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), dsName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	for _, container := range daemonSet.Spec.Template.Spec.Containers {
		requests := container.Resources.Requests
		if requests.Cpu().String() == cpuRequest && requests.Memory().String() == memoryRequest {
			return true
		}
	}

	return false
}

// DaemonSetRunningOnAllNodes verifies if a given DaemonSet is running on all control plane and worker nodes
func DaemonSetRunningOnAllNodes(namespace string, dsName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false
	}

	daemonSet, err := clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), dsName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	expectedPods := len(nodes.Items)
	if int(daemonSet.Status.NumberAvailable) == expectedPods {
		return true
	}

	return false
}
