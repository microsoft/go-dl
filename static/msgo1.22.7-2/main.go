// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.7-2 command runs the go command from Microsoft build of Go 1.22.7-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.7-2@latest
//	$ msgo1.22.7-2 download
//
// And then use the msgo1.22.7-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.7-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.7-2", "msgo1.22.7-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/57f3e2a8-0275-43ca-a7ce-8a24b0a74949/1dfe59d8fef2d45c6af9e76bf18e9362/go1.22.7-20240910.4.linux-amd64.tar.gz", SHA256: "8d5cb866f5428199acea05fd0332f69570e9f1feb7303615c121b8eb2b002a7f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/57f3e2a8-0275-43ca-a7ce-8a24b0a74949/b4180f3d81bd36e261eb684c6092e4d9/go1.22.7-20240910.4.linux-armv6l.tar.gz", SHA256: "00fd5d990f9d17027f68d8cae33941c1447f50f4215e49c14c486aba4c11339f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/57f3e2a8-0275-43ca-a7ce-8a24b0a74949/94ebaaa37b238a5c20841e66d202e2a0/go1.22.7-20240910.4.linux-arm64.tar.gz", SHA256: "adf0111c17c81f88e63fb6cdf9b8cde4019dbfbac7a01ed74ae71210059cb37b"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/57f3e2a8-0275-43ca-a7ce-8a24b0a74949/b0360dac3ad429b5f01c4d9267bd7c0b/go1.22.7-20240910.4.windows-amd64.zip", SHA256: "759e0753f2650982c0351afb14b0e83a38e85cb23324f7f6e3325322b7d62f7a"}
	default:
		return version.MicrosoftArchive{}
	}
}
