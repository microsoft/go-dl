// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.0-1 command runs the go command from Microsoft build of Go 1.22.0-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.0-1@latest
//	$ msgo1.22.0-1 download
//
// And then use the msgo1.22.0-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.0-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.0-1", "msgo1.22.0-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/5d856430-f43c-4203-8a7a-e4668fabe2de/25d0add9fcc5b4d508069d6f489e75a3/go1.22.0-20240208.1.linux-amd64.tar.gz", SHA256: "a9172d8103f9d4bc08e00ffd14929d09a7d4c25f61c81fa768372f49f0258c2c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/5d856430-f43c-4203-8a7a-e4668fabe2de/58d81d39146e57ff5077cbc179077d6a/go1.22.0-20240208.1.linux-armv6l.tar.gz", SHA256: "42721a8ad29b95831d2d03a11c575650cc0dc2ed6928cd9585d466cc2681fcf1"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/5d856430-f43c-4203-8a7a-e4668fabe2de/edfa75dd7724426f4ecd885498960b52/go1.22.0-20240208.1.linux-arm64.tar.gz", SHA256: "ff50182dceee5c8bc3b36efd4c73d185b36e4a042dd6624620a3cf53f2c516a2"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/5d856430-f43c-4203-8a7a-e4668fabe2de/703dbd298f65dcc53611d5f90f376260/go1.22.0-20240208.1.windows-amd64.zip", SHA256: "d52c9144af985c7484401b9cede0294b167bff5916aa5962e886f747a0b2bd99"}
	default:
		return version.MicrosoftArchive{}
	}
}
