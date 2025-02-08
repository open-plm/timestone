package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type ProjectControllerForPublic interface {
	GetProjects(c *gin.Context)
}

type projectControllerForPublic struct {
	ProjectRepository repository.ProjectRepository
}

func (projectController projectControllerForPublic) GetProjects(c *gin.Context) {
	var projectRequest request.ProjectRequest
	if err := c.Bind(&projectRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProjectResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Projects: []response.Project{}})
		return
	}
	res := projectController.ProjectRepository.GetProjects()
	c.JSON(http.StatusOK, res)
	return
}


func NewProjectControllerForPublic(projectRepository repository.ProjectRepository) ProjectControllerForPublic {
	return &projectControllerForPublic{ ProjectRepository: projectRepository}
}