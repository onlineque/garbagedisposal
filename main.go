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
			age := time.Now().Sub(pods[pod].CreationTimestamp.Time)
			status := pods[pod].Status.Phase
			log.Printf("Terminating pod %s from %s namespace (%v), status: %s\n", podName, namespace, age,
				status)
			err := k8sfunctions.TerminatePod(clientset, namespace, podName)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
