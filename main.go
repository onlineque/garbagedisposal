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
	err = k8sfunctions.GetPods(clientset, "", "Succeeded")
	if err != nil {
		log.Fatal(err)
	}
}
