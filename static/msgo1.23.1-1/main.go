// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.1-1 command runs the go command from Microsoft build of Go 1.23.1-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.1-1@latest
//	$ msgo1.23.1-1 download
//
// And then use the msgo1.23.1-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.1-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.1-1", "msgo1.23.1-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e4a8eb45-a308-4caa-b3ee-26e0f51a0856/29c89b378fe41e988f8847e212b838d0/go1.23.1-20240905.5.linux-amd64.tar.gz", SHA256: "6ec113065364ca9cdd0e98b0aa86a127278ee8296d9c1b0605f79f563c2c2eb8"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e4a8eb45-a308-4caa-b3ee-26e0f51a0856/512f64e28aa5d7611e543e0f7208b033/go1.23.1-20240905.5.linux-armv6l.tar.gz", SHA256: "febe70f508757e37bb371637a554f6ca89f2d0c2b82a21bbf5cdf93da67bc2d4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e4a8eb45-a308-4caa-b3ee-26e0f51a0856/ee5f9f8e7863167d1d0539b322f201d4/go1.23.1-20240905.5.linux-arm64.tar.gz", SHA256: "811077d2e7fec5bd95982535470d4ddcc43116ff01d3e177cc3c844687a237e6"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e4a8eb45-a308-4caa-b3ee-26e0f51a0856/0d9c1bead9445238fccb962855b0b6e6/go1.23.1-20240905.5.windows-amd64.zip", SHA256: "b19cde66707199969e9ddf27430b00d89e36386aef039b38b07ba8fd55a1c072"}
	default:
		return version.MicrosoftArchive{}
	}
}
