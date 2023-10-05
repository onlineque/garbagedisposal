package k8sfunctions

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"slices"
)

func InitAPIAccess() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func GetPods(clientset *kubernetes.Clientset, namespace string, statusList []string) ([]v1.Pod, error) {
	var result []v1.Pod

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, p := range pods.Items {
		if slices.Contains(statusList, string(p.Status.Phase)) {
			// if string(p.Status.Phase) == status {
			result = append(result, p)
		}
	}

	return result, nil
}

func TerminatePod(clientset *kubernetes.Clientset, namespace string, pod string) error {
	return clientset.CoreV1().Pods(namespace).Delete(context.TODO(), pod, metav1.DeleteOptions{})
}
