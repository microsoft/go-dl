// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.4-2 command runs the go command from Microsoft build of Go 1.25.4-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.4-2@latest
//	$ msgo1.25.4-2 download
//
// And then use the msgo1.25.4-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.4-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.4-2", "msgo1.25.4-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a95f333f-4441-43a8-aa1c-7eb364a31011/c444b0fd7a976555880670022eec84ec/go1.25.4-20251106.4.darwin-amd64.tar.gz", SHA256: "e441afb1607027d6ae6c4cb98148d80739a2b5c2847066e6471bf9ed603a5b74"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a95f333f-4441-43a8-aa1c-7eb364a31011/cc89e3fa3b03536a1048ff3522408da2/go1.25.4-20251106.4.darwin-arm64.tar.gz", SHA256: "352f6536afd39aa02cd20f0251b58f28928b0dd6f442ad6a6c649d022d7c490b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a95f333f-4441-43a8-aa1c-7eb364a31011/71a87bbcd1fb4054787995152cd20f31/go1.25.4-20251106.4.linux-amd64.tar.gz", SHA256: "602b5ff65678b7a9f88a9ed27ceb40182768fc3565aa3ec58dd0be56867a375f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a95f333f-4441-43a8-aa1c-7eb364a31011/b0280aa1d32c83e4b4e2d53d16e1ec50/go1.25.4-20251106.4.linux-armv6l.tar.gz", SHA256: "81c0a8df7e182afa2916df160a7037062549177dc463c1bf3eaac8b084066709"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a95f333f-4441-43a8-aa1c-7eb364a31011/ea33ba23b5ffbeee2ecdd566ef81d775/go1.25.4-20251106.4.linux-arm64.tar.gz", SHA256: "93dbb86f7e50058f8168113ce99cbe601e8c1485cbd2f89840eb002d9f6170c1"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a95f333f-4441-43a8-aa1c-7eb364a31011/9a91b6da21473ca14d18c87aef6f0723/go1.25.4-20251106.4.windows-amd64.zip", SHA256: "b8fca22a73e0f12c32d4e3360239e4ae6b3bfdf844333ce24ad2b9a2c9cc2fba"}
	default:
		return version.MicrosoftArchive{}
	}
}
