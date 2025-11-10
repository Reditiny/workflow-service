package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"workflow-service/config"
	"workflow-service/handlers"
	"workflow-service/temporal"
)

func main() {
	cfg := config.Load()

	tc, err := temporal.NewClient(cfg)
	if err != nil {
		log.Fatalf("failed to init temporal client: %v", err)
	}
	defer tc.Close()

	go func() {
		if err := temporal.StartWorker(cfg, tc); err != nil {
			log.Fatalf("failed to start worker: %v", err)
		}
	}()

	r := gin.Default()

	api := r.Group("/v1")
	handlers.RegisterWorkflowRoutes(api, tc)
	handlers.RegisterScheduleRoutes(api, tc)

	log.Printf("🚀 Server running on %s", cfg.ServerAddr)
	if err := r.Run(cfg.ServerAddr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
