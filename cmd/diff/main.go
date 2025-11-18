package diff

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/runatlantis/atlantis/server/core/config/raw"
	"gopkg.in/yaml.v3"

	"github.com/bmatcuk/doublestar"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Options struct {
	BaseWorkDir            string
	TargetWorkDir          string
	BaseAtlantisConfPath   string
	TargetAtlantisConfPath string
	OutputPath             string
}

func CreateComparison(opts *Options) error {
	/** TODO
	 */
	var (
		baseRoot string
		newRoot  string
		//maxParallel int = 8
	)
	var err error
	baseRoot, err = filepath.Abs(opts.BaseWorkDir)
	checkErr(err)
	newRoot, err = filepath.Abs(opts.TargetWorkDir)
	checkErr(err)

	baseCfg, err := LoadRawRepoCfg(opts.BaseAtlantisConfPath)
	checkErr(err)
	targetCfg, err := LoadRawRepoCfg(opts.TargetAtlantisConfPath)
	checkErr(err)

	baseByName := make(map[string]raw.Project, len(baseCfg.Projects))
	for _, p := range baseCfg.Projects {
		baseByName[*p.Name] = p
	}

	eg := errgroup.Group{}
	//eg.SetLimit(maxParallel)

	for i := range targetCfg.Projects {
		idx := i
		p := &targetCfg.Projects[idx]
		eg.Go(func() error {
			changed, err := projectChanged(
				p,
				baseByName,
				baseRoot,
				newRoot,
			)
			if err != nil {
				return err
			}
			p.Autoplan.Enabled = &changed // TODO determine if this should always be true
			if changed {
				p.Autoplan.WhenModified = []string{
					".atlantis-project-changed",
				}
			} else {
				p.Autoplan.WhenModified = []string{}
			}
			return nil
		})
	}
	checkErr(eg.Wait())

	yamlBytes, err := yaml.Marshal(&targetCfg)
	if err != nil {
		return err
	}

	// Ensure newline characters are correct on windows machines, as the json encoding function in the stdlib
	// uses "\n" for all newlines regardless of OS: https://github.com/golang/go/blob/master/src/encoding/json/stream.go#L211-L217
	yamlString := string(yamlBytes)
	// TODO load raw from this yaml and then dump that instead.

	if strings.Contains(runtime.GOOS, "windows") {
		yamlString = strings.ReplaceAll(yamlString, "\n", "\r\n")
	}

	// Write output
	if len(opts.OutputPath) != 0 {
		err = os.WriteFile(opts.OutputPath, []byte(yamlString), 0644)
		if err != nil {
			return err
		}
	} else {
		log.Println(yamlString)
	}
	return nil
}

func projectChanged(newProj *raw.Project, baseByName map[string]raw.Project, baseRoot, newRoot string) (bool, error) {
	newDir := normalizeSlash(*newProj.Dir)
	// Step 1: is this a new project
	baseProj, found := baseByName[*newProj.Name]
	if !found {
		return true, nil
	}

	// Step 2: compare unique set of strings in when_modified & depends_on
	newSet := toSet(unique(newProj.Autoplan.WhenModified))
	baseSet := toSet(unique(baseProj.Autoplan.WhenModified))
	if !stringSetsEqual(newSet, baseSet) {
		return true, nil
	}
	newSet = toSet(unique(newProj.DependsOn))
	baseSet = toSet(unique(baseProj.DependsOn))
	if !stringSetsEqual(newSet, baseSet) {
		return true, nil
	}

	// Step 3: expand full set of files for both
	newFiles, err := expandAllGlobs(newRoot, newDir, newProj.Autoplan.WhenModified)
	if err != nil {
		return false, fmt.Errorf("expand new globs (dir=%s): %w", newDir, err)
	}
	baseFiles, err := expandAllGlobs(baseRoot, newDir, baseProj.Autoplan.WhenModified)
	if err != nil {
		return false, fmt.Errorf("expand base globs (dir=%s): %w", newDir, err)
	}

	// Compare the full set of relative paths
	if len(newFiles) != len(baseFiles) || !stringSetsEqual(sliceToSet(newFiles), sliceToSet(baseFiles)) {
		return true, nil
	}

	// Step 4: Compare SHA256 for each matched file (since sets are equal, just iterate one side)
	// We hash files by the relative key (relative to root). Then compare across roots.
	hashNew, err := hashFiles(newRoot, newFiles)
	if err != nil {
		return false, fmt.Errorf("hash new files: %w", err)
	}
	hashBase, err := hashFiles(baseRoot, baseFiles)
	if err != nil {
		return false, fmt.Errorf("hash base files: %w", err)
	}

	for rel, hNew := range hashNew {
		hBase, ok := hashBase[rel]
		if !ok {
			// Shouldn't happen, sets equal — treat as change if it does.
			return true, nil
		}
		if !bytes.Equal(hNew, hBase) {
			return true, nil
		}
	}

	return false, nil
}

func unique(xs []string) []string {
	if len(xs) == 0 {
		return xs
	}
	m := make(map[string]struct{}, len(xs))
	for _, s := range xs {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		m[s] = struct{}{}
	}
	out := make([]string, 0, len(m))
	for s := range m {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func toSet(xs []string) map[string]struct{} {
	m := make(map[string]struct{}, len(xs))
	for _, s := range xs {
		m[s] = struct{}{}
	}
	return m
}

func sliceToSet(xs []string) map[string]struct{} {
	m := make(map[string]struct{}, len(xs))
	for _, s := range xs {
		m[s] = struct{}{}
	}
	return m
}

func stringSetsEqual(a, b map[string]struct{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}

func normalizeSlash(p string) string {
	return filepath.ToSlash(filepath.Clean(p))
}

// expandAllGlobs expands each when_modified pattern relative to root/projectDir,
// returns a sorted list of *relative paths to root* (stable keys for cross-root comparison).
func expandAllGlobs(root, projectDir string, patterns []string) ([]string, error) {
	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	projAbs := filepath.Join(rootAbs, filepath.FromSlash(projectDir))

	seen := map[string]struct{}{}
	for _, pat := range patterns {
		pat = filepath.ToSlash(pat)
		globPattern := filepath.Join(projAbs, pat)

		matches, err := doublestar.Glob(globPattern)
		if err != nil {
			return nil, fmt.Errorf("glob %q: %w", globPattern, err)
		}

		for _, abs := range matches {
			abs = filepath.Clean(abs)
			rel, relErr := filepath.Rel(rootAbs, abs)
			if relErr != nil {
				rel = filepath.ToSlash(abs)
			}
			seen[normalizeSlash(rel)] = struct{}{}
		}
	}

	out := make([]string, 0, len(seen))
	for k := range seen {
		out = append(out, k)
	}
	sort.Strings(out)
	return out, nil
}

func hashFiles(root string, rels []string) (map[string][]byte, error) {
	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	out := make(map[string][]byte, len(rels))
	for _, rel := range rels {
		full := filepath.Join(rootAbs, filepath.FromSlash(rel))
		h, err := fileSHA256(full)
		if err != nil {
			// If a file disappeared between globbing and hashing, treat as an error
			if errors.Is(err, os.ErrNotExist) {
				return nil, fmt.Errorf("file missing during hashing: %s", full)
			}
			return nil, err
		}
		out[rel] = h
	}
	return out, nil
}

func fileSHA256(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func checkErr(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
