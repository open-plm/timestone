package repository

import (
	"time"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/entity/model"
	"gorm.io/gorm"
)

type ProgramRepository interface {
	GetPrograms() []model.Programs
	CreateProgram(program model.Programs) *gorm.DB
	UpdateProgram(program model.Programs) *gorm.DB
	DeleteProgram(uuid string) *gorm.DB
}

type programRepository struct {
	BaseConfig config.BaseConfig
}


func (programRepository programRepository) GetPrograms() []model.Programs {
	var programs []model.Programs
	programRepository.BaseConfig.DBConnection.Find(&programs)
	return programs
}

func (programRepository programRepository) CreateProgram(program model.Programs) *gorm.DB {
	results := programRepository.BaseConfig.DBConnection.Create(&program)
	return results
}

func (programRepository programRepository) UpdateProgram(program model.Programs) *gorm.DB {
	results := programRepository.BaseConfig.DBConnection.Model(model.Programs{}).Where("uuid = ?", program.UUID).Updates(program)
	return results
}

func (programRepository programRepository) DeleteProgram(uuid string) *gorm.DB {
	results := programRepository.BaseConfig.DBConnection.Model(model.Programs{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewProgramRepository(conf config.BaseConfig) ProgramRepository {
	return &programRepository{BaseConfig: conf}
}