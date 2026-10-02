// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.1-1 command runs the go command from Microsoft build of Go 1.26.1-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.1-1@latest
//	$ msgo1.26.1-1 download
//
// And then use the msgo1.26.1-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.1-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.1-1", "msgo1.26.1-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f5319d7f-0a67-4e49-9fac-dd1d041c3a6c/3fa28531a3789b018eedc004abc50558/go1.26.1-20260306.1.darwin-amd64.tar.gz", SHA256: "166e95b9c1508b4f6a44ba303b1b4e97351e6b29df212bb9fa718f0eaa25a7d8"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f5319d7f-0a67-4e49-9fac-dd1d041c3a6c/cab22390f9681deb3d8c0eeac8d18aa4/go1.26.1-20260306.1.darwin-arm64.tar.gz", SHA256: "180bc6df45693c99e48017f3a6c5a351903ea4064a280671dc85e383b4e26850"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f5319d7f-0a67-4e49-9fac-dd1d041c3a6c/7290e09f7c945f2130d9e5a9e7682267/go1.26.1-20260306.1.linux-amd64.tar.gz", SHA256: "a16f9f7dfc8e6e25f6bac2db644d5d87c0f3f0c1bcb1213c6e20c05f452b28d5"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f5319d7f-0a67-4e49-9fac-dd1d041c3a6c/3c31cacec77843fc3bdc29b758c97b69/go1.26.1-20260306.1.linux-armv6l.tar.gz", SHA256: "931345dae1f45eb025de20b9fd3d63f5f0cb056fe5f5aaa73b651476109a67c9"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f5319d7f-0a67-4e49-9fac-dd1d041c3a6c/9837189ebe0dcd55014cc7df0ce303fc/go1.26.1-20260306.1.linux-arm64.tar.gz", SHA256: "4002724aa3f71d88653c0ce2e8860e013e3ab1f9b6e57d1cdedeb519132e8c12"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f5319d7f-0a67-4e49-9fac-dd1d041c3a6c/282a3755b7c0a90de86f240c63bf2a36/go1.26.1-20260306.1.windows-amd64.zip", SHA256: "1f7421dd74845777b2ad24b2a45beebbc407acd04932458394f0678a26f81407"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f5319d7f-0a67-4e49-9fac-dd1d041c3a6c/e1ad044f9b9c8587a2c12d2f3dc248aa/go1.26.1-20260306.1.windows-arm64.zip", SHA256: "6d64b1aec4f36d0f858892b92fa0314b9af87eaecc2ad71cfc3b0fd7ed81114b"}
	default:
		return version.MicrosoftArchive{}
	}
}
