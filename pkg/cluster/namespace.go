package cluster

import (
	"context"
	"fmt"
	"log"

	"github.com/excircle/kill-runner/pkg/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// getKubernetesClient creates and returns a Kubernetes clientset
func getKubernetesClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeConfigFile := utils.GetGlobalKubeConfig()
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigFile)
		if err != nil {
			utils.LogEvent(2, "Failed to load kubeconfig")
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

// NamespaceExists checks if a Kubernetes namespace exists
func NamespaceExists(name string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return false
	}

	return true
}

// CreateNamespace creates a Kubernetes namespace
func CreateNamespace(name string, question string) error {
	clientset, err := getKubernetesClient()
	if err != nil {
		return err
	}

	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create namespace: %v", err)
	}

	fmt.Printf("[%s] Namespace %s created successfully\n", question, name)
	return nil
}

// DestroyNamespace deletes a Kubernetes namespace
func DestroyNamespace(name string, question string) error {
	clientset, err := getKubernetesClient()
	if err != nil {
		return err
	}

	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("Failed to delete namespace %s: %v", name, err)
	}

	fmt.Printf("[%s] Namespace %s deleted successfully\n", question, name)
	return nil
}
