package usecase

import (
	"fmt"

	"github.com/open-plm/timestone/pkg/client/repository"
	"github.com/open-plm/timestone/pkg/entity/request"
)

type TaskUsecase interface {
	BootstrapTaskForDB(request request.TaskRequest)
	GetTaskForPublic(request request.TaskRequest)
	GetTaskForInternal(request request.TaskRequest)
	GetTaskForPrivate(request request.TaskRequest)
	CreateTaskForPublic(request request.TaskRequest)
	CreateTaskForInternal(request request.TaskRequest)
	CreateTaskForPrivate(request request.TaskRequest)
	UpdateTaskForPublic(request request.TaskRequest)
	UpdateTaskForInternal(request request.TaskRequest)
	UpdateTaskForPrivate(request request.TaskRequest)
	DeleteTaskForPublic(request request.TaskRequest)
	DeleteTaskForInternal(request request.TaskRequest)
	DeleteTaskForPrivate(request request.TaskRequest)
}

type taskUsecase struct {
	TaskRepository   repository.TaskRepository
}

//Bootstrap
func (taskUsecase taskUsecase) BootstrapTaskForDB(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.BootstrapTaskForDB(request)
	fmt.Println(tasks)
}

//GET
func (taskUsecase taskUsecase) GetTaskForPublic(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.GetTaskForPublic(request)
	fmt.Println(tasks)
}

func (taskUsecase taskUsecase) GetTaskForInternal(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.GetTaskForInternal(request)
	fmt.Println(tasks)
}

func (taskUsecase taskUsecase) GetTaskForPrivate(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.GetTaskForPrivate(request)
	fmt.Println(tasks)
}

//CREATE
func (taskUsecase taskUsecase) CreateTaskForPublic(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.CreateTaskForPublic(request)
	fmt.Println(tasks)
}

func (taskUsecase taskUsecase) CreateTaskForInternal(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.CreateTaskForInternal(request)
	fmt.Println(tasks)
}

func (taskUsecase taskUsecase) CreateTaskForPrivate(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.CreateTaskForPrivate(request)
	fmt.Println(tasks)
}

//UPDATE
func (taskUsecase taskUsecase) UpdateTaskForPublic(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.UpdateTaskForPublic(request)
	fmt.Println(tasks)
}

func (taskUsecase taskUsecase) UpdateTaskForInternal(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.UpdateTaskForInternal(request)
	fmt.Println(tasks)
}

func (taskUsecase taskUsecase) UpdateTaskForPrivate(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.UpdateTaskForPrivate(request)
	fmt.Println(tasks)
}

//DELETE
func (taskUsecase taskUsecase) DeleteTaskForPublic(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.DeleteTaskForPublic(request)
	fmt.Println(tasks)
}

func (taskUsecase taskUsecase) DeleteTaskForInternal(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.DeleteTaskForInternal(request)
	fmt.Println(tasks)
}

func (taskUsecase taskUsecase) DeleteTaskForPrivate(request request.TaskRequest) {
	tasks := taskUsecase.TaskRepository.DeleteTaskForPrivate(request)
	fmt.Println(tasks)
}

func NewTaskUsecase(taskRepository repository.TaskRepository) TaskUsecase {
	return &taskUsecase{ TaskRepository: taskRepository}
}