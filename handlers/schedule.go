package handlers

import (
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
	"net/http"
)

func RegisterScheduleRoutes(rg *gin.RouterGroup, c client.Client) {
	h := &ScheduleHandler{}

	s := rg.Group("/workflows/schedules")
	{
		s.POST("", h.CreateSchedule)
		s.GET("/:schedule_id", h.GetSchedule)
		s.GET("", h.ListSchedules)
		s.POST("/:schedule_id/pause", h.PauseSchedule)
		s.POST("/:schedule_id/resume", h.ResumeSchedule)
		s.DELETE("/:schedule_id", h.DeleteSchedule)
		s.POST("/:schedule_id/trigger", h.TriggerScheduleNow)
	}
}

type ScheduleHandler struct {
}

func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateSchedule placeholder"})
}

func (h *ScheduleHandler) ListSchedules(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ListSchedules placeholder"})
}

func (h *ScheduleHandler) GetSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetSchedule placeholder"})
}

func (h *ScheduleHandler) PauseSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetSchedule placeholder"})
}

func (h *ScheduleHandler) ResumeSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetSchedule placeholder"})
}

func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteSchedule placeholder"})
}

func (h *ScheduleHandler) TriggerScheduleNow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteSchedule placeholder"})
}
