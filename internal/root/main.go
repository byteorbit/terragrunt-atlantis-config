package root

import (
	"github.com/spf13/cobra"
	"github.com/transcend-io/terragrunt-atlantis-config/cmd/diff"
	"github.com/transcend-io/terragrunt-atlantis-config/cmd/generate"
	"github.com/transcend-io/terragrunt-atlantis-config/cmd/version"
)

func New(_version string) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:          "terragrunt-atlantis-config",
		Short:        "Generates Atlantis Config for Terragrunt projects",
		Long:         "Generates Atlantis Config for Terragrunt projects",
		SilenceUsage: true,
	}
	rootCmd.AddCommand(
		generate.New(),
		version.New(rootCmd, _version),
		diff.New(),
	)
	return rootCmd
}
