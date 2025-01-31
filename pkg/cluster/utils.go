package cluster

import (
	"github.com/excircle/kill-runner/pkg/utils"
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
