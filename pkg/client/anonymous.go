package client

import (
	"github.com/open-plm/timestone/pkg/client/controller"
	"github.com/open-plm/timestone/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAnonymousUser struct {
	Get       *cobra.Command
}

func InitRootCmdForAnonymousUser() *cobra.Command {
	var rootCmdForAnonymousUser = &cobra.Command{
		Use:   "timestone-anonymous",
		Short: "'timestone' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAnonymousUser
}

func InitBaseCmdForAnonymousUser() BaseCmdForAnonymousUser {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	baseCmdForAnonymousUser := BaseCmdForAnonymousUser{
		Get:       getCmd,
	}
	return baseCmdForAnonymousUser
}

func ClientForAnonymousUser(conf config.BaseConfig) {
	rootCmdForAnonymousUser := InitRootCmdForAnonymousUser()
	rootCmdForAnonymousUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAnonymousUser := InitBaseCmdForAnonymousUser()

	//get
	getProgramCmdForAnonymousUser := controller.InitGetProgramCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getProgramCmdForAnonymousUser)
	getProjectCmdForAnonymousUser := controller.InitGetProjectCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getProjectCmdForAnonymousUser)
	getScheduleCmdForAnonymousUser := controller.InitGetScheduleCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getScheduleCmdForAnonymousUser)
	getTaskCmdForAnonymousUser := controller.InitGetTaskCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getTaskCmdForAnonymousUser)
	getEventCmdForAnonymousUser := controller.InitGetEventCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getEventCmdForAnonymousUser)
	rootCmdForAnonymousUser.AddCommand(baseCmdForAnonymousUser.Get)
	
	rootCmdForAnonymousUser.Execute()
}