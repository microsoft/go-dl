// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"testing"
	"testing/fstest"
)

func TestLoad(t *testing.T) {
	fsys := fstest.MapFS{
		"go1.26/go1.26.3-2.assets.json": {
			Data: []byte(`{"version":"1.26.3-2"}`),
		},
		"go1.26/go1.26.3-1.assets.json": {
			Data: []byte(`{"version":"1.26.3-1"}`),
		},
		"go1.25/go1.25.9-1.assets.json": {
			Data: []byte(`{"version":"1.25.9-1"}`),
		},
	}

	entries, err := load(fsys)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(entries), 3; got != want {
		t.Fatalf("entry count = %d, want %d", got, want)
	}
	for i, want := range []string{"1.25.9-1", "1.26.3-1", "1.26.3-2"} {
		if got := entries[i].Version.Full(); got != want {
			t.Errorf("entry %d version = %q, want %q", i, got, want)
		}
	}
}

func TestLoadRejectsMismatchedVersion(t *testing.T) {
	fsys := fstest.MapFS{
		"go1.26/go1.26.3-1.assets.json": {
			Data: []byte(`{"version":"1.26.4-1"}`),
		},
	}

	if _, err := load(fsys); err == nil {
		t.Fatal("load succeeded, want version mismatch error")
	}
}

func TestLoadRejectsMismatchedSeriesDirectory(t *testing.T) {
	fsys := fstest.MapFS{
		"go1.25/go1.26.3-1.assets.json": {
			Data: []byte(`{"version":"1.26.3-1"}`),
		},
	}

	if _, err := load(fsys); err == nil {
		t.Fatal("load succeeded, want path mismatch error")
	}
}
