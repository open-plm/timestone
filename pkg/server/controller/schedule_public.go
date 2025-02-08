package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type ScheduleControllerForPublic interface {
	GetSchedules(c *gin.Context)
}

type scheduleControllerForPublic struct {
	ScheduleRepository repository.ScheduleRepository
}

func (scheduleController scheduleControllerForPublic) GetSchedules(c *gin.Context) {
	var scheduleRequest request.ScheduleRequest
	if err := c.Bind(&scheduleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ScheduleResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Schedules: []response.Schedule{}})
		return
	}
	res := scheduleController.ScheduleRepository.GetSchedules()
	c.JSON(http.StatusOK, res)
	return
}


func NewScheduleControllerForPublic(scheduleRepository repository.ScheduleRepository) ScheduleControllerForPublic {
	return &scheduleControllerForPublic{ ScheduleRepository: scheduleRepository}
}