// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.3-1 command runs the go command from Microsoft build of Go 1.26.3-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.3-1@latest
//	$ msgo1.26.3-1 download
//
// And then use the msgo1.26.3-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.3-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.3-1", "msgo1.26.3-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b5f4cfa-e528-4781-92cb-dbdb5effb00c/75619db85ab8b5467af699f2a5206013/go1.26.3-20260508.1.darwin-amd64.tar.gz", SHA256: "656c8680be43c2fa64ba1e044c446185264ba056191e62029d1d3e7ac88b0f00"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b5f4cfa-e528-4781-92cb-dbdb5effb00c/cf9e99b59a4cf1b7826ed21dcf627f4b/go1.26.3-20260508.1.darwin-arm64.tar.gz", SHA256: "48601476ef9883c84c82d05c13ebb50154825b1c0093186403c011f91eb96bf4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b5f4cfa-e528-4781-92cb-dbdb5effb00c/450d56d408a0b49c0d3a8ba15a0a810d/go1.26.3-20260508.1.linux-amd64.tar.gz", SHA256: "e09fedd6e53dfe835ea00b38bb678388965417ecdd28b02231fc868d4a9064c8"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b5f4cfa-e528-4781-92cb-dbdb5effb00c/f1ab658dc4451cca9405bb305e682444/go1.26.3-20260508.1.linux-armv6l.tar.gz", SHA256: "ff05b991a283f7a6debed1fe5643edf6a949b5f341cd4672cab1b2e352e764bd"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b5f4cfa-e528-4781-92cb-dbdb5effb00c/b9189bfbbd3c981361dc32f42bc57edb/go1.26.3-20260508.1.linux-arm64.tar.gz", SHA256: "782eb2cfeb1c2e21d604d810a3e81a1c7112ba1ca7b1b2aee6131ab0353c8278"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b5f4cfa-e528-4781-92cb-dbdb5effb00c/f48fd2bbefb9cfb4064e7fdca5a7e53a/go1.26.3-20260508.1.windows-amd64.zip", SHA256: "17691c07b572a77aae37bf681b49e77d514a7434e0f5a81e49a0671d7862a437"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b5f4cfa-e528-4781-92cb-dbdb5effb00c/7f0999be2275c62d9b2af2838a0e23d8/go1.26.3-20260508.1.windows-arm64.zip", SHA256: "271e7c1814c8d62d8e07c771d9991950375adda97bfd6a4256fe38b7527fb64b"}
	default:
		return version.MicrosoftArchive{}
	}
}
