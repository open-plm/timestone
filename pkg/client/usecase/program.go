package usecase

import (
	"fmt"

	"github.com/open-plm/timestone/pkg/client/repository"
	"github.com/open-plm/timestone/pkg/entity/request"
)

type ProgramUsecase interface {
	BootstrapProgramForDB(request request.ProgramRequest)
	GetProgramForPublic(request request.ProgramRequest)
	GetProgramForInternal(request request.ProgramRequest)
	GetProgramForPrivate(request request.ProgramRequest)
	CreateProgramForPublic(request request.ProgramRequest)
	CreateProgramForInternal(request request.ProgramRequest)
	CreateProgramForPrivate(request request.ProgramRequest)
	UpdateProgramForPublic(request request.ProgramRequest)
	UpdateProgramForInternal(request request.ProgramRequest)
	UpdateProgramForPrivate(request request.ProgramRequest)
	DeleteProgramForPublic(request request.ProgramRequest)
	DeleteProgramForInternal(request request.ProgramRequest)
	DeleteProgramForPrivate(request request.ProgramRequest)
}

type programUsecase struct {
	ProgramRepository   repository.ProgramRepository
}

//Bootstrap
func (programUsecase programUsecase) BootstrapProgramForDB(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.BootstrapProgramForDB(request)
	fmt.Println(programs)
}

//GET
func (programUsecase programUsecase) GetProgramForPublic(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.GetProgramForPublic(request)
	fmt.Println(programs)
}

func (programUsecase programUsecase) GetProgramForInternal(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.GetProgramForInternal(request)
	fmt.Println(programs)
}

func (programUsecase programUsecase) GetProgramForPrivate(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.GetProgramForPrivate(request)
	fmt.Println(programs)
}

//CREATE
func (programUsecase programUsecase) CreateProgramForPublic(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.CreateProgramForPublic(request)
	fmt.Println(programs)
}

func (programUsecase programUsecase) CreateProgramForInternal(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.CreateProgramForInternal(request)
	fmt.Println(programs)
}

func (programUsecase programUsecase) CreateProgramForPrivate(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.CreateProgramForPrivate(request)
	fmt.Println(programs)
}

//UPDATE
func (programUsecase programUsecase) UpdateProgramForPublic(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.UpdateProgramForPublic(request)
	fmt.Println(programs)
}

func (programUsecase programUsecase) UpdateProgramForInternal(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.UpdateProgramForInternal(request)
	fmt.Println(programs)
}

func (programUsecase programUsecase) UpdateProgramForPrivate(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.UpdateProgramForPrivate(request)
	fmt.Println(programs)
}

//DELETE
func (programUsecase programUsecase) DeleteProgramForPublic(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.DeleteProgramForPublic(request)
	fmt.Println(programs)
}

func (programUsecase programUsecase) DeleteProgramForInternal(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.DeleteProgramForInternal(request)
	fmt.Println(programs)
}

func (programUsecase programUsecase) DeleteProgramForPrivate(request request.ProgramRequest) {
	programs := programUsecase.ProgramRepository.DeleteProgramForPrivate(request)
	fmt.Println(programs)
}

func NewProgramUsecase(programRepository repository.ProgramRepository) ProgramUsecase {
	return &programUsecase{ ProgramRepository: programRepository}
}