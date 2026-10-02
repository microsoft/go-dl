// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import (
	"archive/zip"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// MicrosoftArchive describes the platform archive selected by a generated wrapper.
type MicrosoftArchive struct {
	URL    string
	SHA256 string
}

// RunMicrosoft runs a Microsoft build of Go. command is the wrapper command
// invoked by the user, and version is the specific Microsoft build to run.
func RunMicrosoft(command, version string, archive MicrosoftArchive) {
	log.SetFlags(0)

	root, err := goroot(version)
	if err != nil {
		log.Fatalf("%s: %v", command, err)
	}

	if len(os.Args) == 2 && os.Args[1] == "download" {
		if err := installMicrosoft(root, command, version, archive); err != nil {
			log.Fatalf("%s: download failed: %v", command, err)
		}
		os.Exit(0)
	}

	if _, err := os.Stat(filepath.Join(root, unpackedOkay)); err != nil {
		log.Fatalf("%s: not downloaded. Run '%s download' to install to %v", command, command, root)
	}

	// Preserve the user's GOTOOLCHAIN setting, matching upstream version.Run.
	runGo(root, "")
}

func installMicrosoft(targetDir, command, version string, archive MicrosoftArchive) error {
	if _, err := os.Stat(filepath.Join(targetDir, unpackedOkay)); err == nil {
		log.Printf("%s: already downloaded in %v", command, targetDir)
		return nil
	}
	if err := archive.validate(version); err != nil {
		return err
	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	res, err := http.Head(archive.URL)
	if err != nil {
		return err
	}
	res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return fmt.Errorf("no binary release of %v for %v/%v at %v", version, getOS(), runtime.GOARCH, archive.URL)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %v checking size of %v", http.StatusText(res.StatusCode), archive.URL)
	}
	base := path.Base(archive.URL)
	archiveFile := filepath.Join(targetDir, base)
	if fi, err := os.Stat(archiveFile); err != nil || fi.Size() != res.ContentLength {
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		if err := copyFromURL(archiveFile, archive.URL); err != nil {
			return fmt.Errorf("error downloading %v: %v", archive.URL, err)
		}
		fi, err = os.Stat(archiveFile)
		if err != nil {
			return err
		}
		if fi.Size() != res.ContentLength {
			return fmt.Errorf("downloaded file %s size %v doesn't match server size %v", archiveFile, fi.Size(), res.ContentLength)
		}
	}
	if err := verifySHA256(archiveFile, archive.SHA256); err != nil {
		return fmt.Errorf("error verifying SHA256 of %v: %v", archiveFile, err)
	}
	log.Printf("Unpacking %v ...", archiveFile)
	if err := unpackMicrosoftArchive(targetDir, archiveFile); err != nil {
		return fmt.Errorf("extracting archive %v: %v", archiveFile, err)
	}
	if err := os.WriteFile(filepath.Join(targetDir, unpackedOkay), nil, 0644); err != nil {
		return err
	}
	log.Printf("Success. You may now run '%v'", command)
	return nil
}

func unpackMicrosoftArchive(targetDir, archiveFile string) error {
	// Upstream validates tar entry names but not ZIP entry names. Check every
	// ZIP entry before extraction so an invalid archive cannot be partly unpacked.
	if strings.HasSuffix(archiveFile, ".zip") {
		zr, err := zip.OpenReader(archiveFile)
		if err != nil {
			return err
		}
		for _, f := range zr.File {
			if !validRelPath(f.Name) {
				zr.Close()
				return fmt.Errorf("zip file contained invalid name %q", f.Name)
			}
		}
		if err := zr.Close(); err != nil {
			return err
		}
	}
	return unpackArchive(targetDir, archiveFile)
}

func (archive MicrosoftArchive) validate(version string) error {
	if archive.URL == "" {
		return fmt.Errorf("no binary release of %v for %v/%v", version, getOS(), runtime.GOARCH)
	}
	if archive.SHA256 == "" {
		return fmt.Errorf("binary release of %v for %v/%v has no SHA256 checksum", version, getOS(), runtime.GOARCH)
	}
	return nil
}
