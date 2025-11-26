package root

import (
	"github.com/byteorbit/terragrunt-atlantis-config/internal/diff"
	"github.com/byteorbit/terragrunt-atlantis-config/internal/generate"
	"github.com/byteorbit/terragrunt-atlantis-config/internal/version"
	"github.com/spf13/cobra"
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
