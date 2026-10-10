// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.3-1 command runs the go command from Microsoft build of Go 1.23.3-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.3-1@latest
//	$ msgo1.23.3-1 download
//
// And then use the msgo1.23.3-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.3-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.3-1", "msgo1.23.3-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/221ef1d1-2ed0-4d66-99fc-8d9673448088/b7cad2762bf5b024385854a488fe97c8/go1.23.3-20241107.5.linux-amd64.tar.gz", SHA256: "63f60378974040471ed951d50ef36f088282ca09d65d62fa014287147ea4d4bc"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/221ef1d1-2ed0-4d66-99fc-8d9673448088/b090d7aafa21d63ed8b6046dd52f1ca0/go1.23.3-20241107.5.linux-armv6l.tar.gz", SHA256: "2b65688cdea4b29afc9ac05ce33a09b2e6cbf0e941d2b327fd35a4902b80e879"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/221ef1d1-2ed0-4d66-99fc-8d9673448088/7ace59654ff2932a466c95342899e9ea/go1.23.3-20241107.5.linux-arm64.tar.gz", SHA256: "f51c8b8d7deb4dc6a7750e326fce940840afd822738c1f40f01e763e18e0f3cf"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/221ef1d1-2ed0-4d66-99fc-8d9673448088/78d060820c5ac7d5fc94e425397a07fd/go1.23.3-20241107.5.windows-amd64.zip", SHA256: "d21f4f149b961e5ee39bbfae1a0e34ba14982d1e4c86fe61808419c5579e9cbc"}
	default:
		return version.MicrosoftArchive{}
	}
}
