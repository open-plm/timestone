package client

import (
	"github.com/open-plm/timestone/pkg/client/controller"
	"github.com/open-plm/timestone/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAdminUser struct {
	Bootstrap *cobra.Command
	Create    *cobra.Command
	Get       *cobra.Command
	Update    *cobra.Command
	Delete    *cobra.Command
}

func InitRootCmdForAdminUser() *cobra.Command {
	var rootCmdForAdminUser = &cobra.Command{
		Use:   "timestone-admin",
		Short: "'timestone' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAdminUser
}

func InitBaseCmdForAdminUser() BaseCmdForAdminUser {
	bootstrapCmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
	}
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
	baseCmdForAdminUser := BaseCmdForAdminUser{
		Bootstrap: bootstrapCmd,
		Create:    createCmd,
		Get:       getCmd,
		Update:    updateCmd,
		Delete:    deleteCmd,
	}
	return baseCmdForAdminUser
}

func ClientForAdminUser(conf config.BaseConfig) {
	rootCmdForAdminUser := InitRootCmdForAdminUser()
	rootCmdForAdminUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAdminUser := InitBaseCmdForAdminUser()

	//bootstrap
	bootstrapProgramCmdForAdminUser := controller.InitBootstrapProgramCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapProgramCmdForAdminUser)
	bootstrapProjectCmdForAdminUser := controller.InitBootstrapProjectCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapProjectCmdForAdminUser)
	bootstrapScheduleCmdForAdminUser := controller.InitBootstrapScheduleCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapScheduleCmdForAdminUser)
	bootstrapTaskCmdForAdminUser := controller.InitBootstrapTaskCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapTaskCmdForAdminUser)
	bootstrapEventCmdForAdminUser := controller.InitBootstrapEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapEventCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Bootstrap)

	//create
	createProgramCmdForAdminUser := controller.InitCreateProgramCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createProgramCmdForAdminUser)
	createProjectCmdForAdminUser := controller.InitCreateProjectCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createProjectCmdForAdminUser)
	createScheduleCmdForAdminUser := controller.InitCreateScheduleCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createScheduleCmdForAdminUser)
	createTaskCmdForAdminUser := controller.InitCreateTaskCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createTaskCmdForAdminUser)
	createEventCmdForAdminUser := controller.InitCreateEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createEventCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Create)

	//get
	getProgramCmdForAdminUser := controller.InitGetProgramCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getProgramCmdForAdminUser)
	getProjectCmdForAdminUser := controller.InitGetProjectCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getProjectCmdForAdminUser)
	getScheduleCmdForAdminUser := controller.InitGetScheduleCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getScheduleCmdForAdminUser)
	getTaskCmdForAdminUser := controller.InitGetTaskCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getTaskCmdForAdminUser)
	getEventCmdForAdminUser := controller.InitGetEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getEventCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Get)

	//update
	updateProgramCmdForAdminUser := controller.InitUpdateProgramCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateProgramCmdForAdminUser)
	updateProjectCmdForAdminUser := controller.InitUpdateProjectCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateProjectCmdForAdminUser)
	updateScheduleCmdForAdminUser := controller.InitUpdateScheduleCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateScheduleCmdForAdminUser)
	updateTaskCmdForAdminUser := controller.InitUpdateTaskCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateTaskCmdForAdminUser)
	updateEventCmdForAdminUser := controller.InitUpdateEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateEventCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Update)
	
	//delete
	deleteProgramCmdForAdminUser := controller.InitDeleteProgramCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteProgramCmdForAdminUser)
	deleteProjectCmdForAdminUser := controller.InitDeleteProjectCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteProjectCmdForAdminUser)
	deleteScheduleCmdForAdminUser := controller.InitDeleteScheduleCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteScheduleCmdForAdminUser)
	deleteTaskCmdForAdminUser := controller.InitDeleteTaskCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteTaskCmdForAdminUser)
	deleteEventCmdForAdminUser := controller.InitDeleteEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteEventCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Delete)

	rootCmdForAdminUser.Execute()
}