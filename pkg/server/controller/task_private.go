package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/model"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type TaskControllerForPrivate interface {
	GetTasks(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type taskControllerForPrivate struct {
	TaskRepository repository.TaskRepository
}

func (taskController taskControllerForPrivate) GetTasks(c *gin.Context) {
	var taskRequest request.TaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TaskResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Tasks: []response.Task{}})
		return
	}
	res := taskController.TaskRepository.GetTasks()
	c.JSON(http.StatusOK, res)
	return
}


func (taskController taskControllerForPrivate) CreateTask(c *gin.Context) {
	var taskRequest request.TaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TaskResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Tasks: []response.Task{}})
		return
	}
	var taskModel model.Tasks
	res := taskController.TaskRepository.CreateTask(taskModel)
	c.JSON(http.StatusOK, res)
	return
}


func (taskController taskControllerForPrivate) UpdateTask(c *gin.Context) {
	var taskRequest request.TaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TaskResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Tasks: []response.Task{}})
		return
	}
	var taskModel model.Tasks
	res := taskController.TaskRepository.UpdateTask(taskModel)
	c.JSON(http.StatusOK, res)
	return
}


func (taskController taskControllerForPrivate) DeleteTask(c *gin.Context) {
	var taskRequest request.TaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TaskResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Tasks: []response.Task{}})
		return
	}
	var uuid string
	res := taskController.TaskRepository.DeleteTask(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewTaskControllerForPrivate(taskRepository repository.TaskRepository) TaskControllerForPrivate {
	return &taskControllerForPrivate{ TaskRepository: taskRepository}
}