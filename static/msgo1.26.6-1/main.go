// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.6-1 command runs the go command from Microsoft build of Go 1.26.6-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.6-1@latest
//	$ msgo1.26.6-1 download
//
// And then use the msgo1.26.6-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.6-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.6-1", "msgo1.26.6-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4d610db-5cd8-4e26-a708-47e2bee4d7c7/bbc03c2a3d24b2097a6ed232f0fc34f9/go1.26.6-20260813.3.darwin-amd64.tar.gz", SHA256: "81e90c7778a11e2b231994532a6ed3416df61ada0aceef846958233c6369b179"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4d610db-5cd8-4e26-a708-47e2bee4d7c7/858bc7106af7c999b1597aea447f7a8e/go1.26.6-20260813.3.darwin-arm64.tar.gz", SHA256: "0e8b1f8a53fc50c228acb77bb2903e624c6e9a45366833eb5c3324f265fc47c5"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4d610db-5cd8-4e26-a708-47e2bee4d7c7/f9ce169e3b9aa315ca348f9d91036a13/go1.26.6-20260813.3.linux-amd64.tar.gz", SHA256: "d150752f0dfad988313e5dae993faac0f4adb3f253647714718383621a1f3362"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4d610db-5cd8-4e26-a708-47e2bee4d7c7/30d695864437f994859f9bd9eb889e60/go1.26.6-20260813.3.linux-armv6l.tar.gz", SHA256: "6eb17a73ea82fc0e1aec784ba60dc74d5efbbe535a9f6edfbf007a97a59d52c0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4d610db-5cd8-4e26-a708-47e2bee4d7c7/3794b1f427aa800dc686fd6ce03662b9/go1.26.6-20260813.3.linux-arm64.tar.gz", SHA256: "7ea6883405ac48b5fcf850f214fcd0482dc7f1b64eb68484e046c8417d6378fa"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4d610db-5cd8-4e26-a708-47e2bee4d7c7/294125ec2526dd8dd8bff7e6db84f155/go1.26.6-20260813.3.windows-amd64.zip", SHA256: "18dfcbb815383ef5533958a7d4c566301fe09067de7db6619ff90d7f0d78d332"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4d610db-5cd8-4e26-a708-47e2bee4d7c7/8b5bb3ac21e68e28ae2f1ccd127718bd/go1.26.6-20260813.3.windows-arm64.zip", SHA256: "4ebfc6c20cc91436b631e1e9c39499e0db0bb7e2261d29ee47b1e20b00f37d0a"}
	default:
		return version.MicrosoftArchive{}
	}
}
