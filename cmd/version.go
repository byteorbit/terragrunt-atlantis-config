package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCmd(rootCmd *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Version of terragrunt-atlantis-config",
		Long:  "Version of terragrunt-atlantis-config",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(rootCmd.Use + " " + VERSION)
		},
	}
}
