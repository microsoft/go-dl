// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.9-1 command runs the go command from Microsoft build of Go 1.23.9-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.9-1@latest
//	$ msgo1.23.9-1 download
//
// And then use the msgo1.23.9-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.9-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.9-1", "msgo1.23.9-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/ce33762b-ab2e-4d78-839d-8794468e9468/fbb2f3467f240cb7524afdda41c51877/go1.23.9-20250506.5.linux-amd64.tar.gz", SHA256: "22b9d1a5d4cab9068f424b3c7b3a8169165f8a995a44c5b30f29b77d1c24b560"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/ce33762b-ab2e-4d78-839d-8794468e9468/5af8133ac3a79d8773c669a7bb99d38b/go1.23.9-20250506.5.linux-armv6l.tar.gz", SHA256: "ea1a76b0ceaa0a874d8e50a7e3663c971e86b366113d9034dc8bd35c5f86abf6"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/ce33762b-ab2e-4d78-839d-8794468e9468/89b83471fd67a6c5135ebcdb164d3bdf/go1.23.9-20250506.5.linux-arm64.tar.gz", SHA256: "7e80b5b640a6bb14035c8c546ca35abae278b7b5f47b3c8bf30aacb965fc39ad"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/ce33762b-ab2e-4d78-839d-8794468e9468/9387c79905dd6f5da40cbdd94fdee082/go1.23.9-20250506.5.windows-amd64.zip", SHA256: "46b92142ef89b003957cc2c3bb360e84b021d1b6414fc0e7dc2d56700fc708c6"}
	default:
		return version.MicrosoftArchive{}
	}
}
