// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.10-1 command runs the go command from Microsoft build of Go 1.22.10-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.10-1@latest
//	$ msgo1.22.10-1 download
//
// And then use the msgo1.22.10-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.10-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.10-1", "msgo1.22.10-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/549b0e83-52e7-4b33-bb9f-69335b7c670a/0ca83fa95ffb1500238afa2ff0a026e7/go1.22.10-20241203.4.linux-amd64.tar.gz", SHA256: "4cee97a45df75e231d8bf04ab80547920b0bd830ed5adfbc01b525c83354cc87"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/549b0e83-52e7-4b33-bb9f-69335b7c670a/d8355b08b6d34b3b0f09383d6b9c36dd/go1.22.10-20241203.4.linux-armv6l.tar.gz", SHA256: "704a3388458b9a35f878620c9d0b76b7433428eead41f6ee0e6bcb6994c28dbb"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/549b0e83-52e7-4b33-bb9f-69335b7c670a/4ecb0268aed2bf966a3592c24a745cf4/go1.22.10-20241203.4.linux-arm64.tar.gz", SHA256: "44016e44de306663054762782d332a8796285593f437981a8e612bb27be15910"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/549b0e83-52e7-4b33-bb9f-69335b7c670a/d0fb21520ee8650e724027eb424af109/go1.22.10-20241203.4.windows-amd64.zip", SHA256: "4a2c6e697dac2d38a1a08470c31d350709de131a7b4a1edb14a4b46bcbadf862"}
	default:
		return version.MicrosoftArchive{}
	}
}
