// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mkdl

import (
	"strings"
	"testing"

	"github.com/microsoft/go-dl/internal/gen/db"
	"github.com/microsoft/go-infra/assets"
	"github.com/microsoft/go-infra/goversion"
)

const testSHA256 = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

func TestArchiveTargets(t *testing.T) {
	entry := db.Entry{
		Version: goversion.New("1.27.1-2"),
		Manifest: assets.ToolsetBuild{Arches: []*assets.Arch{
			{Env: &assets.ArchEnv{GOOS: "windows", GOARCH: "amd64"}, URL: "https://example.test/go.zip", SHA256: testSHA256},
			{Env: nil, URL: "https://example.test/src.tar.gz", SHA256: "source"},
			{Env: &assets.ArchEnv{GOOS: "linux", GOARCH: "arm", GOARM: "6"}, URL: "https://example.test/go.tar.gz", SHA256: testSHA256},
		}},
	}

	got, err := archiveTargets(entry)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("archive target count = %d, want 2", len(got))
	}
	if got[0].GOOS != "linux" || got[0].GOARCH != "arm" || got[0].GOARM != "6" || got[0].URL != "https://example.test/go.tar.gz" || got[0].SHA256 != testSHA256 {
		t.Errorf("first archive target = %+v, want sorted linux/arm target", got[0])
	}
	if got[1].GOOS != "windows" || got[1].GOARCH != "amd64" {
		t.Errorf("second archive target = %+v, want sorted windows/amd64 target", got[1])
	}
}

func TestArchiveTargetsRejectsTargetConflict(t *testing.T) {
	entry := db.Entry{
		Version: goversion.New("1.27.1-2"),
		Manifest: assets.ToolsetBuild{Arches: []*assets.Arch{
			{Env: &assets.ArchEnv{GOOS: "linux", GOARCH: "arm", GOARM: "6"}, URL: "https://example.test/go-arm6.tar.gz", SHA256: testSHA256},
			{Env: &assets.ArchEnv{GOOS: "linux", GOARCH: "arm", GOARM: "7"}, URL: "https://example.test/go-arm7.tar.gz", SHA256: testSHA256},
		}},
	}

	_, err := archiveTargets(entry)
	if err == nil {
		t.Fatal("archiveTargets succeeded, want target conflict error")
	}
	for _, want := range []string{"1.27.1-2", "linux/arm", `GOARM "6" and "7"`} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not contain %q", err, want)
		}
	}
}

func TestArchiveTargetsRejectsInvalidSHA256(t *testing.T) {
	for _, test := range []struct {
		name     string
		checksum string
		want     string
	}{
		{name: "wrong length", checksum: "abcd", want: "length 4"},
		{name: "nonhex", checksum: strings.Repeat("z", 64), want: "invalid SHA256 checksum"},
	} {
		t.Run(test.name, func(t *testing.T) {
			entry := db.Entry{
				Version: goversion.New("1.27.1-2"),
				Manifest: assets.ToolsetBuild{Arches: []*assets.Arch{{
					Env:    &assets.ArchEnv{GOOS: "linux", GOARCH: "amd64"},
					URL:    "https://example.test/go.tar.gz",
					SHA256: test.checksum,
				}}},
			}
			_, err := archiveTargets(entry)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("archiveTargets error = %v, want error containing %q", err, test.want)
			}
		})
	}
}

func TestGenerateReadme(t *testing.T) {
	readme := []byte("before\n" + readmeVersionExamplesBegin + "\nstale\n" + readmeVersionExamplesEnd + "\nafter\n")
	current := db.Entry{Version: goversion.New("1.27.1-2")}
	previous := db.Entry{Version: goversion.New("1.26.8-1")}

	got, err := generateReadme(readme, current, previous)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"before\n",
		"`msgo` currently downloads Microsoft build of Go 1.27.1-2.",
		"go install github.com/microsoft/go-dl/latest/msgo-previous@latest",
		"go install github.com/microsoft/go-dl/latest/msgo1.27@latest",
		"go install github.com/microsoft/go-dl/latest/msgo1.26@latest",
		"go install github.com/microsoft/go-dl/static/msgo1.27.1-2@latest",
		"go run github.com/microsoft/go-dl/latest/msgo1.27 download",
		"after\n",
	} {
		if !strings.Contains(string(got), want) {
			t.Errorf("generated README does not contain %q", want)
		}
	}
	if strings.Contains(string(got), "stale") {
		t.Error("generated README retained stale content")
	}
	gotAgain, err := generateReadme(got, current, previous)
	if err != nil {
		t.Fatal(err)
	}
	if string(gotAgain) != string(got) {
		t.Error("generating README a second time changed its contents")
	}
}

func TestGenerateReadmeRequiresOneMarkerPair(t *testing.T) {
	current := db.Entry{Version: goversion.New("1.27.1-2")}
	previous := db.Entry{Version: goversion.New("1.26.8-1")}
	for _, readme := range []string{
		"no markers",
		readmeVersionExamplesBegin + readmeVersionExamplesBegin + readmeVersionExamplesEnd,
		readmeVersionExamplesEnd + readmeVersionExamplesBegin,
	} {
		if _, err := generateReadme([]byte(readme), current, previous); err == nil {
			t.Errorf("generateReadme(%q) succeeded, want marker error", readme)
		}
	}
}
