// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestUpstreamFilesMatchChecksums(t *testing.T) {
	checksums, err := readUpstreamChecksums("upstream-files.sha256")
	if err != nil {
		t.Fatal(err)
	}
	if len(checksums) == 0 {
		t.Fatal("upstream checksum file contains no files")
	}
	for name, want := range checksums {
		data, err := os.ReadFile(name)
		if err != nil {
			t.Fatal(err)
		}
		got := sha256.Sum256(data)
		if got != want {
			t.Errorf("%s SHA-256 = %x, want %x; run go run ./cmd/fetch-upstream from internal/gen after reviewing upstream changes", name, got, want)
		}
	}
}

func readUpstreamChecksums(name string) (map[string][sha256.Size]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	checksums := make(map[string][sha256.Size]byte)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid checksum line %q", line)
		}
		digest, err := hex.DecodeString(fields[0])
		if err != nil || len(digest) != sha256.Size {
			return nil, fmt.Errorf("invalid SHA-256 in line %q", line)
		}
		var checksum [sha256.Size]byte
		copy(checksum[:], digest)
		if _, exists := checksums[fields[1]]; exists {
			return nil, fmt.Errorf("duplicate checksum for %q", fields[1])
		}
		checksums[fields[1]] = checksum
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return checksums, nil
}
