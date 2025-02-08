package server

import (
	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/server/controller"
	"github.com/open-plm/timestone/pkg/server/middleware"
	"github.com/open-plm/timestone/pkg/server/repository"
)

func InitRouter(conf config.BaseConfig) *gin.Engine {
	
	programRepository := repository.NewProgramRepository(conf)
	programControllerForPublic := controller.NewProgramControllerForPublic(programRepository)
	programControllerForInternal := controller.NewProgramControllerForInternal(programRepository)
	programControllerForPrivate := controller.NewProgramControllerForPrivate(programRepository)
	
	projectRepository := repository.NewProjectRepository(conf)
	projectControllerForPublic := controller.NewProjectControllerForPublic(projectRepository)
	projectControllerForInternal := controller.NewProjectControllerForInternal(projectRepository)
	projectControllerForPrivate := controller.NewProjectControllerForPrivate(projectRepository)
	
	scheduleRepository := repository.NewScheduleRepository(conf)
	scheduleControllerForPublic := controller.NewScheduleControllerForPublic(scheduleRepository)
	scheduleControllerForInternal := controller.NewScheduleControllerForInternal(scheduleRepository)
	scheduleControllerForPrivate := controller.NewScheduleControllerForPrivate(scheduleRepository)
	
	taskRepository := repository.NewTaskRepository(conf)
	taskControllerForPublic := controller.NewTaskControllerForPublic(taskRepository)
	taskControllerForInternal := controller.NewTaskControllerForInternal(taskRepository)
	taskControllerForPrivate := controller.NewTaskControllerForPrivate(taskRepository)
	
	eventRepository := repository.NewEventRepository(conf)
	eventControllerForPublic := controller.NewEventControllerForPublic(eventRepository)
	eventControllerForInternal := controller.NewEventControllerForInternal(eventRepository)
	eventControllerForPrivate := controller.NewEventControllerForPrivate(eventRepository)
	
	
    router := gin.Default()
	privateAPI := router.Group("api/private")
	internalAPI := router.Group("api/internal")
	publicAPI := router.Group("api/public")

	publicAPI.Use(middleware.ForPublic(conf))
	internalAPI.Use(middleware.ForInternal(conf))
	privateAPI.Use(middleware.ForPrivate(conf))

	
	//program
	publicAPI.GET("/programs", programControllerForPublic.GetPrograms)
	internalAPI.GET("/programs", programControllerForInternal.GetPrograms)
	internalAPI.POST("/program", programControllerForInternal.CreateProgram)
	internalAPI.PUT("/program", programControllerForInternal.UpdateProgram)
	internalAPI.DELETE("/program", programControllerForInternal.DeleteProgram)
	privateAPI.GET("/programs", programControllerForPrivate.GetPrograms)
	privateAPI.POST("/program", programControllerForPrivate.CreateProgram)
	privateAPI.PUT("/program", programControllerForPrivate.UpdateProgram)
	privateAPI.DELETE("/program", programControllerForPrivate.DeleteProgram)
	
	//project
	publicAPI.GET("/projects", projectControllerForPublic.GetProjects)
	internalAPI.GET("/projects", projectControllerForInternal.GetProjects)
	internalAPI.POST("/project", projectControllerForInternal.CreateProject)
	internalAPI.PUT("/project", projectControllerForInternal.UpdateProject)
	internalAPI.DELETE("/project", projectControllerForInternal.DeleteProject)
	privateAPI.GET("/projects", projectControllerForPrivate.GetProjects)
	privateAPI.POST("/project", projectControllerForPrivate.CreateProject)
	privateAPI.PUT("/project", projectControllerForPrivate.UpdateProject)
	privateAPI.DELETE("/project", projectControllerForPrivate.DeleteProject)
	
	//schedule
	publicAPI.GET("/schedules", scheduleControllerForPublic.GetSchedules)
	internalAPI.GET("/schedules", scheduleControllerForInternal.GetSchedules)
	internalAPI.POST("/schedule", scheduleControllerForInternal.CreateSchedule)
	internalAPI.PUT("/schedule", scheduleControllerForInternal.UpdateSchedule)
	internalAPI.DELETE("/schedule", scheduleControllerForInternal.DeleteSchedule)
	privateAPI.GET("/schedules", scheduleControllerForPrivate.GetSchedules)
	privateAPI.POST("/schedule", scheduleControllerForPrivate.CreateSchedule)
	privateAPI.PUT("/schedule", scheduleControllerForPrivate.UpdateSchedule)
	privateAPI.DELETE("/schedule", scheduleControllerForPrivate.DeleteSchedule)
	
	//task
	publicAPI.GET("/tasks", taskControllerForPublic.GetTasks)
	internalAPI.GET("/tasks", taskControllerForInternal.GetTasks)
	internalAPI.POST("/task", taskControllerForInternal.CreateTask)
	internalAPI.PUT("/task", taskControllerForInternal.UpdateTask)
	internalAPI.DELETE("/task", taskControllerForInternal.DeleteTask)
	privateAPI.GET("/tasks", taskControllerForPrivate.GetTasks)
	privateAPI.POST("/task", taskControllerForPrivate.CreateTask)
	privateAPI.PUT("/task", taskControllerForPrivate.UpdateTask)
	privateAPI.DELETE("/task", taskControllerForPrivate.DeleteTask)
	
	//event
	publicAPI.GET("/events", eventControllerForPublic.GetEvents)
	internalAPI.GET("/events", eventControllerForInternal.GetEvents)
	internalAPI.POST("/event", eventControllerForInternal.CreateEvent)
	internalAPI.PUT("/event", eventControllerForInternal.UpdateEvent)
	internalAPI.DELETE("/event", eventControllerForInternal.DeleteEvent)
	privateAPI.GET("/events", eventControllerForPrivate.GetEvents)
	privateAPI.POST("/event", eventControllerForPrivate.CreateEvent)
	privateAPI.PUT("/event", eventControllerForPrivate.UpdateEvent)
	privateAPI.DELETE("/event", eventControllerForPrivate.DeleteEvent)
	

	return router
}