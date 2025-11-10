package temporal

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

// ExampleWorkflow is a sample Temporal workflow that waits for a signal before proceeding.
func ExampleWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started")

	signalCh := workflow.GetSignalChannel(ctx, "my-signal")
	var signalValue string

	logger.Info("Waiting for signal...")
	signalCh.Receive(ctx, &signalValue)

	logger.Info("Received signal", "signal", signalValue)
	_ = workflow.Sleep(ctx, time.Second*2)
	logger.Info("Workflow finished")
	return nil
}
