package main

import (
	"fmt"
	"garbagedisposal/k8sfunctions"
	"log"
	"os"
	"time"
)

func main() {
	succeededAge := os.Getenv("SUCCEEDED_AGE")
	if succeededAge == "" {
		succeededAge = "15m"
	}
	succeededAgeDuration, err := time.ParseDuration(succeededAge)
	if err != nil {
		log.Fatal("SUCCEEDED_AGE parameter is invalid:", err)
	}

	failedAge := os.Getenv("FAILED_AGE")
	if failedAge == "" {
		failedAge = "60m"
	}
	failedAgeDuration, err := time.ParseDuration(failedAge)
	if err != nil {
		log.Fatal("FAILED_AGE parameter is invalid:", err)
	}

	fmt.Printf("Garbage Disposal\nParameters:\n - SUCCEEDED_AGE: %s\n - FAILED_AGE: %s\n", succeededAge,
		failedAge)

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
			if (status == "Succeeded" && age >= succeededAgeDuration) || (status == "Failed" && age >= failedAgeDuration) {
				log.Printf("Terminating pod %s from %s namespace (%v), status: %s\n", podName, namespace, age,
					status)
				err := k8sfunctions.TerminatePod(clientset, namespace, podName)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
