package repository

import (
	"time"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/entity/model"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	GetSchedules() []model.Schedules
	CreateSchedule(schedule model.Schedules) *gorm.DB
	UpdateSchedule(schedule model.Schedules) *gorm.DB
	DeleteSchedule(uuid string) *gorm.DB
}

type scheduleRepository struct {
	BaseConfig config.BaseConfig
}


func (scheduleRepository scheduleRepository) GetSchedules() []model.Schedules {
	var schedules []model.Schedules
	scheduleRepository.BaseConfig.DBConnection.Find(&schedules)
	return schedules
}

func (scheduleRepository scheduleRepository) CreateSchedule(schedule model.Schedules) *gorm.DB {
	results := scheduleRepository.BaseConfig.DBConnection.Create(&schedule)
	return results
}

func (scheduleRepository scheduleRepository) UpdateSchedule(schedule model.Schedules) *gorm.DB {
	results := scheduleRepository.BaseConfig.DBConnection.Model(model.Schedules{}).Where("uuid = ?", schedule.UUID).Updates(schedule)
	return results
}

func (scheduleRepository scheduleRepository) DeleteSchedule(uuid string) *gorm.DB {
	results := scheduleRepository.BaseConfig.DBConnection.Model(model.Schedules{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewScheduleRepository(conf config.BaseConfig) ScheduleRepository {
	return &scheduleRepository{BaseConfig: conf}
}