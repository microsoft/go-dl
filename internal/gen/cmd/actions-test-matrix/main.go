// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command actions-test-matrix produces the dynamic inputs for go-dl CI.
//
// The ordinary test matrix always runs ./... at the minimum and current Go
// versions. Live tests are limited to directly changed downloader commands,
// plus representative commands when code used by every wrapper changes.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type packageInfo struct {
	Dir string
}

type goListPackage struct {
	Dir string
}

type dlMatrix struct {
	// An empty version asks the setup action to install its default toolchain.
	GoVersion []string `json:"go-version"`
}

type outputs struct {
	DLMatrix     dlMatrix
	LivePackages []string
}

type sharedChangeLiveTests struct {
	packages             []string
	triggeredByChangesTo []string
}

// liveTests is the repository policy for networked tests.
var liveTests = struct {
	packageRoots []string
	sharedChange sharedChangeLiveTests
}{
	packageRoots: []string{
		"latest",
		"static",
	},
	sharedChange: sharedChangeLiveTests{
		packages: []string{
			"latest/msgo",
			"latest/msgo-previous",
		},
		// Each entry matches that Git path and anything below it.
		triggeredByChangesTo: []string{
			"internal/version",
			"internal/integration",
			"go.mod",
			"go.sum",
		},
	},
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	base := flag.String("base", "", "Base commit SHA for the diff.")
	head := flag.String("head", "", "Head commit SHA for the diff.")
	mergeBase := flag.Bool("merge-base", false, "Compare head with the merge base of base and head.")
	flag.Parse()

	if *base == "" || *head == "" {
		return fmt.Errorf("-base and -head are required")
	}

	root, err := repoRoot()
	if err != nil {
		return fmt.Errorf("finding repo root: %w", err)
	}
	changedFiles, err := gitDiffNames(root, *base, *head, *mergeBase)
	if err != nil {
		return fmt.Errorf("getting changed files: %w", err)
	}
	packages, err := goListPackages(root)
	if err != nil {
		return fmt.Errorf("listing dl module packages: %w", err)
	}
	minGoVersion, err := goModVersion(root)
	if err != nil {
		return fmt.Errorf("reading dl module Go version: %w", err)
	}

	return writeOutputs(calculateOutputs(changedFiles, packages, minGoVersion))
}

func repoRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("running git rev-parse: %w: %s", err, strings.TrimSpace(string(out)))
	}
	return strings.TrimSpace(string(out)), nil
}

func gitDiffNames(root, base, head string, mergeBase bool) ([]string, error) {
	// GitHub uses an all-zero base for the first push of a branch. Compare that
	// commit with Git's empty tree so every file is treated as added.
	if strings.Trim(base, "0") == "" {
		cmd := exec.Command("git", "-C", root, "hash-object", "-t", "tree", "--stdin")
		cmd.Stdin = strings.NewReader("")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("finding empty Git tree: %w: %s", err, strings.TrimSpace(string(out)))
		}
		base = strings.TrimSpace(string(out))
		mergeBase = false
	}

	revisions := []string{base, head}
	if mergeBase {
		revisions = []string{base + "..." + head}
	}
	// Disabling rename detection reports both sides of a move. This matters when
	// a command moves into or out of a live-tested package.
	args := append([]string{"-C", root, "diff", "--name-only", "-z", "--no-renames"}, revisions...)
	args = append(args, "--")
	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("running git diff: %w: %s", err, strings.TrimSpace(string(out)))
	}

	var files []string
	for _, name := range strings.Split(string(out), "\x00") {
		if name != "" {
			files = append(files, filepath.ToSlash(name))
		}
	}
	return files, nil
}

