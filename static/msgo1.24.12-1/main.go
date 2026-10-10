// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.12-1 command runs the go command from Microsoft build of Go 1.24.12-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.12-1@latest
//	$ msgo1.24.12-1 download
//
// And then use the msgo1.24.12-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.12-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.12-1", "msgo1.24.12-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/05f68804-c1ba-4142-9c57-ae52bbddfd89/7878f248a53f88dbce60dfb857ad264a/go1.24.12-20260116.10.darwin-amd64.tar.gz", SHA256: "312c79aa103e7e333dde36beb1cbf3de032f43abd6e4fd7d7cac7f48fd7a34a9"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/05f68804-c1ba-4142-9c57-ae52bbddfd89/60b3081237ccaa500a94cd574a51be05/go1.24.12-20260116.10.darwin-arm64.tar.gz", SHA256: "94b2c51ac791462e799977f68c558edc5a36eced4360255df0b4e2ba4135c66d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/05f68804-c1ba-4142-9c57-ae52bbddfd89/63a1f776f7d0501d066dde8caadf1002/go1.24.12-20260116.10.linux-amd64.tar.gz", SHA256: "166247a9dd0fc726a819ba0e11796da9b5b3b03082011361835857a782894d23"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/05f68804-c1ba-4142-9c57-ae52bbddfd89/0d1261d71f2706a76e686d2754103f8a/go1.24.12-20260116.10.linux-armv6l.tar.gz", SHA256: "15efd07bba7279f0095c28df9f06474b373f59a791a8e6896405c3e4147eda8c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/05f68804-c1ba-4142-9c57-ae52bbddfd89/691e88c1d18366425bd01c87f110e94e/go1.24.12-20260116.10.linux-arm64.tar.gz", SHA256: "5e7949083542b6cbecd3cf3b5e6580415418965424569fbda144a2bad729f755"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/05f68804-c1ba-4142-9c57-ae52bbddfd89/2081f7398dddf5a7eb696e52a198c7cf/go1.24.12-20260116.10.windows-amd64.zip", SHA256: "1b30fe3ed48054a62301a587f9add8bbc304dd37129ed92f44bce86acafc313e"}
	default:
		return version.MicrosoftArchive{}
	}
}
