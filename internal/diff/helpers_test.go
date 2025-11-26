package diff_test

import (
	"testing"

	"github.com/transcend-io/terragrunt-atlantis-config/internal/diff"
)

func TestWhenModifiedPatternFromDir_Root(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		want string
	}{
		{name: "empty", dir: "", want: "**/*"},
		{name: "dot", dir: ".", want: "**/*"},
		{name: "spaces only", dir: "   ", want: "**/*"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diff.WhenModifiedGlobFromDir(tt.dir)
			if got != tt.want {
				t.Fatalf("whenModifiedPatternFromDir(%q) = %q, want %q", tt.dir, got, tt.want)
			}
		})
	}
}

func TestWhenModifiedPatternFromDir_SingleLevel(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		want string
	}{
		{name: "simple", dir: "foo", want: "../**/*"},
		{name: "with_dot_prefix", dir: "./foo", want: "../**/*"},
		{name: "with_trailing_slash", dir: "foo/", want: "../**/*"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diff.WhenModifiedGlobFromDir(tt.dir)
			if got != tt.want {
				t.Fatalf("whenModifiedPatternFromDir(%q) = %q, want %q", tt.dir, got, tt.want)
			}
		})
	}
}

func TestWhenModifiedPatternFromDir_MultiLevel(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		want string
	}{
		{name: "two_levels", dir: "foo/bar", want: "../../**/*"},
		{name: "two_levels_with_dot", dir: "./foo/bar", want: "../../**/*"},
		{name: "two_levels_with_trailing_slash", dir: "foo/bar/", want: "../../**/*"},
		{name: "three_levels", dir: "a/b/c", want: "../../../**/*"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diff.WhenModifiedGlobFromDir(tt.dir)
			if got != tt.want {
				t.Fatalf("whenModifiedPatternFromDir(%q) = %q, want %q", tt.dir, got, tt.want)
			}
		})
	}
}
