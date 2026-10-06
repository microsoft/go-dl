// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestCalculateOutputs(t *testing.T) {
	packages := []packageInfo{
		{Dir: "internal/integration"},
		{Dir: "internal/version"},
		{Dir: "latest/msgo"},
		{Dir: "latest/msgo-previous"},
		{Dir: "latest/msgo1.27"},
		{Dir: "static/msgo1.27.1-2"},
	}

	tests := []struct {
		name         string
		changedFiles []string
		wantLive     []string
	}{
		{
			name:         "one command",
			changedFiles: []string{"static/msgo1.27.1-2/main.go"},
			wantLive:     []string{"./static/msgo1.27.1-2"},
		},
		{
			name: "multiple commands",
			changedFiles: []string{
				"latest/msgo/main.go",
				"latest/msgo1.27/main.go",
			},
			wantLive: []string{"./latest/msgo", "./latest/msgo1.27"},
		},
		{
			name:         "shared implementation",
			changedFiles: []string{"internal/version/version.go"},
			wantLive:     []string{"./latest/msgo", "./latest/msgo-previous"},
		},
		{
			name:         "live test implementation",
			changedFiles: []string{"internal/integration/live_test.go"},
			wantLive:     []string{"./latest/msgo", "./latest/msgo-previous"},
		},
		{
			name:         "module metadata",
			changedFiles: []string{"go.mod"},
			wantLive:     []string{"./latest/msgo", "./latest/msgo-previous"},
		},
		{
			name:         "generator",
			changedFiles: []string{"internal/gen/mkdl/mkdl.go"},
			wantLive:     []string{},
		},
		{
			name:         "workflow",
			changedFiles: []string{".github/workflows/ci.yml"},
			wantLive:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateOutputs(tt.changedFiles, packages, "1.18.x")
			if !reflect.DeepEqual(got.DLMatrix.GoVersion, []string{"1.18.x", ""}) {
				t.Errorf("Go versions = %q", got.DLMatrix.GoVersion)
			}
			if !reflect.DeepEqual(got.LivePackages, tt.wantLive) {
				t.Errorf("live packages = %q, want %q", got.LivePackages, tt.wantLive)
			}
		})
	}
}

func TestGitDiffNamesDisablesRenameDetection(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.name", "Test")
	runGit(t, root, "config", "user.email", "test@example.com")
	writeTestFile(t, root, "old.go")
	runGit(t, root, "add", "old.go")
	runGit(t, root, "commit", "-m", "old")
	base := runGit(t, root, "rev-parse", "HEAD")

	if err := os.Rename(filepath.Join(root, "old.go"), filepath.Join(root, "new.go")); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "-A")
	runGit(t, root, "commit", "-m", "rename")
	head := runGit(t, root, "rev-parse", "HEAD")

	got, err := gitDiffNames(root, base, head, false)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"new.go", "old.go"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("gitDiffNames() = %q, want %q", got, want)
	}
}

func TestGitDiffNamesUsesMergeBase(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.name", "Test")
	runGit(t, root, "config", "user.email", "test@example.com")
	writeTestFile(t, root, "common.go")
	runGit(t, root, "add", "common.go")
	runGit(t, root, "commit", "-m", "common")

	runGit(t, root, "checkout", "-b", "feature")
	writeTestFile(t, root, "feature.go")
	runGit(t, root, "add", "feature.go")
	runGit(t, root, "commit", "-m", "feature")
	featureHead := runGit(t, root, "rev-parse", "HEAD")

	runGit(t, root, "checkout", "-")
	writeTestFile(t, root, "base.go")
	runGit(t, root, "add", "base.go")
	runGit(t, root, "commit", "-m", "base")
	baseHead := runGit(t, root, "rev-parse", "HEAD")

	got, err := gitDiffNames(root, baseHead, featureHead, true)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, []string{"feature.go"}) {
		t.Errorf("gitDiffNames() = %q, want only feature changes", got)
	}
}

func TestGoModVersion(t *testing.T) {
	got, err := goModVersion("../../../../")
	if err != nil {
		t.Fatal(err)
	}
	if got != "1.18.x" {
		t.Errorf("goModVersion() = %q, want %q", got, "1.18.x")
	}
}

func writeTestFile(t *testing.T, root, name string) {
	t.Helper()
	content := []byte("package " + strings.TrimSuffix(name, ".go") + "\n")
	if err := os.WriteFile(filepath.Join(root, name), content, 0o600); err != nil {
		t.Fatal(err)
	}
}

func runGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	commandArgs := append([]string{"-C", root}, args...)
	out, err := exec.Command("git", commandArgs...).CombinedOutput()
	if err != nil {
		t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
	}
	return strings.TrimSpace(string(out))
}
