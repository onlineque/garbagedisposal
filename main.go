package main

import (
	"garbagedisposal/k8sfunctions"
	"log"
)

func main() {
	clientset, err := k8sfunctions.InitAPIAccess()
	if err != nil {
		log.Fatal("Error initializing API:", err)
	}
	err = k8sfunctions.GetPods(clientset, "kube-system", "Completed")
	if err != nil {
		log.Fatal("Error getting list of pods: ", err)
	}
}
