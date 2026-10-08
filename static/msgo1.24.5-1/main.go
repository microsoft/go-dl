// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.5-1 command runs the go command from Microsoft build of Go 1.24.5-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.5-1@latest
//	$ msgo1.24.5-1 download
//
// And then use the msgo1.24.5-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.5-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.5-1", "msgo1.24.5-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4ba02eab-f12f-46f0-8d2e-561a36985ed3/786fc7d13212d2cb94298d9e42397772/go1.24.5-20250708.7.darwin-amd64.tar.gz", SHA256: "1b32a6d3844b4bb28acc4e6b6cec743773d7805251dd3ef6fad4be72e5a8f103"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4ba02eab-f12f-46f0-8d2e-561a36985ed3/591b79316cd2f19626bc4913e82d09da/go1.24.5-20250708.7.darwin-arm64.tar.gz", SHA256: "a9b199003beaafcb693707327302ac996add811ebacf3bd17ad8fcb8417f3e99"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4ba02eab-f12f-46f0-8d2e-561a36985ed3/efd09946006859e3eef6281d7e4365cf/go1.24.5-20250708.7.linux-amd64.tar.gz", SHA256: "109829f011521e7fa68fe9eab62243265a22d1386a7495cbcab92359f000fe09"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4ba02eab-f12f-46f0-8d2e-561a36985ed3/46baf965f9b88d5bdd0c7809010805bf/go1.24.5-20250708.7.linux-armv6l.tar.gz", SHA256: "a097eb6c197f85914082b77fcf2cdddc246c37c72e9beddd11cb4ed62a7a331a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4ba02eab-f12f-46f0-8d2e-561a36985ed3/028f3a99498f240dff599068b0f21514/go1.24.5-20250708.7.linux-arm64.tar.gz", SHA256: "4fe69b51d1dde0442fd8cbae12559828d12236ca17c656d71eff7ff8adfdc82e"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4ba02eab-f12f-46f0-8d2e-561a36985ed3/d0cf2310da75825439d8c643138618ed/go1.24.5-20250708.7.windows-amd64.zip", SHA256: "ae82ab73f975687b27a9ab7dc9f07b16e19f979d64ea3fc8e4e8f019b8e8f655"}
	default:
		return version.MicrosoftArchive{}
	}
}