func goListPackages(root string) ([]packageInfo, error) {
	// Use Go's package view instead of walking directories. This excludes the
	// nested generator module and ignores deleted directories from the diff.
	cmd := exec.Command("go", "list", "-json", "./...")
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("running go list: %w: %s", err, strings.TrimSpace(string(out)))
	}

	decoder := json.NewDecoder(strings.NewReader(string(out)))
	var packages []packageInfo
	for {
		var listed goListPackage
		if err := decoder.Decode(&listed); err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("decoding go list output: %w", err)
		}

		relDir, err := filepath.Rel(root, listed.Dir)
		if err != nil {
			return nil, fmt.Errorf("finding relative package path for %q: %w", listed.Dir, err)
		}
		packages = append(packages, packageInfo{Dir: filepath.ToSlash(relDir)})
	}
	return packages, nil
}

func goModVersion(root string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f={{.GoVersion}}")
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("running go list: %w: %s", err, strings.TrimSpace(string(out)))
	}

	parts := strings.Split(strings.TrimSpace(string(out)), ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("unexpected Go version %q", strings.TrimSpace(string(out)))
	}
	return parts[0] + "." + parts[1] + ".x", nil
}

func calculateOutputs(changedFiles []string, packages []packageInfo, minGoVersion string) outputs {
	return outputs{
		DLMatrix: dlMatrix{
			GoVersion: []string{minGoVersion, ""},
		},
		LivePackages: changedLivePackages(changedFiles, packages),
	}
}

func changedLivePackages(changedFiles []string, packages []packageInfo) []string {
	changed := make(map[string]bool)
	available := make(map[string]bool)
	// Direct changes test the command itself. Intersecting the diff with current
	// packages also means a deleted command is not sent to go run.
	for _, pkg := range packages {
		if !matchesAnyPath(pkg.Dir, liveTests.packageRoots) {
			continue
		}
		available[pkg.Dir] = true
		for _, file := range changedFiles {
			if pathHasPrefix(file, pkg.Dir) {
				changed["./"+pkg.Dir] = true
				break
			}
		}
	}

	if anyFileMatches(changedFiles, requiresSharedLiveTests) {
		// The floating current and previous commands exercise code used by every
		// wrapper without retesting every historical package.
		for _, pkg := range liveTests.sharedChange.packages {
			if available[pkg] {
				changed["./"+pkg] = true
			}
		}
	}

	result := make([]string, 0, len(changed))
	for pkg := range changed {
		result = append(result, pkg)
	}
	sort.Strings(result)
	return result
}

func requiresSharedLiveTests(file string) bool {
	return matchesAnyPath(file, liveTests.sharedChange.triggeredByChangesTo)
}

func matchesAnyPath(file string, paths []string) bool {
	for _, path := range paths {
		if pathHasPrefix(file, path) {
			return true
		}
	}
	return false
}

func anyFileMatches(files []string, match func(string) bool) bool {
	for _, file := range files {
		if match(file) {
			return true
		}
	}
	return false
}

func pathHasPrefix(file, dir string) bool {
	return file == dir || strings.HasPrefix(file, dir+"/")
}

func writeOutputs(result outputs) (err error) {
	// GitHub Actions reads step outputs from GITHUB_OUTPUT. Writing to stdout
	// keeps the command useful for local inspection.
	w, err := outputWriter()
	if err != nil {
		return err
	}
	if w != os.Stdout {
		defer func() {
			if closeErr := w.Close(); closeErr != nil && err == nil {
				err = fmt.Errorf("closing GITHUB_OUTPUT: %w", closeErr)
			}
		}()
	}

	matrixJSON, err := json.Marshal(result.DLMatrix)
	if err != nil {
		return fmt.Errorf("marshaling dl matrix: %w", err)
	}
	livePackagesJSON, err := json.Marshal(result.LivePackages)
	if err != nil {
		return fmt.Errorf("marshaling live packages: %w", err)
	}
	_, err = fmt.Fprintf(
		w,
		"dl-matrix=%s\nlive-packages=%s\nhas-live=%t\n",
		matrixJSON,
		livePackagesJSON,
		len(result.LivePackages) > 0,
	)
	return err
}

func outputWriter() (*os.File, error) {
	if path := os.Getenv("GITHUB_OUTPUT"); path != "" {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o600)
		if err != nil {
			return nil, fmt.Errorf("opening GITHUB_OUTPUT: %w", err)
		}
		return file, nil
	}
	return os.Stdout, nil
}
