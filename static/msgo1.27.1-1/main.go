// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.27.1-1 command runs the go command from Microsoft build of Go 1.27.1-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.27.1-1@latest
//	$ msgo1.27.1-1 download
//
// And then use the msgo1.27.1-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.27.1-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.27.1-1", "msgo1.27.1-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3385953-fcfb-4ab3-a2c7-9fee00b1a099/27d41f9c8bbadebe6d8bd92d8f5ec498/go1.27.1-20260901.1.darwin-amd64.tar.gz", SHA256: "bdee35be74d7d705077760f3f1d0245dacfceb990b162e4f0bf1ecdd874a58d1"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3385953-fcfb-4ab3-a2c7-9fee00b1a099/9ddc91cc34ca9a32a96aa35ea5bd97ee/go1.27.1-20260901.1.darwin-arm64.tar.gz", SHA256: "1f777b34d38f3136582c7e8453e53e7dc0564a71d54167d94777ebcd6d4eeb2d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3385953-fcfb-4ab3-a2c7-9fee00b1a099/b11c6e30cdc20c35e091cf402cc3a376/go1.27.1-20260901.1.linux-amd64.tar.gz", SHA256: "3bb61524a969db8116f2a65ba0650d7d394caadd5d6f9a2fc0a4a56714b18ed2"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3385953-fcfb-4ab3-a2c7-9fee00b1a099/ac92e12dd224618c246d3633d47d9258/go1.27.1-20260901.1.linux-armv6l.tar.gz", SHA256: "bf03e7f93d6d2eada76e3d5813f94603737a8d7e76de79732ecf13c06f56d191"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3385953-fcfb-4ab3-a2c7-9fee00b1a099/3cd1de54c241e1142e6abe80e545be4f/go1.27.1-20260901.1.linux-arm64.tar.gz", SHA256: "4eea08235a737dbd7edc33d0fc3001402e1edc531da6b0b736589e38af30f88c"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3385953-fcfb-4ab3-a2c7-9fee00b1a099/bdb0af7b42ea960f8d131119cefcc6ed/go1.27.1-20260901.1.windows-amd64.zip", SHA256: "ba24d6f4cc6c6d57b47b34e01810dd74d37dd6d3a58e9dff0839dcc329dec45c"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3385953-fcfb-4ab3-a2c7-9fee00b1a099/ee45262b8dee19a3bda1fbff3247fc96/go1.27.1-20260901.1.windows-arm64.zip", SHA256: "7d212da5f4da1d69b7bc3334b6db7bc5c581fdd9df28f8f2fec4a5afcb816ee2"}
	default:
		return version.MicrosoftArchive{}
	}
}
