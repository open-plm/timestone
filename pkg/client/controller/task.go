package controller

import (
	"fmt"
	"log"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapTaskCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	bootstrapTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapTaskCmd
}

func InitCreateTaskCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return createTaskCmd
}

func InitCreateTaskCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return createTaskCmd
}

func InitCreateTaskCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return createTaskCmd
}

func InitGetTaskCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return getTaskCmd
}

func InitGetTaskCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return getTaskCmd
}

func InitGetTaskCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return getTaskCmd
}

func InitUpdateTaskCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTaskCmd
}

func InitUpdateTaskCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTaskCmd
}

func InitUpdateTaskCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "update the value of a key",
		Long:  "update the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTaskCmd
}

func InitDeleteTaskCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTaskCmd
}

func InitDeleteTaskCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTaskCmd
}

func InitDeleteTaskCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteTaskCmd := &cobra.Command{
		Use:   "task",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteTaskCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTaskCmd
}