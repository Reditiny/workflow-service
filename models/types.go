package models

import "time"

type StartWorkflowRequest struct {
	WorkflowType string      `json:"workflow_type" binding:"required"`
	WorkflowID   string      `json:"workflow_id,omitempty"`
	TaskQueue    string      `json:"task_queue" binding:"required"`
	ArgsJSON     interface{} `json:"args_json,omitempty"`
	Timeout      string      `json:"timeout,omitempty"` // ISO8601 duration or string like "1h30m"
}

type StartWorkflowResponse struct {
	WorkflowID string    `json:"workflow_id"`
	RunID      string    `json:"run_id"`
	StartedAt  time.Time `json:"started_at"`
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

type SignalWorkflowRequest struct {
	SignalName    string      `json:"signal_name" binding:"required"`
	SignalPayload interface{} `json:"signal_payload,omitempty"`
}

type SignalWorkflowResponse struct {
	WorkflowID   string    `json:"workflow_id"`
	SignalSentAt time.Time `json:"signal_sent_at"`
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
