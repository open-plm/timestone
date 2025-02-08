package usecase

import (
	"fmt"

	"github.com/open-plm/timestone/pkg/client/repository"
	"github.com/open-plm/timestone/pkg/entity/request"
)

type ProjectUsecase interface {
	BootstrapProjectForDB(request request.ProjectRequest)
	GetProjectForPublic(request request.ProjectRequest)
	GetProjectForInternal(request request.ProjectRequest)
	GetProjectForPrivate(request request.ProjectRequest)
	CreateProjectForPublic(request request.ProjectRequest)
	CreateProjectForInternal(request request.ProjectRequest)
	CreateProjectForPrivate(request request.ProjectRequest)
	UpdateProjectForPublic(request request.ProjectRequest)
	UpdateProjectForInternal(request request.ProjectRequest)
	UpdateProjectForPrivate(request request.ProjectRequest)
	DeleteProjectForPublic(request request.ProjectRequest)
	DeleteProjectForInternal(request request.ProjectRequest)
	DeleteProjectForPrivate(request request.ProjectRequest)
}

type projectUsecase struct {
	ProjectRepository   repository.ProjectRepository
}

//Bootstrap
func (projectUsecase projectUsecase) BootstrapProjectForDB(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.BootstrapProjectForDB(request)
	fmt.Println(projects)
}

//GET
func (projectUsecase projectUsecase) GetProjectForPublic(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.GetProjectForPublic(request)
	fmt.Println(projects)
}

func (projectUsecase projectUsecase) GetProjectForInternal(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.GetProjectForInternal(request)
	fmt.Println(projects)
}

func (projectUsecase projectUsecase) GetProjectForPrivate(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.GetProjectForPrivate(request)
	fmt.Println(projects)
}

//CREATE
func (projectUsecase projectUsecase) CreateProjectForPublic(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.CreateProjectForPublic(request)
	fmt.Println(projects)
}

func (projectUsecase projectUsecase) CreateProjectForInternal(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.CreateProjectForInternal(request)
	fmt.Println(projects)
}

func (projectUsecase projectUsecase) CreateProjectForPrivate(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.CreateProjectForPrivate(request)
	fmt.Println(projects)
}

//UPDATE
func (projectUsecase projectUsecase) UpdateProjectForPublic(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.UpdateProjectForPublic(request)
	fmt.Println(projects)
}

func (projectUsecase projectUsecase) UpdateProjectForInternal(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.UpdateProjectForInternal(request)
	fmt.Println(projects)
}

func (projectUsecase projectUsecase) UpdateProjectForPrivate(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.UpdateProjectForPrivate(request)
	fmt.Println(projects)
}

//DELETE
func (projectUsecase projectUsecase) DeleteProjectForPublic(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.DeleteProjectForPublic(request)
	fmt.Println(projects)
}

func (projectUsecase projectUsecase) DeleteProjectForInternal(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.DeleteProjectForInternal(request)
	fmt.Println(projects)
}

func (projectUsecase projectUsecase) DeleteProjectForPrivate(request request.ProjectRequest) {
	projects := projectUsecase.ProjectRepository.DeleteProjectForPrivate(request)
	fmt.Println(projects)
}

func NewProjectUsecase(projectRepository repository.ProjectRepository) ProjectUsecase {
	return &projectUsecase{ ProjectRepository: projectRepository}
}