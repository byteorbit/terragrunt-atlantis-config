package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New(rootCmd *cobra.Command, version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Version of terragrunt-atlantis-config",
		Long:  "Version of terragrunt-atlantis-config",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(rootCmd.Use + " " + version)
		},
	}
}
