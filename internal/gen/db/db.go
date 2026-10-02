// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package db provides the build asset manifests used to generate go-dl packages.
package db

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"regexp"
	"sort"
	"strconv"

	"github.com/microsoft/go-infra/assets"
	"github.com/microsoft/go-infra/goversion"
)

//go:embed go1.*/*.assets.json
var files embed.FS

var manifestPathPattern = regexp.MustCompile(`^go(1\.\d+)/go(1\.\d+\.\d+-\d+)\.assets\.json$`)

// Entry contains one embedded build asset manifest and its parsed metadata.
type Entry struct {
	Version  *goversion.GoVersion
	Manifest assets.ToolsetBuild
}

// Load reads, validates, and sorts all embedded build asset manifests.
func Load() ([]Entry, error) {
	return load(files)
}

func load(fsys fs.FS) ([]Entry, error) {
	var entries []Entry
	seen := make(map[string]string)
	err := fs.WalkDir(fsys, ".", func(name string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() || path.Ext(name) != ".json" {
			return nil
		}

		version, err := parseManifestPath(name)
		if err != nil {
			return err
		}
		if previous, ok := seen[version.Full()]; ok {
			return fmt.Errorf("build asset manifests %q and %q both identify version %s", previous, name, version.Full())
		}

		data, err := fs.ReadFile(fsys, name)
		if err != nil {
			return fmt.Errorf("read build asset manifest %q: %w", name, err)
		}
		var manifest assets.ToolsetBuild
		if err := json.Unmarshal(data, &manifest); err != nil {
			return fmt.Errorf("decode build asset manifest %q: %w", name, err)
		}
		if manifest.Version != version.Full() {
			return fmt.Errorf("build asset manifest %q contains version %q, want %q", name, manifest.Version, version.Full())
		}

		seen[version.Full()] = name
		entries = append(entries, Entry{
			Version:  version,
			Manifest: manifest,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("build asset database is empty")
	}

	sort.Slice(entries, func(i, j int) bool {
		return versionLess(entries[i].Version, entries[j].Version)
	})
	return entries, nil
}

func parseManifestPath(name string) (*goversion.GoVersion, error) {
	match := manifestPathPattern.FindStringSubmatch(name)
	if match == nil {
		return nil, fmt.Errorf("build asset manifest path %q does not match go1.x/go1.x.y-z.assets.json", name)
	}
	version := goversion.New(match[2])
	if version.MajorMinor() != match[1] || version.Full() != match[2] {
		return nil, fmt.Errorf("build asset manifest path %q has inconsistent or noncanonical versions", name)
	}
	return version, nil
}

func versionLess(a, b *goversion.GoVersion) bool {
	for _, parts := range [][2]string{
		{a.Major, b.Major},
		{a.Minor, b.Minor},
		{a.Patch, b.Patch},
		{a.Revision, b.Revision},
	} {
		aPart, _ := strconv.Atoi(parts[0])
		bPart, _ := strconv.Atoi(parts[1])
		if aPart != bPart {
			return aPart < bPart
		}
	}
	return false
}
