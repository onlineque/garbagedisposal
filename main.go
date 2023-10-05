package main

import (
	"fmt"
	"garbagedisposal/k8sfunctions"
	"log"
)

func main() {
	clientset, err := k8sfunctions.InitAPIAccess()
	if err != nil {
		log.Fatal("Error initializing API:", err)
	}

	pods, err := k8sfunctions.GetPods(clientset, "", "Succeeded")
	if err != nil {
		log.Fatal(err)
	}

	for pod := range pods {
		fmt.Printf("%s - %s\n", pods[pod].ObjectMeta.Namespace, pods[pod].ObjectMeta.Name)
	}
}
