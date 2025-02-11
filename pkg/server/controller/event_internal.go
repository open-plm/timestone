package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/model"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type EventControllerForInternal interface {
	GetEvents(c *gin.Context)
	CreateEvent(c *gin.Context)
	UpdateEvent(c *gin.Context)
	DeleteEvent(c *gin.Context)
}

type eventControllerForInternal struct {
	EventRepository repository.EventRepository
}

func (eventController eventControllerForInternal) GetEvents(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	res := eventController.EventRepository.GetEvents()
	c.JSON(http.StatusOK, res)
	return
}


func (eventController eventControllerForInternal) CreateEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var eventModel model.Events
	res := eventController.EventRepository.CreateEvent(eventModel)
	c.JSON(http.StatusOK, res)
	return
}


func (eventController eventControllerForInternal) UpdateEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var eventModel model.Events
	res := eventController.EventRepository.UpdateEvent(eventModel)
	c.JSON(http.StatusOK, res)
	return
}


func (eventController eventControllerForInternal) DeleteEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var uuid string
	res := eventController.EventRepository.DeleteEvent(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewEventControllerForInternal(eventRepository repository.EventRepository) EventControllerForInternal {
	return &eventControllerForInternal{ EventRepository: eventRepository}
}