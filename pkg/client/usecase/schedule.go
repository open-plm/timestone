package usecase

import (
	"fmt"

	"github.com/open-plm/timestone/pkg/client/repository"
	"github.com/open-plm/timestone/pkg/entity/request"
)

type ScheduleUsecase interface {
	BootstrapScheduleForDB(request request.ScheduleRequest)
	GetScheduleForPublic(request request.ScheduleRequest)
	GetScheduleForInternal(request request.ScheduleRequest)
	GetScheduleForPrivate(request request.ScheduleRequest)
	CreateScheduleForPublic(request request.ScheduleRequest)
	CreateScheduleForInternal(request request.ScheduleRequest)
	CreateScheduleForPrivate(request request.ScheduleRequest)
	UpdateScheduleForPublic(request request.ScheduleRequest)
	UpdateScheduleForInternal(request request.ScheduleRequest)
	UpdateScheduleForPrivate(request request.ScheduleRequest)
	DeleteScheduleForPublic(request request.ScheduleRequest)
	DeleteScheduleForInternal(request request.ScheduleRequest)
	DeleteScheduleForPrivate(request request.ScheduleRequest)
}

type scheduleUsecase struct {
	ScheduleRepository   repository.ScheduleRepository
}

//Bootstrap
func (scheduleUsecase scheduleUsecase) BootstrapScheduleForDB(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.BootstrapScheduleForDB(request)
	fmt.Println(schedules)
}

//GET
func (scheduleUsecase scheduleUsecase) GetScheduleForPublic(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.GetScheduleForPublic(request)
	fmt.Println(schedules)
}

func (scheduleUsecase scheduleUsecase) GetScheduleForInternal(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.GetScheduleForInternal(request)
	fmt.Println(schedules)
}

func (scheduleUsecase scheduleUsecase) GetScheduleForPrivate(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.GetScheduleForPrivate(request)
	fmt.Println(schedules)
}

//CREATE
func (scheduleUsecase scheduleUsecase) CreateScheduleForPublic(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.CreateScheduleForPublic(request)
	fmt.Println(schedules)
}

func (scheduleUsecase scheduleUsecase) CreateScheduleForInternal(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.CreateScheduleForInternal(request)
	fmt.Println(schedules)
}

func (scheduleUsecase scheduleUsecase) CreateScheduleForPrivate(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.CreateScheduleForPrivate(request)
	fmt.Println(schedules)
}

//UPDATE
func (scheduleUsecase scheduleUsecase) UpdateScheduleForPublic(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.UpdateScheduleForPublic(request)
	fmt.Println(schedules)
}

func (scheduleUsecase scheduleUsecase) UpdateScheduleForInternal(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.UpdateScheduleForInternal(request)
	fmt.Println(schedules)
}

func (scheduleUsecase scheduleUsecase) UpdateScheduleForPrivate(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.UpdateScheduleForPrivate(request)
	fmt.Println(schedules)
}

//DELETE
func (scheduleUsecase scheduleUsecase) DeleteScheduleForPublic(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.DeleteScheduleForPublic(request)
	fmt.Println(schedules)
}

func (scheduleUsecase scheduleUsecase) DeleteScheduleForInternal(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.DeleteScheduleForInternal(request)
	fmt.Println(schedules)
}

func (scheduleUsecase scheduleUsecase) DeleteScheduleForPrivate(request request.ScheduleRequest) {
	schedules := scheduleUsecase.ScheduleRepository.DeleteScheduleForPrivate(request)
	fmt.Println(schedules)
}

func NewScheduleUsecase(scheduleRepository repository.ScheduleRepository) ScheduleUsecase {
	return &scheduleUsecase{ ScheduleRepository: scheduleRepository}
}