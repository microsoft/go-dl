// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.11-1 command runs the go command from Microsoft build of Go 1.24.11-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.11-1@latest
//	$ msgo1.24.11-1 download
//
// And then use the msgo1.24.11-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.11-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.11-1", "msgo1.24.11-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3b8b6a35-ce1d-4b9d-8769-f15c9db0c583/d3effe7909ac540102081365fc659f4e/go1.24.11-20251202.3.darwin-amd64.tar.gz", SHA256: "e19d1c55367565b5b7d90883fc5a24c4805d024350bff5dc34e6533ea6f906fb"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3b8b6a35-ce1d-4b9d-8769-f15c9db0c583/934d50deeacaa03a7765fb783cc9b7d4/go1.24.11-20251202.3.darwin-arm64.tar.gz", SHA256: "40ee2f4e95420c377d41f0e28219b87fd53bd666895eea75ec27b2998fdd9848"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3b8b6a35-ce1d-4b9d-8769-f15c9db0c583/41f7246beebc77837298bdacb1d333a5/go1.24.11-20251202.3.linux-amd64.tar.gz", SHA256: "863b16abbb23c88023dcbee28754d316352064af4083885afe3a23c43720818b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3b8b6a35-ce1d-4b9d-8769-f15c9db0c583/9ed860964127278d3ca5249381bc96f7/go1.24.11-20251202.3.linux-armv6l.tar.gz", SHA256: "2548f16f38e2a85c6a1ff29f5a5bbc1077f7daaafd0c117ea13ed337275ae568"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3b8b6a35-ce1d-4b9d-8769-f15c9db0c583/be60270b9261e9d11f6418e0e867a8c7/go1.24.11-20251202.3.linux-arm64.tar.gz", SHA256: "0b4b0fd6e74f02e9fc44ae6f1125dfedd668cf72d9cd06ffd7781be43bd13af5"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3b8b6a35-ce1d-4b9d-8769-f15c9db0c583/31b952572863e469ee1c1f897cc9d53c/go1.24.11-20251202.3.windows-amd64.zip", SHA256: "00b4f189001a72a8710ac276384d8a8c5897d7eb867739f8c517d97a0dfde97a"}
	default:
		return version.MicrosoftArchive{}
	}
}
