package repository

import (
	"fmt"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
)

type ScheduleRepository interface {
	BootstrapScheduleForDB(request request.ScheduleRequest)(response response.ScheduleResponse) 
	GetScheduleForPublic(request request.ScheduleRequest)(response response.ScheduleResponse) 
	GetScheduleForInternal(request request.ScheduleRequest)(response response.ScheduleResponse) 
	GetScheduleForPrivate(request request.ScheduleRequest)(response response.ScheduleResponse) 
	CreateScheduleForPublic(request request.ScheduleRequest)   (response response.ScheduleResponse) 
	CreateScheduleForInternal(request request.ScheduleRequest) (response response.ScheduleResponse)
	CreateScheduleForPrivate(request request.ScheduleRequest)  (response response.ScheduleResponse)
	UpdateScheduleForPublic(request request.ScheduleRequest)   (response response.ScheduleResponse)
	UpdateScheduleForInternal(request request.ScheduleRequest) (response response.ScheduleResponse)
	UpdateScheduleForPrivate(request request.ScheduleRequest)  (response response.ScheduleResponse)
	DeleteScheduleForPublic(request request.ScheduleRequest)   (response response.ScheduleResponse)
	DeleteScheduleForInternal(request request.ScheduleRequest) (response response.ScheduleResponse)
	DeleteScheduleForPrivate(request request.ScheduleRequest)  (response response.ScheduleResponse)
}

type scheduleRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (scheduleRepository scheduleRepository) BootstrapScheduleForDB(request request.ScheduleRequest) (response response.ScheduleResponse) {
	fmt.Println("BootstrapScheduleForDB")
	return response
}

//GET
func (scheduleRepository scheduleRepository) GetScheduleForPublic(request request.ScheduleRequest) (response response.ScheduleResponse) {
	fmt.Println("GetScheduleForPublic")
	return response
}

func (scheduleRepository scheduleRepository) GetScheduleForInternal(request request.ScheduleRequest) (response response.ScheduleResponse ){
	fmt.Println("GetScheduleForInternal")
	return response
}

func (scheduleRepository scheduleRepository) GetScheduleForPrivate(request request.ScheduleRequest) (response response.ScheduleResponse) {
	fmt.Println("GetScheduleForPrivate")
	return response
}

//CREATE
func (scheduleRepository scheduleRepository) CreateScheduleForPublic(request request.ScheduleRequest) (response response.ScheduleResponse ){
	fmt.Println("CreateScheduleForPublic")
	return response
}

func (scheduleRepository scheduleRepository) CreateScheduleForInternal(request request.ScheduleRequest) (response response.ScheduleResponse) {
	fmt.Println("CreateScheduleForInternal()")
	return response
}

func (scheduleRepository scheduleRepository) CreateScheduleForPrivate(request request.ScheduleRequest) (response response.ScheduleResponse){
	fmt.Println("CreateScheduleForPrivate()")
	return response
}

//UPDATE
func (scheduleRepository scheduleRepository) UpdateScheduleForPublic(request request.ScheduleRequest) (response response.ScheduleResponse){
	fmt.Println("UpdateScheduleForPublic()")
	return response
}

func (scheduleRepository scheduleRepository) UpdateScheduleForInternal(request request.ScheduleRequest) (response response.ScheduleResponse){
	fmt.Println("UpdateScheduleForInternal")
	return response
}

func (scheduleRepository scheduleRepository) UpdateScheduleForPrivate(request request.ScheduleRequest) (response response.ScheduleResponse){
	fmt.Println("UpdateScheduleForPrivate")
	return response
}

//DELETE
func (scheduleRepository scheduleRepository) DeleteScheduleForPublic(request request.ScheduleRequest) (response response.ScheduleResponse){
	fmt.Println("DeleteScheduleForPublic")
	return response
}

func (scheduleRepository scheduleRepository) DeleteScheduleForInternal(request request.ScheduleRequest) (response response.ScheduleResponse ){
	fmt.Println("DeleteScheduleForInternal")
	return response
}

func (scheduleRepository scheduleRepository) DeleteScheduleForPrivate(request request.ScheduleRequest) (response response.ScheduleResponse){
	fmt.Println("DeleteScheduleForPrivate")
	return response
}

func NewScheduleRepository(conf config.BaseConfig) ScheduleRepository {
	return &scheduleRepository{BaseConfig: conf}
}