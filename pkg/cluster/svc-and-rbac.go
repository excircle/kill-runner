package cluster

import (
	"context"

	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CheckSaAcctExists checks if a service account exists in a given namespace
func CheckSaAcctExists(namespace string, svcAcctName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.CoreV1().ServiceAccounts(namespace).Get(context.TODO(), svcAcctName, metav1.GetOptions{})
	return err == nil
}

// CheckRoleExists checks if a role exists in a given namespace
func CheckRoleExists(namespace string, roleName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.RbacV1().Roles(namespace).Get(context.TODO(), roleName, metav1.GetOptions{})
	return err == nil
}

// CheckRoleBindingExists checks if a role binding exists in a given namespace
func CheckRoleBindingExists(namespace string, roleBindingName string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.RbacV1().RoleBindings(namespace).Get(context.TODO(), roleBindingName, metav1.GetOptions{})
	return err == nil
}

func CanI(namespace string, svcAcctName string, action string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	sar := &authorizationv1.SubjectAccessReview{
		Spec: authorizationv1.SubjectAccessReviewSpec{
			User: "system:serviceaccount:" + namespace + ":" + svcAcctName,
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Namespace: namespace,
				Verb:      "create",
				Resource:  action,
			},
		},
	}

	response, err := clientset.AuthorizationV1().SubjectAccessReviews().Create(context.TODO(), sar, metav1.CreateOptions{})
	if err != nil {
		return false
	}

	return response.Status.Allowed
}
