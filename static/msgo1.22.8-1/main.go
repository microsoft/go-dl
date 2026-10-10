// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.8-1 command runs the go command from Microsoft build of Go 1.22.8-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.8-1@latest
//	$ msgo1.22.8-1 download
//
// And then use the msgo1.22.8-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.8-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.8-1", "msgo1.22.8-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/5eb583bd-ddfb-4f9e-846f-fd78d9cccca2/3bad0f443f15b53c30ad07c755a403e8/go1.22.8-20241001.6.linux-amd64.tar.gz", SHA256: "f3f812fe9ce2e6f7ba14a38ad29b1e1907153bfb5589967086cb3ce367c13565"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/5eb583bd-ddfb-4f9e-846f-fd78d9cccca2/789ea17076a107f2c1c2c6b715a1e8a0/go1.22.8-20241001.6.linux-armv6l.tar.gz", SHA256: "047c824a677a35e0e9ae57eb8bc9a76f0be4a507e0d465f2167bf96a1a7a104b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/5eb583bd-ddfb-4f9e-846f-fd78d9cccca2/c93c7c16f780ad2903f38724b4a4f454/go1.22.8-20241001.6.linux-arm64.tar.gz", SHA256: "9d77b18a2b1413ce41d925465d8132a982368a9d937252880d298cdb49978b3a"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/5eb583bd-ddfb-4f9e-846f-fd78d9cccca2/8de78384ed7d75807373814be4e6f9e3/go1.22.8-20241001.6.windows-amd64.zip", SHA256: "fe045e6b19f1c513269200fea25d9ecc70c25276ce5701229fe58db2d06de6f2"}
	default:
		return version.MicrosoftArchive{}
	}
}
