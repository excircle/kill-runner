package utils

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func KubeConnect(kubeconfig string) error {
	// Try to load in-cluster configuration (if running inside a pod)
	config, err := rest.InClusterConfig()
	if err != nil {
		// If not in a pod, load kubeconfig from the default location
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			LogEvent(2, "Failed to load kubeconfig")
			return err
		}
	}

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %v", err)
	}

	// Attempt to list nodes to validate the connection
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list nodes: %v", err)
	}

	message := fmt.Sprintf("Successfully connected! Found %d node(s).\n", len(nodes.Items))
	LogEvent(0, message)
	return nil
}
