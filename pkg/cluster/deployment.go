package cluster

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateDeployment creates a deployment in the specified namespace
func CreateDeployment(namespace string, deploymentName string, deployImage string, replicas int32) error {
	clientset, err := getKubernetesClient()
	if err != nil {
		return err
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": deploymentName},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": deploymentName},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: deployImage,
						},
					},
				},
			},
		},
	}

	_, err = clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	return err
}

// CheckDeployment checks if a deployment exists in the specified namespace
func CheckDeployment(namespace string, deploymentName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	return err == nil
}

// CheckReplicaCount checks if a deployment has the specified number of replicas
func CheckReplicaCount(namespace string, deploymentName string, expectedReplicas int32) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return false
	}

	return *deployment.Spec.Replicas == expectedReplicas
}
