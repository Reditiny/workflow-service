package temporal

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"workflow-service/config"
)

func StartWorker(cfg *config.Config, c client.Client) error {
	w := worker.New(c, cfg.TaskQueue, worker.Options{})

	w.RegisterWorkflow(ExampleWorkflow)
	log.Println("Worker started...")

	log.Printf("Temporal Worker started. Task Queue: %s", cfg.TaskQueue)
	return w.Run(worker.InterruptCh())
}
