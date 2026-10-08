// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.1-1 command runs the go command from Microsoft build of Go 1.22.1-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.1-1@latest
//	$ msgo1.22.1-1 download
//
// And then use the msgo1.22.1-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.1-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.1-1", "msgo1.22.1-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/010ede32-b10a-456f-8111-2194f1dec0fe/b4bbc37bcb1b4b0262235e72547ad687/go1.22.1-20240305.9.linux-amd64.tar.gz", SHA256: "49b9ca5b5b1ca9211b58a14507c73d3d68d571b8c40dd5c0d5bc128f8048ebfa"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/010ede32-b10a-456f-8111-2194f1dec0fe/cbb62e7422878240a7806c9f5be67ecc/go1.22.1-20240305.9.linux-armv6l.tar.gz", SHA256: "57ad8864893c8bce729d5830241347c8556e2eafe7b528fad24df2196d33c6b8"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/010ede32-b10a-456f-8111-2194f1dec0fe/41eae1bfca1fc1722c5e18f0426737e2/go1.22.1-20240305.9.linux-arm64.tar.gz", SHA256: "9ef1ca23ee6c768fba8047cf1192911126bd0b6de12525911d3f839babcdc097"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/010ede32-b10a-456f-8111-2194f1dec0fe/b5d4855fe86591eef9091b1ed866177f/go1.22.1-20240305.9.windows-amd64.zip", SHA256: "3e2136ee2c76c2f5fe6abfc09d34c65be806aca95d858656054c6dcf64414ac3"}
	default:
		return version.MicrosoftArchive{}
	}
}
