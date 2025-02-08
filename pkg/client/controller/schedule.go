package controller

import (
	"fmt"
	"log"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapScheduleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	bootstrapScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapScheduleCmd
}

func InitCreateScheduleCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	createScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return createScheduleCmd
}

func InitCreateScheduleCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	createScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return createScheduleCmd
}

func InitCreateScheduleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	createScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return createScheduleCmd
}

func InitGetScheduleCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	getScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return getScheduleCmd
}

func InitGetScheduleCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	getScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return getScheduleCmd
}

func InitGetScheduleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	getScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return getScheduleCmd
}

func InitUpdateScheduleCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	updateScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return updateScheduleCmd
}

func InitUpdateScheduleCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	updateScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return updateScheduleCmd
}

func InitUpdateScheduleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	updateScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return updateScheduleCmd
}

func InitDeleteScheduleCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	deleteScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteScheduleCmd
}

func InitDeleteScheduleCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	deleteScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteScheduleCmd
}

func InitDeleteScheduleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteScheduleCmd := &cobra.Command{
		Use:   "schedule",
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
	deleteScheduleCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteScheduleCmd
}