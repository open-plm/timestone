package repository

import (
	"time"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/entity/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks() []model.Tasks
	CreateTask(task model.Tasks) *gorm.DB
	UpdateTask(task model.Tasks) *gorm.DB
	DeleteTask(uuid string) *gorm.DB
}

type taskRepository struct {
	BaseConfig config.BaseConfig
}


func (taskRepository taskRepository) GetTasks() []model.Tasks {
	var tasks []model.Tasks
	taskRepository.BaseConfig.DBConnection.Find(&tasks)
	return tasks
}

func (taskRepository taskRepository) CreateTask(task model.Tasks) *gorm.DB {
	results := taskRepository.BaseConfig.DBConnection.Create(&task)
	return results
}

func (taskRepository taskRepository) UpdateTask(task model.Tasks) *gorm.DB {
	results := taskRepository.BaseConfig.DBConnection.Model(model.Tasks{}).Where("uuid = ?", task.UUID).Updates(task)
	return results
}

func (taskRepository taskRepository) DeleteTask(uuid string) *gorm.DB {
	results := taskRepository.BaseConfig.DBConnection.Model(model.Tasks{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewTaskRepository(conf config.BaseConfig) TaskRepository {
	return &taskRepository{BaseConfig: conf}
}