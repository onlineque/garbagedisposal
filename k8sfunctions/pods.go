package k8sfunctions

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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

func GetPods(clientset *kubernetes.Clientset, namespace string, status string) error {
	fmt.Println(status)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	fmt.Println("pod count: ", len(pods.Items))

	for _, p := range pods.Items {
		if string(p.Status.Phase) == status {
			fmt.Printf("%s - %s\n", p.ObjectMeta.Namespace, p.ObjectMeta.Name)
		}
	}

	return nil
}
