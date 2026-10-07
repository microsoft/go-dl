// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.4-1 command runs the go command from Microsoft build of Go 1.23.4-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.4-1@latest
//	$ msgo1.23.4-1 download
//
// And then use the msgo1.23.4-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.4-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.4-1", "msgo1.23.4-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c93d49fc-cbd4-4e48-912a-0dcaeda0d8ce/3406ebb6e1f8fa2cea8056839cc5e67c/go1.23.4-20241203.2.linux-amd64.tar.gz", SHA256: "b428b028b482c76b4db0b75e5904a0f3efd6b432db7e7717cc1a25c53eb5ab56"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c93d49fc-cbd4-4e48-912a-0dcaeda0d8ce/912668e3f5b7711cd3f60dbc9fcb295c/go1.23.4-20241203.2.linux-armv6l.tar.gz", SHA256: "d418d81cf9ff989aea03fcd047a0c0f7e97303710eb810b9a1a304d715ccb116"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c93d49fc-cbd4-4e48-912a-0dcaeda0d8ce/7c157b265b957d6021162320e0f1b736/go1.23.4-20241203.2.linux-arm64.tar.gz", SHA256: "3164753c517442977e49532b92f657c6c93323025bb99b63c2b14963dbba8115"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c93d49fc-cbd4-4e48-912a-0dcaeda0d8ce/c9426c644f61c72f38994a16580ac9b9/go1.23.4-20241203.2.windows-amd64.zip", SHA256: "0c95f2d9bfb02f7a3bff5b4f23f0373bf6f8d9e8d5a27d3d6a03fff1c1881517"}
	default:
		return version.MicrosoftArchive{}
	}
}
