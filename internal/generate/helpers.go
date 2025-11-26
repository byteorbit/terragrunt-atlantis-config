package generate

import (
	"path/filepath"

	"github.com/gruntwork-io/terragrunt/util"
)

// addTerragruntValuesDependency ensures that when a terragrunt.hcl is tracked,
// we also track the associated terragrunt.values.hcl,
// so changes to stack value files will trigger Atlantis plans.
func addTerragruntValuesDependency(relativeDependencies *[]string, baseDir string, depAbsPath string) {
	depDir := filepath.Dir(depAbsPath)

	candidateFiles := []string{
		filepath.Join(depDir, "terragrunt.values.hcl"),
	}

	for _, valuesPath := range candidateFiles {
		if !util.FileExists(valuesPath) {
			continue
		}

		relValues, err := filepath.Rel(baseDir, valuesPath)
		if err != nil {
			continue
		}

		*relativeDependencies = append(*relativeDependencies, filepath.ToSlash(relValues))
	}
}
