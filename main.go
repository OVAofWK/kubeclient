package main

import updateImage "kubeclient/update/image"

func main() {
	update := updateImage.NewUpdateImage()
	// pods.GetPods(clientset)
	// pods.StartTest()
	err := update.UpdateDeploymentsImage("default", "nginx", "nginx", "redis:6.0.16")
	if err != nil {
		panic(err)
	}
}
