package models

import (
	"time"
)

type ResponseMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseMessage(message string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Message: message,
		Data:    data,
	}
}

/*
StartWorkflowRequest example:

	{
	  "workflow_type": "DisbursementWorkflow",
	  "workflow_id": "f3b8f77c-9a2d-4f81-aaa8-1f1d6c39f512",
	  "args_json": {
		"user_id": "user_123",
		"target_bank_account": "bank_account_456",
		"advance_id": "adv_987",
		"amount": 150.0
	  },
	  "timeout": 86400
	}
*/
type StartWorkflowRequest struct {
	WorkflowType string      `json:"workflow_type" binding:"required"` // Required. Registered Temporal workflow type
	WorkflowID   string      `json:"workflow_id"`                      // Optional. Client-supplied idempotency key
	ArgsJSON     interface{} `json:"args_json"`                        // Optional. Serialized workflow arguments
	Timeout      int         `json:"timeout"`                          // Optional. Workflow execution timeout in seconds
	StartAt      time.Time   `json:"start_at"`                         // Optional. Scheduled start time (ISO8601)
}

/*
StartWorkflowResponse example:

	{
		"workflow_id": "f3b8f77c-9a2d-4f81-aaa8-1f1d6c39f512",
		"run_id": "3fe4d9a8-0b4b-4520-b8ad-12f67dc862c1",
		"created_at": "2025-11-10T08:30:00Z",
		"start_at": "2025-11-10T08:30:00Z"
	}
*/
type StartWorkflowResponse struct {
	WorkflowID string    `json:"workflow_id"` // Unique workflow identifier
	RunID      string    `json:"run_id"`      // Temporal run identifier
	CreatedAt  time.Time `json:"created_at"`  // Workflow creation time
	StartAt    time.Time `json:"start_at"`    // Scheduled or immediate execution start time
}

/*
SignalWorkflowRequest example:

	{
	  "run_id": "3fe4d9a8-0b4b-4520-b8ad-12f67dc862c1",
	  "signal_name": "CreditWebhook",
	  "signal_payload": {
	    "transaction_id": "tx_456",
	    "status": "SETTLED"
	  }
	}
*/
type SignalWorkflowRequest struct {
	WorkflowID    string      `json:"workflow_id" binding:"required"`    // Required. Running workflow
	RunID         string      `json:"run_id"`                            // Optional. Specific run to signal
	SignalName    string      `json:"signal_name" binding:"required"`    // Required. Registered signal channel
	SignalPayload interface{} `json:"signal_payload" binding:"required"` // Required. Signal data
}

/*
SignalWorkflowResponse example:

	{
		"workflow_id": "f3b8f77c-9a2d-4f81-aaa8-1f1d6c39f512",
		"run_id": "3fe4d9a8-0b4b-4520-b8ad-12f67dc862c1",
		"signal_name": "CreditWebhook",
		"signal_sent_at": "2025-11-10T09:10:00Z"
	}
*/
type SignalWorkflowResponse struct {
	WorkflowID   string    `json:"workflow_id"`    // Signalled workflow
	RunID        string    `json:"run_id"`         // Actual run that received the signal
	SignalName   string    `json:"signal_name"`    // Echoed channel
	SignalSentAt time.Time `json:"signal_sent_at"` // Server timestamp
}

type DescribeWorkflowResponse struct {
	WorkflowID   string     `json:"workflow_id"`
	RunID        string     `json:"run_id"`
	WorkflowType string     `json:"workflow_type"`
	Status       string     `json:"status"` // PENDING/PROCESSING/COMPLETED/FAILED/CANCELLED
	StartedAt    *time.Time `json:"started_at,omitempty"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
	Error        string     `json:"error,omitempty"`
}

type CancelWorkflowRequest struct {
	Reason string `json:"reason,omitempty"`
}

type CancelWorkflowResponse struct {
	WorkflowID  string    `json:"workflow_id"`
	CancelledAt time.Time `json:"cancelled_at"`
}

type ListWorkflowsRequest struct {
	WorkflowType string `form:"workflow_type,omitempty"`
	Status       string `form:"status,omitempty"`
	StartTime    string `form:"start_time,omitempty"`
	EndTime      string `form:"end_time,omitempty"`
	Page         int    `form:"page,omitempty"`
	PageSize     int    `form:"page_size,omitempty"`
}

type WorkflowItem struct {
	WorkflowID   string     `json:"workflow_id"`
	RunID        string     `json:"run_id"`
	WorkflowType string     `json:"workflow_type"`
	Status       string     `json:"status"`
	StartedAt    *time.Time `json:"started_at,omitempty"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
}

type ListWorkflowsResponse struct {
	Workflows      []WorkflowItem `json:"workflows"`
	PaginationMeta PaginationMeta `json:"pagination_meta"`
}

type CreateScheduleRequest struct {
	ScheduleID    string      `json:"schedule_id" binding:"required"`
	WorkflowType  string      `json:"workflow_type" binding:"required"`
	TaskQueue     string      `json:"task_queue" binding:"required"`
	Spec          string      `json:"spec" binding:"required"`  // cron/interval/calendar
	ArgsJSON      interface{} `json:"args_json,omitempty"`      // workflow args
	OverlapPolicy string      `json:"overlap_policy,omitempty"` // Allow / Reject / Queue etc.
}

type CreateScheduleResponse struct {
	ScheduleID string    `json:"schedule_id"`
	NextRunAt  time.Time `json:"next_run_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetScheduleResponse struct {
	ScheduleID    string      `json:"schedule_id"`
	WorkflowType  string      `json:"workflow_type"`
	Spec          string      `json:"spec"`
	State         string      `json:"state"` // ACTIVE/PAUSED/DELETED
	NextRunAt     *time.Time  `json:"next_run_at,omitempty"`
	LastRunResult interface{} `json:"last_run_result,omitempty"`
}

type ListSchedulesRequest struct {
	WorkflowType string `form:"workflow_type,omitempty"`
	State        string `form:"state,omitempty"`
	Page         int    `form:"page,omitempty"`
	PageSize     int    `form:"page_size,omitempty"`
}

type ScheduleItem struct {
	ScheduleID    string      `json:"schedule_id"`
	WorkflowType  string      `json:"workflow_type"`
	Spec          string      `json:"spec"`
	State         string      `json:"state"`
	NextRunAt     *time.Time  `json:"next_run_at,omitempty"`
	LastRunResult interface{} `json:"last_run_result,omitempty"`
}

type ListSchedulesResponse struct {
	Schedules      []ScheduleItem `json:"schedules"`
	PaginationMeta PaginationMeta `json:"pagination_meta"`
}

type DeleteScheduleResponse struct {
	ScheduleID string    `json:"schedule_id"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}
