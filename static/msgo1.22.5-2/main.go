// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.5-2 command runs the go command from Microsoft build of Go 1.22.5-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.5-2@latest
//	$ msgo1.22.5-2 download
//
// And then use the msgo1.22.5-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.5-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.5-2", "msgo1.22.5-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/724b7dab-d078-4dec-91e1-1bd0b44ef633/06dceb5a69d1318939c8c466d496dbe0/go1.22.5-20240722.5.linux-amd64.tar.gz", SHA256: "59963327dc86e8a5f48d9485ed800c26f80b9975c9331736c33005947a02686f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/724b7dab-d078-4dec-91e1-1bd0b44ef633/254bc806f38765e6d02981f1481c4aa9/go1.22.5-20240722.5.linux-armv6l.tar.gz", SHA256: "c641ce883b1967fb5df679ee652ebd2482dcbc14e528b36fea942ab69b4ca2f2"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/724b7dab-d078-4dec-91e1-1bd0b44ef633/eeb1608811d4a63e40bdf4cbb0a5d76f/go1.22.5-20240722.5.linux-arm64.tar.gz", SHA256: "92078bdb0a2cbe1a5c175b3046427eca30f137b57106157d0696099254bbe03d"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/724b7dab-d078-4dec-91e1-1bd0b44ef633/2d8f121a1e9df25420d7f0b4bd70c80a/go1.22.5-20240722.5.windows-amd64.zip", SHA256: "d7b897d0c594e9a0cdc115353c337f10090bcf716138b50cee99aaa84143688d"}
	default:
		return version.MicrosoftArchive{}
	}
}
