package repository

import (
	"fmt"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
)

type ProgramRepository interface {
	BootstrapProgramForDB(request request.ProgramRequest)(response response.ProgramResponse) 
	GetProgramForPublic(request request.ProgramRequest)(response response.ProgramResponse) 
	GetProgramForInternal(request request.ProgramRequest)(response response.ProgramResponse) 
	GetProgramForPrivate(request request.ProgramRequest)(response response.ProgramResponse) 
	CreateProgramForPublic(request request.ProgramRequest)   (response response.ProgramResponse) 
	CreateProgramForInternal(request request.ProgramRequest) (response response.ProgramResponse)
	CreateProgramForPrivate(request request.ProgramRequest)  (response response.ProgramResponse)
	UpdateProgramForPublic(request request.ProgramRequest)   (response response.ProgramResponse)
	UpdateProgramForInternal(request request.ProgramRequest) (response response.ProgramResponse)
	UpdateProgramForPrivate(request request.ProgramRequest)  (response response.ProgramResponse)
	DeleteProgramForPublic(request request.ProgramRequest)   (response response.ProgramResponse)
	DeleteProgramForInternal(request request.ProgramRequest) (response response.ProgramResponse)
	DeleteProgramForPrivate(request request.ProgramRequest)  (response response.ProgramResponse)
}

type programRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (programRepository programRepository) BootstrapProgramForDB(request request.ProgramRequest) (response response.ProgramResponse) {
	fmt.Println("BootstrapProgramForDB")
	return response
}

//GET
func (programRepository programRepository) GetProgramForPublic(request request.ProgramRequest) (response response.ProgramResponse) {
	fmt.Println("GetProgramForPublic")
	return response
}

func (programRepository programRepository) GetProgramForInternal(request request.ProgramRequest) (response response.ProgramResponse ){
	fmt.Println("GetProgramForInternal")
	return response
}

func (programRepository programRepository) GetProgramForPrivate(request request.ProgramRequest) (response response.ProgramResponse) {
	fmt.Println("GetProgramForPrivate")
	return response
}

//CREATE
func (programRepository programRepository) CreateProgramForPublic(request request.ProgramRequest) (response response.ProgramResponse ){
	fmt.Println("CreateProgramForPublic")
	return response
}

func (programRepository programRepository) CreateProgramForInternal(request request.ProgramRequest) (response response.ProgramResponse) {
	fmt.Println("CreateProgramForInternal()")
	return response
}

func (programRepository programRepository) CreateProgramForPrivate(request request.ProgramRequest) (response response.ProgramResponse){
	fmt.Println("CreateProgramForPrivate()")
	return response
}

//UPDATE
func (programRepository programRepository) UpdateProgramForPublic(request request.ProgramRequest) (response response.ProgramResponse){
	fmt.Println("UpdateProgramForPublic()")
	return response
}

func (programRepository programRepository) UpdateProgramForInternal(request request.ProgramRequest) (response response.ProgramResponse){
	fmt.Println("UpdateProgramForInternal")
	return response
}

func (programRepository programRepository) UpdateProgramForPrivate(request request.ProgramRequest) (response response.ProgramResponse){
	fmt.Println("UpdateProgramForPrivate")
	return response
}

//DELETE
func (programRepository programRepository) DeleteProgramForPublic(request request.ProgramRequest) (response response.ProgramResponse){
	fmt.Println("DeleteProgramForPublic")
	return response
}

func (programRepository programRepository) DeleteProgramForInternal(request request.ProgramRequest) (response response.ProgramResponse ){
	fmt.Println("DeleteProgramForInternal")
	return response
}

func (programRepository programRepository) DeleteProgramForPrivate(request request.ProgramRequest) (response response.ProgramResponse){
	fmt.Println("DeleteProgramForPrivate")
	return response
}

func NewProgramRepository(conf config.BaseConfig) ProgramRepository {
	return &programRepository{BaseConfig: conf}
}