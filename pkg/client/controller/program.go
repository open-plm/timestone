package controller

import (
	"fmt"
	"log"

	"github.com/open-plm/timestone/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapProgramCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapProgramCmd := &cobra.Command{
		Use:   "program",
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
	bootstrapProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapProgramCmd
}

func InitCreateProgramCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createProgramCmd := &cobra.Command{
		Use:   "program",
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
	createProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return createProgramCmd
}

func InitCreateProgramCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createProgramCmd := &cobra.Command{
		Use:   "program",
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
	createProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return createProgramCmd
}

func InitCreateProgramCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createProgramCmd := &cobra.Command{
		Use:   "program",
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
	createProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return createProgramCmd
}

func InitGetProgramCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getProgramCmd := &cobra.Command{
		Use:   "program",
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
	getProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return getProgramCmd
}

func InitGetProgramCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getProgramCmd := &cobra.Command{
		Use:   "program",
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
	getProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return getProgramCmd
}

func InitGetProgramCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getProgramCmd := &cobra.Command{
		Use:   "program",
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
	getProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return getProgramCmd
}

func InitUpdateProgramCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateProgramCmd := &cobra.Command{
		Use:   "program",
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
	updateProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProgramCmd
}

func InitUpdateProgramCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateProgramCmd := &cobra.Command{
		Use:   "program",
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
	updateProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProgramCmd
}

func InitUpdateProgramCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateProgramCmd := &cobra.Command{
		Use:   "program",
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
	updateProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProgramCmd
}

func InitDeleteProgramCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteProgramCmd := &cobra.Command{
		Use:   "program",
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
	deleteProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProgramCmd
}

func InitDeleteProgramCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteProgramCmd := &cobra.Command{
		Use:   "program",
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
	deleteProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProgramCmd
}

func InitDeleteProgramCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteProgramCmd := &cobra.Command{
		Use:   "program",
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
	deleteProgramCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProgramCmd
}