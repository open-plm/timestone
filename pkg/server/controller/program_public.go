package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
	"github.com/open-plm/timestone/pkg/server/repository"
)

type ProgramControllerForPublic interface {
	GetPrograms(c *gin.Context)
}

type programControllerForPublic struct {
	ProgramRepository repository.ProgramRepository
}

func (programController programControllerForPublic) GetPrograms(c *gin.Context) {
	var programRequest request.ProgramRequest
	if err := c.Bind(&programRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProgramResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Programs: []response.Program{}})
		return
	}
	res := programController.ProgramRepository.GetPrograms()
	c.JSON(http.StatusOK, res)
	return
}


func NewProgramControllerForPublic(programRepository repository.ProgramRepository) ProgramControllerForPublic {
	return &programControllerForPublic{ ProgramRepository: programRepository}
}