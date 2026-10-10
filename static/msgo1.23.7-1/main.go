// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.7-1 command runs the go command from Microsoft build of Go 1.23.7-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.7-1@latest
//	$ msgo1.23.7-1 download
//
// And then use the msgo1.23.7-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.7-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.7-1", "msgo1.23.7-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/7aae5053-b83d-415d-a0a2-d1438effd694/f74ab547e22c6dcb50a6420a8d72baf8/go1.23.7-20250304.2.linux-amd64.tar.gz", SHA256: "aafcfa90c65ba5d44f38875801adf812f82f37d2ef5ea57ef9bb71cab3e3612e"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/7aae5053-b83d-415d-a0a2-d1438effd694/af742b58553c4de2b3f300334a72be8c/go1.23.7-20250304.2.linux-armv6l.tar.gz", SHA256: "841a8d78288f5681844a2b98786dc7fa0f33e7d700bbd6020230672cab3f9110"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/7aae5053-b83d-415d-a0a2-d1438effd694/10c8c667ee3ddd881bb6e2ad20312f30/go1.23.7-20250304.2.linux-arm64.tar.gz", SHA256: "edbbeb2bba93a0956ff238867da2864a71e0b3813211864c355b5966fa95ea67"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/7aae5053-b83d-415d-a0a2-d1438effd694/193bffc7137156d580f11f4d5e8500d1/go1.23.7-20250304.2.windows-amd64.zip", SHA256: "eaa90e5c158501611df23a1635eb46d81d2706a440351f48e2f523ca773bcd41"}
	default:
		return version.MicrosoftArchive{}
	}
}
