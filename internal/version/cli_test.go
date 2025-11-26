package version_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/transcend-io/terragrunt-atlantis-config/internal/version"
)

func TestNewVersionCommand(t *testing.T) {
	rootCmd := &cobra.Command{
		Use: "terragrunt-atlantis-config",
	}
	const ver = "v1.2.3"

	cmd := version.New(rootCmd, ver)

	// Basic command metadata assertions
	if cmd.Use != "version" {
		t.Errorf("expected Use to be %q, got %q", "version", cmd.Use)
	}
	if cmd.Short != "Version of terragrunt-atlantis-config" {
		t.Errorf("unexpected Short: %q", cmd.Short)
	}
	if cmd.Long != "Version of terragrunt-atlantis-config" {
		t.Errorf("unexpected Long: %q", cmd.Long)
	}

	// Capture stdout since the Run function uses fmt.Println
	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	// Invoke the Run function
	cmd.Run(cmd, []string{})

	// Restore stdout and read what was printed
	w.Close()
	os.Stdout = origStdout

	outBytes, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read from pipe: %v", err)
	}
	output := strings.TrimSpace(string(outBytes))
	want := rootCmd.Use + " " + ver

	if output != want {
		t.Errorf("expected output %q, got %q", want, output)
	}
}
