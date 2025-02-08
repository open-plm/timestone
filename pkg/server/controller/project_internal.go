package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/model"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type ProjectControllerForInternal interface {
	GetProjects(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
}

type projectControllerForInternal struct {
	ProjectRepository repository.ProjectRepository
}

func (projectController projectControllerForInternal) GetProjects(c *gin.Context) {
	var projectRequest request.ProjectRequest
	if err := c.Bind(&projectRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProjectResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Projects: []response.Project{}})
		return
	}
	res := projectController.ProjectRepository.GetProjects()
	c.JSON(http.StatusOK, res)
	return
}


func (projectController projectControllerForInternal) CreateProject(c *gin.Context) {
	var projectRequest request.ProjectRequest
	if err := c.Bind(&projectRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProjectResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Projects: []response.Project{}})
		return
	}
	var projectModel model.Projects
	res := projectController.ProjectRepository.CreateProject(projectModel)
	c.JSON(http.StatusOK, res)
	return
}


func (projectController projectControllerForInternal) UpdateProject(c *gin.Context) {
	var projectRequest request.ProjectRequest
	if err := c.Bind(&projectRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProjectResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Projects: []response.Project{}})
		return
	}
	var projectModel model.Projects
	res := projectController.ProjectRepository.UpdateProject(projectModel)
	c.JSON(http.StatusOK, res)
	return
}


func (projectController projectControllerForInternal) DeleteProject(c *gin.Context) {
	var projectRequest request.ProjectRequest
	if err := c.Bind(&projectRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProjectResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Projects: []response.Project{}})
		return
	}
	var uuid string
	res := projectController.ProjectRepository.DeleteProject(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewProjectControllerForInternal(projectRepository repository.ProjectRepository) ProjectControllerForInternal {
	return &projectControllerForInternal{ ProjectRepository: projectRepository}
}