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

	for {
		// once per minute
		time.Sleep(1 * time.Minute)

		pods, err := k8sfunctions.GetPods(clientset, "", "Succeeded")
		if err != nil {
			log.Fatal(err)
		}

		for pod := range pods {
			namespace := pods[pod].ObjectMeta.Namespace
			podName := pods[pod].ObjectMeta.Name
			log.Printf("%s - %s\n", pods[pod].ObjectMeta.Namespace, pods[pod].ObjectMeta.Name)
			err := k8sfunctions.TerminatePod(clientset, namespace, podName)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
