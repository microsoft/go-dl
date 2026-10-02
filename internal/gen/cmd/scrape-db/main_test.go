// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScrape(t *testing.T) {
	const manifest = "{\n  \"version\": \"1.22.0-1\"\n}\n"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/go1.22.0-1.assets.json" {
			_, _ = w.Write([]byte(manifest))
			return
		}
		_, _ = w.Write([]byte("<html>not found</html>"))
	}))
	defer server.Close()

	root := t.TempDir()
	s := scraper{
		client:    server.Client(),
		dbRoot:    root,
		urlFormat: server.URL + "/go1.%d.%d-%d.assets.json",
	}
	if err := s.scrape(context.Background()); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(filepath.Join(root, "go1.22", "go1.22.0-1.assets.json"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != manifest {
		t.Fatalf("manifest contents = %q, want %q", got, manifest)
	}
}

func TestWriteManifestRejectsChangedContents(t *testing.T) {
	root := t.TempDir()
	s := scraper{dbRoot: root}
	v := version{Series: 22, Patch: 0, Revision: 1}
	if err := s.writeManifest(v, []byte("original")); err != nil {
		t.Fatal(err)
	}

	err := s.writeManifest(v, []byte("changed"))
	if err == nil || !strings.Contains(err.Error(), "differs") {
		t.Fatalf("writeManifest error = %v, want differing manifest error", err)
	}
}
