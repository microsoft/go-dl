// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/microsoft/go-dl/internal/gen/db"
	"github.com/microsoft/go-dl/internal/gen/mkdl"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	entries, err := db.Load()
	if err != nil {
		return err
	}
	return mkdl.Regenerate(entries)
}
