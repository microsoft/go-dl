// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.6-1 command runs the go command from Microsoft build of Go 1.25.6-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.6-1@latest
//	$ msgo1.25.6-1 download
//
// And then use the msgo1.25.6-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.6-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.6-1", "msgo1.25.6-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/25b7e44d-d224-4680-9b8a-820242e2854c/3e08bf081d51d7cc5a6820f80d8790f5/go1.25.6-20260116.7.darwin-amd64.tar.gz", SHA256: "8fc9fecb5b52d8cf902356f7c4c8a45331dbbe2938650aba776aa2c511d5040c"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/25b7e44d-d224-4680-9b8a-820242e2854c/21bb083304f8f726f40a9979d9890bc5/go1.25.6-20260116.7.darwin-arm64.tar.gz", SHA256: "10ff8debef44ea56fd8103b8a89a196171016df0a4698f11320757fa9c7dd046"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/25b7e44d-d224-4680-9b8a-820242e2854c/4d7bff6e6044a01a9b5478395e9e8acb/go1.25.6-20260116.7.linux-amd64.tar.gz", SHA256: "0b93d9d645ea83f70a2cdd605b2189f4abf43650b278cedd654389cb319f4c4c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/25b7e44d-d224-4680-9b8a-820242e2854c/f231b4e184df0d51cc910abc2dc493be/go1.25.6-20260116.7.linux-armv6l.tar.gz", SHA256: "b1950cd3cd046c932db421deb72a4512d8ec2555a0d72a6d6118889a996ead7d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/25b7e44d-d224-4680-9b8a-820242e2854c/87319434ff811c7eb60ae1bae8d87e8b/go1.25.6-20260116.7.linux-arm64.tar.gz", SHA256: "949471b6fe51e9c5437dc7b547006ee1b73a97717def54474d2d9a034303bcdd"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/25b7e44d-d224-4680-9b8a-820242e2854c/418c7ccca44ca30ae4fc991966021a88/go1.25.6-20260116.7.windows-amd64.zip", SHA256: "6bf7431f1b383450bcaf0a834433b733b680df101ea9e4e24d1f5abd275c1813"}
	default:
		return version.MicrosoftArchive{}
	}
}
