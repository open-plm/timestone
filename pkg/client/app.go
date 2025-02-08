package client

import (
	"github.com/open-plm/timestone/pkg/client/controller"
	"github.com/open-plm/timestone/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAppUser struct {
	Create    *cobra.Command
	Get       *cobra.Command
	Update    *cobra.Command
	Delete    *cobra.Command
}

func InitRootCmdForAppUser() *cobra.Command {
	var rootCmdForAppUser = &cobra.Command{
		Use:   "timestone-app",
		Short: "'timestone' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAppUser
}

func InitBaseCmdForAppUser() BaseCmdForAppUser {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create the value of a key",
		Long:  "create the value of a key",
	}
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update the value of a key",
		Long:  "update the value of a key",
	}
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
	}
	baseCmdForAppUser := BaseCmdForAppUser{
		Create:    createCmd,
		Get:       getCmd,
		Update:    updateCmd,
		Delete:    deleteCmd,
	}
	return baseCmdForAppUser
}

func ClientForAppUser(conf config.BaseConfig) {
	rootCmdForAppUser := InitRootCmdForAppUser()
	rootCmdForAppUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAppUser := InitBaseCmdForAppUser()

	//create
	createProgramCmdForAppUser := controller.InitCreateProgramCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createProgramCmdForAppUser)
	createProjectCmdForAppUser := controller.InitCreateProjectCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createProjectCmdForAppUser)
	createScheduleCmdForAppUser := controller.InitCreateScheduleCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createScheduleCmdForAppUser)
	createTaskCmdForAppUser := controller.InitCreateTaskCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createTaskCmdForAppUser)
	createEventCmdForAppUser := controller.InitCreateEventCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createEventCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Create)
	
	//get
	getProgramCmdForAppUser := controller.InitGetProgramCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getProgramCmdForAppUser)
	getProjectCmdForAppUser := controller.InitGetProjectCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getProjectCmdForAppUser)
	getScheduleCmdForAppUser := controller.InitGetScheduleCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getScheduleCmdForAppUser)
	getTaskCmdForAppUser := controller.InitGetTaskCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getTaskCmdForAppUser)
	getEventCmdForAppUser := controller.InitGetEventCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getEventCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Get)
	
	//update
	updateProgramCmdForAppUser := controller.InitUpdateProgramCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateProgramCmdForAppUser)
	updateProjectCmdForAppUser := controller.InitUpdateProjectCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateProjectCmdForAppUser)
	updateScheduleCmdForAppUser := controller.InitUpdateScheduleCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateScheduleCmdForAppUser)
	updateTaskCmdForAppUser := controller.InitUpdateTaskCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateTaskCmdForAppUser)
	updateEventCmdForAppUser := controller.InitUpdateEventCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateEventCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Update)
	
	//delete
	deleteProgramCmdForAppUser := controller.InitDeleteProgramCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteProgramCmdForAppUser)
	deleteProjectCmdForAppUser := controller.InitDeleteProjectCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteProjectCmdForAppUser)
	deleteScheduleCmdForAppUser := controller.InitDeleteScheduleCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteScheduleCmdForAppUser)
	deleteTaskCmdForAppUser := controller.InitDeleteTaskCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteTaskCmdForAppUser)
	deleteEventCmdForAppUser := controller.InitDeleteEventCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteEventCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Delete)

	rootCmdForAppUser.Execute()
}