package main

import (
	"log"
	"sync"
	"github.com/hasirciogli/multicloud-csi-driver/controller"
    "github.com/hasirciogli/multicloud-csi-driver/node"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // 2 server başlatacağımız için

	// Controller Server'ı başlat
	go func() {
		defer wg.Done()
		if err := controller.StartControllerServer(); err != nil {
			log.Fatalf("Failed to start controller server: %v", err)
		}
	}()

	// Node Server'ı başlat
	go func() {
		defer wg.Done()
		if err := node.StartNodeServer(); err != nil {
			log.Fatalf("Failed to start node server: %v", err)
		}
	}()

	// Her iki server'ı da bekle
	wg.Wait()
}
