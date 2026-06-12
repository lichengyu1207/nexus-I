package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Worker starting...")

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Worker heartbeat - processing tasks...")
	}
}
