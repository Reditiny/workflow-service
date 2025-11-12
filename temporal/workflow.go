package temporal

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

// WorkflowFunc defines the signature for Temporal workflow functions.
type WorkflowFunc func(ctx workflow.Context, input map[string]interface{}) error

var Workflows = map[string]WorkflowFunc{}

func init() {
	Workflows["ExampleWorkflow"] = ExampleWorkflow
}

// ExampleWorkflow is a sample Temporal workflow that waits for a signal before proceeding.
func ExampleWorkflow(ctx workflow.Context, input map[string]interface{}) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started")

	logger.Info("Activity to perform", "name", input["name"], "action", input["act"])

	signalCh := workflow.GetSignalChannel(ctx, "my-signal")
	var signalValue map[string]interface{}

	logger.Info("Waiting for signal...")
	signalCh.Receive(ctx, &signalValue)

	logger.Info("Received signal", "signal", signalValue)
	_ = workflow.Sleep(ctx, time.Second*2)
	logger.Info("Workflow finished")
	return nil
}
