package handlers

import (
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
	"log"
	"net/http"
	"time"
	"workflow-service/models"
	"workflow-service/temporal"
)

func RegisterWorkflowRoutes(rg *gin.RouterGroup, c client.Client) {
	h := &WorkflowHandler{
		Cli: c,
	}

	wf := rg.Group("/workflows")
	{
		wf.POST("/run", h.StartWorkflow)
		wf.GET("/:workflow_id", h.DescribeWorkflow)
		wf.POST("/:workflow_id/cancel", h.CancelWorkflow)
		wf.POST("/:workflow_id/signals", h.SignalWorkflow)
		wf.POST("/:workflow_id/batch-signal", h.BatchSignal)
		wf.GET("", h.ListWorkflows)
	}
}

type WorkflowHandler struct {
	Cli client.Client
}

/*
StartWorkflow
Input: workflow_type, workflow_id (optional), task_queue, args_json, timeout
Internal: validate workflow_type, start workflow via temporal client
Output: workflow_id, run_id, started_at
*/
func (h *WorkflowHandler) StartWorkflow(c *gin.Context) {
	var (
		req  models.StartWorkflowRequest
		resp models.StartWorkflowResponse
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	opts := client.StartWorkflowOptions{
		ID:        req.WorkflowID,
		TaskQueue: "default-task-queue",
	}

	//workflow, ok := temporal.Workflows[req.WorkflowType]
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown workflow type"})
	//	return
	//}

	wr, err := h.Cli.ExecuteWorkflow(c, opts, temporal.ExampleWorkflow, req.ArgsJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp.WorkflowID = wr.GetID()
	resp.RunID = wr.GetRunID()
	resp.CreatedAt = time.Now()
	resp.StartAt = time.Now()
	c.JSON(http.StatusOK, models.NewResponseMessage("Workflow started", resp))
}

func (h *WorkflowHandler) DescribeWorkflow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DescribeWorkflow placeholder"})
}

func (h *WorkflowHandler) CancelWorkflow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CancelWorkflow placeholder"})
}

/*
SignalWorkflow
Input: workflow_id, signal_name, signal_payload
Internal: send signal to running workflow (e.g. webhook events, user consent actions)
Output: workflow_id, signal_sent_at
*/
func (h *WorkflowHandler) SignalWorkflow(c *gin.Context) {
	var (
		req  models.SignalWorkflowRequest
		resp models.SignalWorkflowResponse
	)

	workflowID := c.Param("workflow_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.Cli.SignalWorkflow(c, workflowID, req.RunID, req.SignalName, req.SignalPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp.WorkflowID = workflowID
	resp.RunID = req.RunID
	resp.SignalName = req.SignalName
	resp.SignalSentAt = time.Now()
	c.JSON(http.StatusOK, models.NewResponseMessage("Signal dispatched", resp))
}

func (h *WorkflowHandler) BatchSignal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SignalWorkflow placeholder"})
}

func (h *WorkflowHandler) ListWorkflows(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ListWorkflows placeholder"})
}
