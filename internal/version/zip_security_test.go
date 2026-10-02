// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

func TestUnpackZipRejectsInvalidPath(t *testing.T) {
	tempDir := t.TempDir()
	archiveFile := filepath.Join(tempDir, "go.zip")
	file, err := os.Create(archiveFile)
	if err != nil {
		t.Fatal(err)
	}
	zw := zip.NewWriter(file)
	entry, err := zw.Create("go/../../outside")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := entry.Write([]byte("unexpected")); err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	targetDir := filepath.Join(tempDir, "sdk", "nested")
	err = unpackZip(targetDir, archiveFile)
	if err == nil {
		t.Fatal("unpackZip succeeded with a traversing path")
	}
	if _, err := os.Stat(filepath.Join(tempDir, "outside")); !os.IsNotExist(err) {
		t.Errorf("traversing path was written outside target directory: %v", err)
	}
}
