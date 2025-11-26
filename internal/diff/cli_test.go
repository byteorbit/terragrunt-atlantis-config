package diff_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	diff2 "github.com/transcend-io/terragrunt-atlantis-config/internal/diff"
)

func runWithFlags(args []string) error {
	diffCmd := diff2.New()
	diffCmd.SetArgs(args)
	if err := diffCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func TestMissingFlags(t *testing.T) {
	err := runWithFlags([]string{})
	if err == nil {
		t.Error("Expected to fail")
	}
}

func runBasicDiffTest(t *testing.T, testDir string) {
	runBasicDiffTestWithDirs(t, testDir, "base", "target")
}
func runBasicDiffTestWithDirs(t *testing.T, testDir, _baseDir, _targetDir string) {
	expectedYaml := "expected.yaml"
	outputYaml := filepath.Join(t.TempDir(), "atlantis.yaml")
	baseDir := filepath.Join("testdata", testDir, _baseDir)
	targetDir := filepath.Join("testdata", testDir, _targetDir)
	// TODO copy the base and target data dir into the artifact dirs
	//  This in conjunction with generating the stacks and the atlantis.yaml in those directories, instead of committing.
	baseAtlantis := filepath.Join(baseDir, "atlantis.yaml")
	targetAtlantis := filepath.Join(targetDir, "atlantis.yaml")
	err := runWithFlags([]string{
		"--base-work-dir",
		baseDir,
		"--target-work-dir",
		targetDir,
		"--base-atlantis-config-path",
		baseAtlantis,
		"--target-atlantis-config-path",
		targetAtlantis,
		"--output",
		outputYaml,
	})
	if err != nil {
		t.Fatalf("Failed run diff")
	}

	expectedContent, err := diff2.LoadRawRepoCfg(filepath.Join("testdata", testDir, expectedYaml))
	if err != nil {
		t.Fatal(err)
	}
	actualContent, err := diff2.LoadRawRepoCfg(outputYaml)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedContent, actualContent)
}

func TestEmptyBase(t *testing.T) {
	runBasicDiffTest(t, "emptyBase")
}

func TestUnchangesFiles(t *testing.T) {
	runBasicDiffTestWithDirs(t, "unchangedFiles", "base", "base")
}

func TestChangesFilesGlobs(t *testing.T) {
	runBasicDiffTest(t, "changedFilesGlobs")
}

func TestChangesFilesHashes(t *testing.T) {
	runBasicDiffTest(t, "changedFilesHashes")
}

func TestChangesFilesSet(t *testing.T) {
	runBasicDiffTest(t, "changedFilesSet")
}

func TestChangesRefsSelf(t *testing.T) {
	runBasicDiffTest(t, "changedRefsSelf")
}

func TestChangesRefsDependent(t *testing.T) {
	runBasicDiffTest(t, "changedRefsDependent")
}

func TestChangesRefsDeep(t *testing.T) {
	runBasicDiffTest(t, "changedRefsDeep")
}

func TestChangesRefsPartial(t *testing.T) {
	runBasicDiffTest(t, "changedRefsPartial")
}
