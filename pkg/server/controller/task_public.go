package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type TaskControllerForPublic interface {
	GetTasks(c *gin.Context)
}

type taskControllerForPublic struct {
	TaskRepository repository.TaskRepository
}

func (taskController taskControllerForPublic) GetTasks(c *gin.Context) {
	var taskRequest request.TaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TaskResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Tasks: []response.Task{}})
		return
	}
	res := taskController.TaskRepository.GetTasks()
	c.JSON(http.StatusOK, res)
	return
}


func NewTaskControllerForPublic(taskRepository repository.TaskRepository) TaskControllerForPublic {
	return &taskControllerForPublic{ TaskRepository: taskRepository}
}