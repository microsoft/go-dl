// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.14-1 command runs the go command from Microsoft build of Go 1.25.14-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.14-1@latest
//	$ msgo1.25.14-1 download
//
// And then use the msgo1.25.14-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.14-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.14-1", "msgo1.25.14-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/45c9587d-44d5-49d5-93bb-7523b6a10fa1/4e15601183260f4ca346d48faebe6a04/go1.25.14-20260819.1.darwin-amd64.tar.gz", SHA256: "10c30c205ce2d4d74138ba0efa70246a2548a2d5aea19a94315b86f4e4255f7d"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/45c9587d-44d5-49d5-93bb-7523b6a10fa1/50bcaea7a1ad76f4dc046ba0884921e9/go1.25.14-20260819.1.darwin-arm64.tar.gz", SHA256: "39ce7e3e11bd8bcae608e857797daee1d99467ab9d39554f393965e6183ab6a0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/45c9587d-44d5-49d5-93bb-7523b6a10fa1/d6b8a06a6ed07b16914e0af653478360/go1.25.14-20260819.1.linux-amd64.tar.gz", SHA256: "25e61684d45ee160b0a9f92b567ec8a2ba96859ccb00b186aa1fba89c5f25ed4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/45c9587d-44d5-49d5-93bb-7523b6a10fa1/0f5e5ce4918ba024e193f22aaa7f14d1/go1.25.14-20260819.1.linux-armv6l.tar.gz", SHA256: "9e2e6cac3705f329d572b33e630d95599b31d53d94f864dc5eeac5cb12f992db"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/45c9587d-44d5-49d5-93bb-7523b6a10fa1/bb0e2ec09ee38bad654400b1740bc7c3/go1.25.14-20260819.1.linux-arm64.tar.gz", SHA256: "3e97236ea70c2b8743de9d3e68c554d9d642d4ad9ed44f5e1d27d5504d86ad62"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/45c9587d-44d5-49d5-93bb-7523b6a10fa1/50bf8e9cc5b6ef37d2d3fad3a2922162/go1.25.14-20260819.1.windows-amd64.zip", SHA256: "ed0c61fc1df342cb9b9a2194f5fec3e8974d8b3ae829c175febec7b9fe0c9cbb"}
	default:
		return version.MicrosoftArchive{}
	}
}
