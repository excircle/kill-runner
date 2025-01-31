package cluster

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
