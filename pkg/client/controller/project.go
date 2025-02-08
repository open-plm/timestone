package controller

import (
	"fmt"
	"log"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapProjectCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapProjectCmd := &cobra.Command{
		Use:   "project",
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
	bootstrapProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapProjectCmd
}

func InitCreateProjectCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createProjectCmd := &cobra.Command{
		Use:   "project",
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
	createProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return createProjectCmd
}

func InitCreateProjectCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createProjectCmd := &cobra.Command{
		Use:   "project",
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
	createProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return createProjectCmd
}

func InitCreateProjectCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createProjectCmd := &cobra.Command{
		Use:   "project",
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
	createProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return createProjectCmd
}

func InitGetProjectCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getProjectCmd := &cobra.Command{
		Use:   "project",
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
	getProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return getProjectCmd
}

func InitGetProjectCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getProjectCmd := &cobra.Command{
		Use:   "project",
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
	getProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return getProjectCmd
}

func InitGetProjectCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getProjectCmd := &cobra.Command{
		Use:   "project",
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
	getProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return getProjectCmd
}

func InitUpdateProjectCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateProjectCmd := &cobra.Command{
		Use:   "project",
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
	updateProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProjectCmd
}

func InitUpdateProjectCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateProjectCmd := &cobra.Command{
		Use:   "project",
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
	updateProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProjectCmd
}

func InitUpdateProjectCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateProjectCmd := &cobra.Command{
		Use:   "project",
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
	updateProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProjectCmd
}

func InitDeleteProjectCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteProjectCmd := &cobra.Command{
		Use:   "project",
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
	deleteProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProjectCmd
}

func InitDeleteProjectCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteProjectCmd := &cobra.Command{
		Use:   "project",
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
	deleteProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProjectCmd
}

func InitDeleteProjectCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteProjectCmd := &cobra.Command{
		Use:   "project",
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
	deleteProjectCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProjectCmd
}