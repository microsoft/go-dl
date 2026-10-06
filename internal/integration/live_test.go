// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestLiveWrappers(t *testing.T) {
	if os.Getenv("GO_DL_LIVE_TEST") == "" {
		t.Skip("set GO_DL_LIVE_TEST=1 to download and run Microsoft build of Go packages")
	}

	packages := strings.Fields(os.Getenv("GO_DL_LIVE_PACKAGES"))
	if len(packages) == 0 {
		t.Fatal("GO_DL_LIVE_PACKAGES must list at least one package")
	}

	home := t.TempDir()
	for _, pkg := range packages {
		t.Run(strings.TrimPrefix(pkg, "./"), func(t *testing.T) {
			runWrapper(t, home, pkg, "download")
			output := runWrapper(t, home, pkg, "version")
			version := strings.TrimSpace(string(output))
			if !strings.HasPrefix(version, "go version go") || !strings.Contains(version, "-microsoft ") {
				t.Errorf("version output = %q, want Microsoft build of Go version output", version)
			}
			platform := runtime.GOOS + "/" + runtime.GOARCH
			if !strings.Contains(version, platform) {
				t.Errorf("version output = %q, want platform %q", version, platform)
			}
		})
	}
}

func runWrapper(t *testing.T, home, pkg, arg string) []byte {
	t.Helper()

	cmd := exec.Command("go", "run", pkg, arg)
	cmd.Dir = filepath.Join("..", "..")
	cmd.Env = append(os.Environ(), "HOME="+home, "USERPROFILE="+home, "GODEBUG=ms_version=1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run %s %s failed: %v\n%s", pkg, arg, err, output)
	}
	return output
}
