package main

import (
	"log"
	"time"

	"github.com/VishwasMallikarjuna/request-counter/core/config"
	"github.com/VishwasMallikarjuna/request-counter/core/counter"
	"github.com/VishwasMallikarjuna/request-counter/core/server"
)

func main() {
	cfg := config.New()

	memoryCounter, err := counter.LoadFromFile(cfg.Filename)
	if err != nil {
		log.Printf("Cannot read counter from %s: %s", cfg.Filename, err)
		memoryCounter = counter.NewMemoryCounter()
	}

	tick := time.NewTicker(cfg.SaveInterval)
	defer tick.Stop()

	go func() {
		for range tick.C {
			if err := counter.SaveToFile(cfg.Filename, memoryCounter); err != nil {
				log.Printf("Error saving counter: %s", err)
			}
		}
	}()

	handler := server.NewHandler(memoryCounter)

	srv := server.New(cfg.Port)
	srv.Register(handler)

	if err := srv.Run(); err != nil {
		log.Fatalf("Running server failed: %s", err)
	}
}
