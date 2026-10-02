// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/microsoft/go-infra/assets"
	"golang.org/x/sync/errgroup"
)

const assetURLFormat = "https://aka.ms/golang/release/latest/go1.%d.%d-%d.assets.json"

var errVersionNotFound = errors.New("version not found")

// Series is the Go release series after "1.", while Patch and Revision are the
// upstream patch release and Microsoft revision.
type version struct {
	Series   int
	Patch    int
	Revision int
}

func (v version) String() string {
	return fmt.Sprintf("1.%d.%d-%d", v.Series, v.Patch, v.Revision)
}

type scraper struct {
	client    *http.Client
	dbRoot    string
	urlFormat string
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	if _, err := os.Stat(filepath.Join("db", "db.go")); err != nil {
		return fmt.Errorf("expected to run from the internal/gen module: %w", err)
	}
	s := scraper{
		client: &http.Client{
			Timeout: time.Minute,
		},
		dbRoot:    "db",
		urlFormat: assetURLFormat,
	}
	return s.scrape(context.Background())
}

func (s scraper) scrape(ctx context.Context) error {
	group, groupCtx := errgroup.WithContext(ctx)

	var mu sync.Mutex
	checked := make(map[version]struct{})

	var check func(version)
	check = func(v version) {
		mu.Lock()
		if _, ok := checked[v]; ok {
			mu.Unlock()
			return
		}
		checked[v] = struct{}{}
		mu.Unlock()

		group.Go(func() error {
			data, err := s.download(groupCtx, v)
			if errors.Is(err, errVersionNotFound) {
				log.Printf("No manifest found for %s", v)
				return nil
			}
			if err != nil {
				return err
			}
			if err := s.writeManifest(v, data); err != nil {
				return err
			}

			// Every successful probe can reveal a subsequent Microsoft
			// revision, Go patch release, or Go release series.
			check(version{Series: v.Series, Patch: v.Patch, Revision: v.Revision + 1})
			check(version{Series: v.Series, Patch: v.Patch + 1, Revision: 1})
			check(version{Series: v.Series + 1, Patch: 0, Revision: 1})
			return nil
		})
	}

	check(version{Series: 22, Patch: 0, Revision: 1})
	if err := group.Wait(); err != nil {
		return fmt.Errorf("scrape build asset database: %w", err)
	}
	return nil
}

func (s scraper) download(ctx context.Context, v version) ([]byte, error) {
	url := fmt.Sprintf(s.urlFormat, v.Series, v.Patch, v.Revision)
	log.Printf("Downloading %s", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request for version %s: %w", v, err)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch version %s: %w", v, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch version %s: unexpected status %s", v, resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read version %s: %w", v, err)
	}
	trimmed := bytes.TrimSpace(data)
	// aka.ms returns an HTTP 200 Bing page for a link that doesn't exist.
	if len(trimmed) == 0 || trimmed[0] == '<' {
		return nil, fmt.Errorf("%s: %w", v, errVersionNotFound)
	}

	var manifest assets.ToolsetBuild
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("decode version %s: %w", v, err)
	}
	if manifest.Version != v.String() {
		return nil, fmt.Errorf("manifest for version %s contains version %q", v, manifest.Version)
	}
	return data, nil
}

func (s scraper) writeManifest(v version, data []byte) error {
	dir := filepath.Join(s.dbRoot, fmt.Sprintf("go1.%d", v.Series))
	filename := fmt.Sprintf("go%s.assets.json", v)
	target := filepath.Join(dir, filename)

	existing, err := os.ReadFile(target)
	if err == nil {
		// Versioned manifests are immutable because generated wrappers pin
		// their exact bytes. Never silently rewrite one.
		if !bytes.Equal(existing, data) {
			return fmt.Errorf("manifest %q differs from the published manifest", target)
		}
		log.Printf("Unchanged %s", target)
		return nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("read manifest %q: %w", target, err)
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create manifest directory %q: %w", dir, err)
	}
	temp, err := os.CreateTemp(dir, "."+filename+".tmp-*")
	if err != nil {
		return fmt.Errorf("create temporary manifest for %q: %w", target, err)
	}
	tempName := temp.Name()
	defer os.Remove(tempName)

	if _, err := temp.Write(data); err != nil {
		temp.Close()
		return fmt.Errorf("write temporary manifest for %q: %w", target, err)
	}
	if err := temp.Chmod(0o644); err != nil {
		temp.Close()
		return fmt.Errorf("set permissions on temporary manifest for %q: %w", target, err)
	}
	if err := temp.Close(); err != nil {
		return fmt.Errorf("close temporary manifest for %q: %w", target, err)
	}
	if err := os.Rename(tempName, target); err != nil {
		return fmt.Errorf("install manifest %q: %w", target, err)
	}
	log.Printf("Added %s", target)
	return nil
}
