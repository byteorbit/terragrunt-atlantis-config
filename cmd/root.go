package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/transcend-io/terragrunt-atlantis-config/cmd/diff"
)

var (
	VERSION string
)

func NewRoot() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:          "terragrunt-atlantis-config",
		Short:        "Generates Atlantis Config for Terragrunt projects",
		Long:         "Generates Atlantis Config for Terragrunt projects",
		SilenceUsage: true,
	}
	rootCmd.AddCommand(
		New(),
		NewVersionCmd(rootCmd),
		diff.New(),
	)
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(rootCmd *cobra.Command, version string) {
	VERSION = version

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
