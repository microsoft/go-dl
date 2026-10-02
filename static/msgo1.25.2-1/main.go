// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.2-1 command runs the go command from Microsoft build of Go 1.25.2-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.2-1@latest
//	$ msgo1.25.2-1 download
//
// And then use the msgo1.25.2-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.2-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.2-1", "msgo1.25.2-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1c274553-d506-4db3-a558-600f256e580f/3af1011b0a3b0f22d6fd4921fcc1c010/go1.25.2-20251008.3.darwin-amd64.tar.gz", SHA256: "3c2dda8b0a294e6232fab65b42da9b46057f211e0fd7f21fc4644c0b5b38d1a1"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1c274553-d506-4db3-a558-600f256e580f/9003a8486c27ee7ece63b592ea145ca0/go1.25.2-20251008.3.darwin-arm64.tar.gz", SHA256: "5b8a048f6a652a1777741a047acd833e52b7065615ae9c734273d01109423661"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1c274553-d506-4db3-a558-600f256e580f/8f81512d8476c405511fbc353cc10ece/go1.25.2-20251008.3.linux-amd64.tar.gz", SHA256: "42746b2774e8c142a1d19699e56cbaef373beeab0f9110c2a39bf22ce71ff232"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1c274553-d506-4db3-a558-600f256e580f/4bde2f5b4d77f286514e87d4ece3b893/go1.25.2-20251008.3.linux-armv6l.tar.gz", SHA256: "23930438de3cd06363efd68b4a43033f11299ede796e2d3a4a5338ad55bbc534"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1c274553-d506-4db3-a558-600f256e580f/b655db36010d0dc2367b302c9c909351/go1.25.2-20251008.3.linux-arm64.tar.gz", SHA256: "9740a895843b92ce4c1e3d559dbeea265541efc0fda4d14e7ca995caefd7b64f"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1c274553-d506-4db3-a558-600f256e580f/137648ecb162b96844b20b7dc366324e/go1.25.2-20251008.3.windows-amd64.zip", SHA256: "042e90a1d02989c79e50cfc8bfe20bc2a365b532a6319758d4c64f103f3dab94"}
	default:
		return version.MicrosoftArchive{}
	}
}
