package repository

import (
	"fmt"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/entity/request"
	"github.com/open-plm/timestone/pkg/entity/response"
)

type TaskRepository interface {
	BootstrapTaskForDB(request request.TaskRequest)(response response.TaskResponse) 
	GetTaskForPublic(request request.TaskRequest)(response response.TaskResponse) 
	GetTaskForInternal(request request.TaskRequest)(response response.TaskResponse) 
	GetTaskForPrivate(request request.TaskRequest)(response response.TaskResponse) 
	CreateTaskForPublic(request request.TaskRequest)   (response response.TaskResponse) 
	CreateTaskForInternal(request request.TaskRequest) (response response.TaskResponse)
	CreateTaskForPrivate(request request.TaskRequest)  (response response.TaskResponse)
	UpdateTaskForPublic(request request.TaskRequest)   (response response.TaskResponse)
	UpdateTaskForInternal(request request.TaskRequest) (response response.TaskResponse)
	UpdateTaskForPrivate(request request.TaskRequest)  (response response.TaskResponse)
	DeleteTaskForPublic(request request.TaskRequest)   (response response.TaskResponse)
	DeleteTaskForInternal(request request.TaskRequest) (response response.TaskResponse)
	DeleteTaskForPrivate(request request.TaskRequest)  (response response.TaskResponse)
}

type taskRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (taskRepository taskRepository) BootstrapTaskForDB(request request.TaskRequest) (response response.TaskResponse) {
	fmt.Println("BootstrapTaskForDB")
	return response
}

//GET
func (taskRepository taskRepository) GetTaskForPublic(request request.TaskRequest) (response response.TaskResponse) {
	fmt.Println("GetTaskForPublic")
	return response
}

func (taskRepository taskRepository) GetTaskForInternal(request request.TaskRequest) (response response.TaskResponse ){
	fmt.Println("GetTaskForInternal")
	return response
}

func (taskRepository taskRepository) GetTaskForPrivate(request request.TaskRequest) (response response.TaskResponse) {
	fmt.Println("GetTaskForPrivate")
	return response
}

//CREATE
func (taskRepository taskRepository) CreateTaskForPublic(request request.TaskRequest) (response response.TaskResponse ){
	fmt.Println("CreateTaskForPublic")
	return response
}

func (taskRepository taskRepository) CreateTaskForInternal(request request.TaskRequest) (response response.TaskResponse) {
	fmt.Println("CreateTaskForInternal()")
	return response
}

func (taskRepository taskRepository) CreateTaskForPrivate(request request.TaskRequest) (response response.TaskResponse){
	fmt.Println("CreateTaskForPrivate()")
	return response
}

//UPDATE
func (taskRepository taskRepository) UpdateTaskForPublic(request request.TaskRequest) (response response.TaskResponse){
	fmt.Println("UpdateTaskForPublic()")
	return response
}

func (taskRepository taskRepository) UpdateTaskForInternal(request request.TaskRequest) (response response.TaskResponse){
	fmt.Println("UpdateTaskForInternal")
	return response
}

func (taskRepository taskRepository) UpdateTaskForPrivate(request request.TaskRequest) (response response.TaskResponse){
	fmt.Println("UpdateTaskForPrivate")
	return response
}

//DELETE
func (taskRepository taskRepository) DeleteTaskForPublic(request request.TaskRequest) (response response.TaskResponse){
	fmt.Println("DeleteTaskForPublic")
	return response
}

func (taskRepository taskRepository) DeleteTaskForInternal(request request.TaskRequest) (response response.TaskResponse ){
	fmt.Println("DeleteTaskForInternal")
	return response
}

func (taskRepository taskRepository) DeleteTaskForPrivate(request request.TaskRequest) (response response.TaskResponse){
	fmt.Println("DeleteTaskForPrivate")
	return response
}

func NewTaskRepository(conf config.BaseConfig) TaskRepository {
	return &taskRepository{BaseConfig: conf}
}