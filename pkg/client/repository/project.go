package repository

import (
	"fmt"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
)

type ProjectRepository interface {
	BootstrapProjectForDB(request request.ProjectRequest)(response response.ProjectResponse) 
	GetProjectForPublic(request request.ProjectRequest)(response response.ProjectResponse) 
	GetProjectForInternal(request request.ProjectRequest)(response response.ProjectResponse) 
	GetProjectForPrivate(request request.ProjectRequest)(response response.ProjectResponse) 
	CreateProjectForPublic(request request.ProjectRequest)   (response response.ProjectResponse) 
	CreateProjectForInternal(request request.ProjectRequest) (response response.ProjectResponse)
	CreateProjectForPrivate(request request.ProjectRequest)  (response response.ProjectResponse)
	UpdateProjectForPublic(request request.ProjectRequest)   (response response.ProjectResponse)
	UpdateProjectForInternal(request request.ProjectRequest) (response response.ProjectResponse)
	UpdateProjectForPrivate(request request.ProjectRequest)  (response response.ProjectResponse)
	DeleteProjectForPublic(request request.ProjectRequest)   (response response.ProjectResponse)
	DeleteProjectForInternal(request request.ProjectRequest) (response response.ProjectResponse)
	DeleteProjectForPrivate(request request.ProjectRequest)  (response response.ProjectResponse)
}

type projectRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (projectRepository projectRepository) BootstrapProjectForDB(request request.ProjectRequest) (response response.ProjectResponse) {
	fmt.Println("BootstrapProjectForDB")
	return response
}

//GET
func (projectRepository projectRepository) GetProjectForPublic(request request.ProjectRequest) (response response.ProjectResponse) {
	fmt.Println("GetProjectForPublic")
	return response
}

func (projectRepository projectRepository) GetProjectForInternal(request request.ProjectRequest) (response response.ProjectResponse ){
	fmt.Println("GetProjectForInternal")
	return response
}

func (projectRepository projectRepository) GetProjectForPrivate(request request.ProjectRequest) (response response.ProjectResponse) {
	fmt.Println("GetProjectForPrivate")
	return response
}

//CREATE
func (projectRepository projectRepository) CreateProjectForPublic(request request.ProjectRequest) (response response.ProjectResponse ){
	fmt.Println("CreateProjectForPublic")
	return response
}

func (projectRepository projectRepository) CreateProjectForInternal(request request.ProjectRequest) (response response.ProjectResponse) {
	fmt.Println("CreateProjectForInternal()")
	return response
}

func (projectRepository projectRepository) CreateProjectForPrivate(request request.ProjectRequest) (response response.ProjectResponse){
	fmt.Println("CreateProjectForPrivate()")
	return response
}

//UPDATE
func (projectRepository projectRepository) UpdateProjectForPublic(request request.ProjectRequest) (response response.ProjectResponse){
	fmt.Println("UpdateProjectForPublic()")
	return response
}

func (projectRepository projectRepository) UpdateProjectForInternal(request request.ProjectRequest) (response response.ProjectResponse){
	fmt.Println("UpdateProjectForInternal")
	return response
}

func (projectRepository projectRepository) UpdateProjectForPrivate(request request.ProjectRequest) (response response.ProjectResponse){
	fmt.Println("UpdateProjectForPrivate")
	return response
}

//DELETE
func (projectRepository projectRepository) DeleteProjectForPublic(request request.ProjectRequest) (response response.ProjectResponse){
	fmt.Println("DeleteProjectForPublic")
	return response
}

func (projectRepository projectRepository) DeleteProjectForInternal(request request.ProjectRequest) (response response.ProjectResponse ){
	fmt.Println("DeleteProjectForInternal")
	return response
}

func (projectRepository projectRepository) DeleteProjectForPrivate(request request.ProjectRequest) (response response.ProjectResponse){
	fmt.Println("DeleteProjectForPrivate")
	return response
}

func NewProjectRepository(conf config.BaseConfig) ProjectRepository {
	return &projectRepository{BaseConfig: conf}
}