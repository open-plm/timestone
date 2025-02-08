package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/model"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type ProgramControllerForInternal interface {
	GetPrograms(c *gin.Context)
	CreateProgram(c *gin.Context)
	UpdateProgram(c *gin.Context)
	DeleteProgram(c *gin.Context)
}

type programControllerForInternal struct {
	ProgramRepository repository.ProgramRepository
}

func (programController programControllerForInternal) GetPrograms(c *gin.Context) {
	var programRequest request.ProgramRequest
	if err := c.Bind(&programRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProgramResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Programs: []response.Program{}})
		return
	}
	res := programController.ProgramRepository.GetPrograms()
	c.JSON(http.StatusOK, res)
	return
}


func (programController programControllerForInternal) CreateProgram(c *gin.Context) {
	var programRequest request.ProgramRequest
	if err := c.Bind(&programRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProgramResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Programs: []response.Program{}})
		return
	}
	var programModel model.Programs
	res := programController.ProgramRepository.CreateProgram(programModel)
	c.JSON(http.StatusOK, res)
	return
}


func (programController programControllerForInternal) UpdateProgram(c *gin.Context) {
	var programRequest request.ProgramRequest
	if err := c.Bind(&programRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProgramResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Programs: []response.Program{}})
		return
	}
	var programModel model.Programs
	res := programController.ProgramRepository.UpdateProgram(programModel)
	c.JSON(http.StatusOK, res)
	return
}


func (programController programControllerForInternal) DeleteProgram(c *gin.Context) {
	var programRequest request.ProgramRequest
	if err := c.Bind(&programRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProgramResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Programs: []response.Program{}})
		return
	}
	var uuid string
	res := programController.ProgramRepository.DeleteProgram(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewProgramControllerForInternal(programRepository repository.ProgramRepository) ProgramControllerForInternal {
	return &programControllerForInternal{ ProgramRepository: programRepository}
}