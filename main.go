package main

import (
	"garbagedisposal/k8sfunctions"
	"log"
	"time"
)

func main() {
	clientset, err := k8sfunctions.InitAPIAccess()
	if err != nil {
		log.Fatal("Error initializing API:", err)
	}

	statusList := []string{"Succeeded", "Failed"}

	for {
		time.Sleep(1 * time.Minute)

		pods, err := k8sfunctions.GetPods(clientset, "", statusList)
		if err != nil {
			log.Fatal(err)
		}

		for pod := range pods {
			namespace := pods[pod].ObjectMeta.Namespace
			podName := pods[pod].ObjectMeta.Name
			age := pods[pod].CreationTimestamp.Time
			status := pods[pod].Status
			log.Printf("Terminating pod %s - %s (%v), status: \n", namespace, podName, age, status)
			err := k8sfunctions.TerminatePod(clientset, namespace, podName)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
