package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/model"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type ScheduleControllerForPrivate interface {
	GetSchedules(c *gin.Context)
	CreateSchedule(c *gin.Context)
	UpdateSchedule(c *gin.Context)
	DeleteSchedule(c *gin.Context)
}

type scheduleControllerForPrivate struct {
	ScheduleRepository repository.ScheduleRepository
}

func (scheduleController scheduleControllerForPrivate) GetSchedules(c *gin.Context) {
	var scheduleRequest request.ScheduleRequest
	if err := c.Bind(&scheduleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ScheduleResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Schedules: []response.Schedule{}})
		return
	}
	res := scheduleController.ScheduleRepository.GetSchedules()
	c.JSON(http.StatusOK, res)
	return
}


func (scheduleController scheduleControllerForPrivate) CreateSchedule(c *gin.Context) {
	var scheduleRequest request.ScheduleRequest
	if err := c.Bind(&scheduleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ScheduleResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Schedules: []response.Schedule{}})
		return
	}
	var scheduleModel model.Schedules
	res := scheduleController.ScheduleRepository.CreateSchedule(scheduleModel)
	c.JSON(http.StatusOK, res)
	return
}


func (scheduleController scheduleControllerForPrivate) UpdateSchedule(c *gin.Context) {
	var scheduleRequest request.ScheduleRequest
	if err := c.Bind(&scheduleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ScheduleResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Schedules: []response.Schedule{}})
		return
	}
	var scheduleModel model.Schedules
	res := scheduleController.ScheduleRepository.UpdateSchedule(scheduleModel)
	c.JSON(http.StatusOK, res)
	return
}


func (scheduleController scheduleControllerForPrivate) DeleteSchedule(c *gin.Context) {
	var scheduleRequest request.ScheduleRequest
	if err := c.Bind(&scheduleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ScheduleResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Schedules: []response.Schedule{}})
		return
	}
	var uuid string
	res := scheduleController.ScheduleRepository.DeleteSchedule(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewScheduleControllerForPrivate(scheduleRepository repository.ScheduleRepository) ScheduleControllerForPrivate {
	return &scheduleControllerForPrivate{ ScheduleRepository: scheduleRepository}
}