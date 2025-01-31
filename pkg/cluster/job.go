package cluster

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CheckJobExists(ns string, job string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	_, err = clientset.BatchV1().Jobs(ns).Get(context.TODO(), job, metav1.GetOptions{})
	return err == nil
}

func CheckJobSpec(ns string, job string, parallelism int32, completions int32, label string) (bool, string) {
	var issue string

	clientset, err := getKubernetesClient()
	if err != nil {
		return false, issue
	}

	j, err := clientset.BatchV1().Jobs(ns).Get(context.TODO(), job, metav1.GetOptions{})
	if err != nil {
		return false, issue
	}

	if *j.Spec.Parallelism != parallelism {
		issue = "parallelism"
		return false, issue
	}

	if *j.Spec.Completions != completions {
		issue = "completions"
		return false, issue
	}

	if j.Spec.Template.ObjectMeta.Labels["id"] != label {
		issue = "label"
		return false, issue
	}

	return true, issue
}

func CheckContainerJobsForImage(ns string, job string, img string) bool {
	clientset, err := getKubernetesClient()
	if err != nil {
		return false
	}

	jobs, err := clientset.BatchV1().Jobs(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false
	}

	for _, j := range jobs.Items {
		if j.Name == job {
			for _, container := range j.Spec.Template.Spec.Containers {
				if container.Image == img {
					return true
				}
			}
		}
	}

	return false
}
