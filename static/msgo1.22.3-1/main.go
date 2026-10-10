// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.3-1 command runs the go command from Microsoft build of Go 1.22.3-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.3-1@latest
//	$ msgo1.22.3-1 download
//
// And then use the msgo1.22.3-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.3-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.3-1", "msgo1.22.3-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0ed7281-aac5-4709-b0c2-5b42700fb055/579e563d275a15708a429abe1edf1d4d/go1.22.3-20240507.3.linux-amd64.tar.gz", SHA256: "dbfa3440c927f2b780556737e0991ce038caafc964b391c362b46a101ac48792"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0ed7281-aac5-4709-b0c2-5b42700fb055/d1f255a5bb056eafaeebc16088453841/go1.22.3-20240507.3.linux-armv6l.tar.gz", SHA256: "a2d62080c10bcd36af8e47e36c5a811ba98bccb91555ed05680347acb7c7f3ab"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0ed7281-aac5-4709-b0c2-5b42700fb055/96566490195026681ab7976ced7d3eb4/go1.22.3-20240507.3.linux-arm64.tar.gz", SHA256: "15ab54184a844101fb146e4bc21129fdbf88c134b1bb1faa92bffcc950737b80"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0ed7281-aac5-4709-b0c2-5b42700fb055/7ffe54674cebe079c4a608da91d23e18/go1.22.3-20240507.3.windows-amd64.zip", SHA256: "759117af3df20ff528176ac747bc083dd376f1762e9c90f847ba1918db0ee7b9"}
	default:
		return version.MicrosoftArchive{}
	}
}
