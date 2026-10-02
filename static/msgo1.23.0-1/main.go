// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.0-1 command runs the go command from Microsoft build of Go 1.23.0-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.0-1@latest
//	$ msgo1.23.0-1 download
//
// And then use the msgo1.23.0-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.0-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.0-1", "msgo1.23.0-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f348728d-8bff-41f6-af0a-6ed0138cd848/ceb5c69a57705dfed4fd61b552c39a5d/go1.23.0-20240814.2.linux-amd64.tar.gz", SHA256: "b7255b3d4daab96863f4d678db9b1415cb3a1abab2755b8bc630d16aba6c763b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f348728d-8bff-41f6-af0a-6ed0138cd848/1b9e15217b07da018fba99ece14fcffb/go1.23.0-20240814.2.linux-armv6l.tar.gz", SHA256: "1ff0837a61f094dabde817a1387b7c48b22673ab24e44d2dade3fd5e26324b7c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f348728d-8bff-41f6-af0a-6ed0138cd848/785ace49221185dcca49733d1872f62c/go1.23.0-20240814.2.linux-arm64.tar.gz", SHA256: "b3e5d1bc55b26c5bae66dc484577ba55c45dd1752699d0f8683a9b6eb982454b"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f348728d-8bff-41f6-af0a-6ed0138cd848/88681653ee6e9631a58900722086be0f/go1.23.0-20240814.2.windows-amd64.zip", SHA256: "5ade140045b3a002c6874748854b4427153c81ac1e1df0b60eaa4faaab0ee7a6"}
	default:
		return version.MicrosoftArchive{}
	}
}
