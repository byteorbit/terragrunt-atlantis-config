package diff

import (
	"fmt"

	"github.com/gruntwork-io/terragrunt/util"
	"github.com/spf13/cobra"
)

const (
	baseWorkDirFlag            = "base-work-dir"
	targetWorkDirFlag          = "target-work-dir"
	baseAtlantisConfPathFlag   = "base-atlantis-config-path"
	targetAtlantisConfPathFlag = "target-atlantis-config-path"
	outputPathFlag             = "output"

	baseWorkDirFlagError            = "--%s must be a directory"
	targetWorkDirFlagError          = "--%s must be a directory"
	baseAtlantisConfPathFlagError   = "--%s must be a file"
	targetAtlantisConfPathFlagError = "--%s must be a file"
)

func New() *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:   "diff",
		Short: "Compares Atlantis config files.",
		Long:  `TODO`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !util.IsDir(opts.BaseWorkDir) {
				return fmt.Errorf(baseWorkDirFlagError, baseWorkDirFlag)
			}
			if !util.IsDir(opts.TargetWorkDir) {
				return fmt.Errorf(targetWorkDirFlagError, targetWorkDirFlag)
			}
			if !util.IsFile(opts.BaseAtlantisConfPath) {
				return fmt.Errorf(baseAtlantisConfPathFlagError, baseAtlantisConfPathFlag)
			}
			if !util.IsFile(opts.TargetAtlantisConfPath) {
				return fmt.Errorf(targetAtlantisConfPathFlagError, targetAtlantisConfPathFlag)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return CreateComparison(opts)
		},
	}

	// Bind flags directly to fields on opts
	f := cmd.Flags()
	f.StringVar(&opts.BaseWorkDir, baseWorkDirFlag, "", "Working directory for base project.")
	f.StringVar(&opts.TargetWorkDir, targetWorkDirFlag, "", "Working directory for target project.")
	f.StringVar(&opts.BaseAtlantisConfPath, baseAtlantisConfPathFlag, "", "Absolute path to the Atlantis config generated for the base repo.")
	f.StringVar(&opts.TargetAtlantisConfPath, targetAtlantisConfPathFlag, "", "Absolute path to the Atlantis config generated for the target repo.")
	f.StringVar(&opts.OutputPath, outputPathFlag, "", "Path of the file where configuration will be generated. Default is not to write to file")

	_ = cmd.MarkFlagRequired(baseWorkDirFlag)
	_ = cmd.MarkFlagRequired(targetWorkDirFlag)
	_ = cmd.MarkFlagRequired(baseAtlantisConfPathFlag)
	_ = cmd.MarkFlagRequired(targetAtlantisConfPathFlag)

	return cmd
}
