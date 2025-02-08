package repository

import (
	"time"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/entity/model"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetProjects() []model.Projects
	CreateProject(project model.Projects) *gorm.DB
	UpdateProject(project model.Projects) *gorm.DB
	DeleteProject(uuid string) *gorm.DB
}

type projectRepository struct {
	BaseConfig config.BaseConfig
}


func (projectRepository projectRepository) GetProjects() []model.Projects {
	var projects []model.Projects
	projectRepository.BaseConfig.DBConnection.Find(&projects)
	return projects
}

func (projectRepository projectRepository) CreateProject(project model.Projects) *gorm.DB {
	results := projectRepository.BaseConfig.DBConnection.Create(&project)
	return results
}

func (projectRepository projectRepository) UpdateProject(project model.Projects) *gorm.DB {
	results := projectRepository.BaseConfig.DBConnection.Model(model.Projects{}).Where("uuid = ?", project.UUID).Updates(project)
	return results
}

func (projectRepository projectRepository) DeleteProject(uuid string) *gorm.DB {
	results := projectRepository.BaseConfig.DBConnection.Model(model.Projects{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewProjectRepository(conf config.BaseConfig) ProjectRepository {
	return &projectRepository{BaseConfig: conf}
}