package cluster

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// ExposePod exposes a Kubernetes pod in the specified namespace
func ExposePod(namespace string, podName string, svcName string, port int32, targetPort int32, exposeType string) error {
	clientset, err := getKubernetesClient()
	if err != nil {
		return err
	}

	svcType := corev1.ServiceTypeClusterIP
	if exposeType == "NodePort" {
		svcType = corev1.ServiceTypeNodePort
	} else if exposeType == "LoadBalancer" {
		svcType = corev1.ServiceTypeLoadBalancer
	}

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svcName,
			Namespace: namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": podName},
			Ports: []corev1.ServicePort{
				{
					Port:       port,
					TargetPort: intstr.FromInt(int(targetPort)),
				},
			},
			Type: svcType,
		},
	}

	_, err = clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	return err
}

func CheckService(namespace string, svcName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.CoreV1().Services(namespace).Get(context.TODO(), svcName, metav1.GetOptions{})
	return err == nil
}
