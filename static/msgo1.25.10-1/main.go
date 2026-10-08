// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.10-1 command runs the go command from Microsoft build of Go 1.25.10-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.10-1@latest
//	$ msgo1.25.10-1 download
//
// And then use the msgo1.25.10-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.10-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.10-1", "msgo1.25.10-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/dbbb9a00-7906-4517-9b07-7674cb0f973c/12bde1a14b33ae283d90202549d0dba3/go1.25.10-20260507.2.darwin-amd64.tar.gz", SHA256: "d1a775a330bb62dc13d60234b6cff529ad8eb796b31d93649e061c149688b3b8"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/dbbb9a00-7906-4517-9b07-7674cb0f973c/a455e6a29edcb6cd740d88a33b6cfdcb/go1.25.10-20260507.2.darwin-arm64.tar.gz", SHA256: "3e441d81766ea70d6e4b76a5a9f8620d6bfa448c3aa7510bc9da62b611911cbd"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/dbbb9a00-7906-4517-9b07-7674cb0f973c/30f05072bd2afa6dbf07137de0ca3d7d/go1.25.10-20260507.2.linux-amd64.tar.gz", SHA256: "fd15a085c94974835913c2c5fa72375bf70a5bd541ad63cb1d11c2ad904ee243"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/dbbb9a00-7906-4517-9b07-7674cb0f973c/ffd0253be4b2ee398a700aac189db370/go1.25.10-20260507.2.linux-armv6l.tar.gz", SHA256: "38ace168481ad68fa38e2c9dab952e2b2e3b4882c5e4e7ffc972961852765b65"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/dbbb9a00-7906-4517-9b07-7674cb0f973c/22c49451b3b74d0358a765f06e70b4b8/go1.25.10-20260507.2.linux-arm64.tar.gz", SHA256: "64be28120ad89694a223f67a4badb4bd6f6ed6cfa4d156e3eaa047c6fb626c19"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/dbbb9a00-7906-4517-9b07-7674cb0f973c/14958e43cc573d9a0dc69e1a8e192e2a/go1.25.10-20260507.2.windows-amd64.zip", SHA256: "7210c1b8804a022fda0b391613f1429d82e946358dbd2ca17f4cf0fd7c27d837"}
	default:
		return version.MicrosoftArchive{}
	}
}
