// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
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

	for _, pkg := range packages {
		t.Run(strings.TrimPrefix(pkg, "./"), func(t *testing.T) {
			home := t.TempDir()
			runWrapper(t, home, pkg, "download")
			output := runWrapper(t, home, pkg, "version")
			version := strings.TrimSpace(string(output))
			if !strings.HasPrefix(version, "go version go") ||
				(requiresMicrosoftVersionSuffix(version) && !strings.Contains(version, "-microsoft ")) {

				t.Errorf("version output = %q, want Microsoft build of Go version output", version)
			}
			platform := runtime.GOOS + "/" + runtime.GOARCH
			if !strings.Contains(version, platform) {
				t.Errorf("version output = %q, want platform %q", version, platform)
			}
		})
	}
}

// requiresMicrosoftVersionSuffix reports whether the Go version in output
// requires the Microsoft version suffix.
//
// Parse the version locally to avoid adding go-infra as a dependency of the
// otherwise dependency-free go-dl module.
func requiresMicrosoftVersionSuffix(output string) bool {
	fields := strings.Fields(output)
	if len(fields) < 3 {
		return false
	}

	version := strings.TrimPrefix(fields[2], "go")
	majorText, minorText, ok := strings.Cut(version, ".")
	if !ok {
		return false
	}
	minorEnd := strings.IndexFunc(minorText, func(r rune) bool {
		return r < '0' || r > '9'
	})
	if minorEnd >= 0 {
		minorText = minorText[:minorEnd]
	}
	major, majorErr := strconv.Atoi(majorText)
	minor, minorErr := strconv.Atoi(minorText)
	if majorErr != nil || minorErr != nil {
		return false
	}
	return major > 1 || (major == 1 && minor >= 26)
}

func runWrapper(t *testing.T, home, pkg, arg string) []byte {
	t.Helper()

	cmd := exec.Command("go", "run", pkg, arg)
	cmd.Dir = filepath.Join("..", "..")
	cmd.Env = wrapperEnvironment(t, home)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run %s %s failed: %v\n%s", pkg, arg, err, output)
	}
	return output
}

func wrapperEnvironment(t *testing.T, home string) []string {
	t.Helper()

	// Disable telemetry before either Go toolchain can start a background
	// process that writes to the temporary home during cleanup.
	telemetryDir := filepath.Join(home, "telemetry")
	if err := os.MkdirAll(telemetryDir, 0o700); err != nil {
		t.Fatalf("create telemetry directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(telemetryDir, "mode"), []byte("off"), 0o600); err != nil {
		t.Fatalf("disable telemetry: %v", err)
	}
	return append(os.Environ(), "HOME="+home, "USERPROFILE="+home,
		"GODEBUG=ms_version=1", "TEST_TELEMETRY_DIR="+telemetryDir)
}

func TestWrapperEnvironmentDisablesTelemetry(t *testing.T) {
	inheritedTelemetryDir := t.TempDir()
	t.Setenv("TEST_TELEMETRY_DIR", inheritedTelemetryDir)
	home := t.TempDir()
	telemetryDir := filepath.Join(home, "telemetry")
	env := wrapperEnvironment(t, home)

	values := make(map[string]string)
	for _, entry := range env {
		name, value, _ := strings.Cut(entry, "=")
		values[name] = value
	}
	for name, want := range map[string]string{
		"HOME":               home,
		"USERPROFILE":        home,
		"GODEBUG":            "ms_version=1",
		"TEST_TELEMETRY_DIR": telemetryDir,
	} {
		if got := values[name]; got != want {
			t.Errorf("%s = %q, want %q", name, got, want)
		}
	}

	mode, err := os.ReadFile(filepath.Join(telemetryDir, "mode"))
	if err != nil {
		t.Fatal(err)
	}
	if string(mode) != "off" {
		t.Errorf("telemetry mode = %q, want off", mode)
	}

	cmd := exec.Command("go", "env", "-json", "GOTELEMETRY", "GOTELEMETRYDIR")
	cmd.Env = env
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("read Go telemetry settings: %v\n%s", err, output)
	}
	var settings struct {
		Mode string `json:"GOTELEMETRY"`
		Dir  string `json:"GOTELEMETRYDIR"`
	}
	if err := json.Unmarshal(output, &settings); err != nil {
		t.Fatal(err)
	}
	if settings.Mode != "" && settings.Mode != "off" {
		t.Errorf("Go telemetry mode = %q, want off", settings.Mode)
	}
	if settings.Dir != "" && settings.Dir != telemetryDir {
		t.Errorf("Go telemetry directory = %q, want %q", settings.Dir, telemetryDir)
	}

	files, err := os.ReadDir(inheritedTelemetryDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 0 {
		t.Errorf("inherited telemetry directory was modified: %v", files)
	}
}

func TestRequiresMicrosoftVersionSuffix(t *testing.T) {
	tests := []struct {
		output string
		want   bool
	}{
		{output: "go version go1.25.14-1 linux/amd64", want: false},
		{output: "go version go1.26.8-1 linux/amd64", want: true},
		{output: "go version go1.27rc1 linux/amd64", want: true},
		{output: "go version go2.0 linux/amd64", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.output, func(t *testing.T) {
			if got := requiresMicrosoftVersionSuffix(tt.output); got != tt.want {
				t.Errorf("requiresMicrosoftVersionSuffix(%q) = %t, want %t", tt.output, got, tt.want)
			}
		})
	}
}
