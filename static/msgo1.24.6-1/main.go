// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.6-1 command runs the go command from Microsoft build of Go 1.24.6-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.6-1@latest
//	$ msgo1.24.6-1 download
//
// And then use the msgo1.24.6-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.6-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.6-1", "msgo1.24.6-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2494a00a-04e6-4e77-bef0-d209c086028b/c516ff27e43b5e533d020792a6591047/go1.24.6-20250806.4.darwin-amd64.tar.gz", SHA256: "2db70fd37cf8bcadb45536f7abcbb4e20f5840c383072a7c4bfc1db2e1d5b8a3"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2494a00a-04e6-4e77-bef0-d209c086028b/b2fb4a21da4c869de586a72d91a16304/go1.24.6-20250806.4.darwin-arm64.tar.gz", SHA256: "7efaa48190fc0f33655e07a076f6b768a9d01122508663cd709247699d64c731"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2494a00a-04e6-4e77-bef0-d209c086028b/b5fce78da80cb514412a586e3ad28e80/go1.24.6-20250806.4.linux-amd64.tar.gz", SHA256: "b9635be15acdb15bc048fd71965de96b5107fac620bc75241968c189c0fc1da9"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2494a00a-04e6-4e77-bef0-d209c086028b/7053cf8ba4bdf89ba2b2707b11817d5a/go1.24.6-20250806.4.linux-armv6l.tar.gz", SHA256: "6dd6ce1d5a4f9f602ad2bd6fba11f72f4d288f6270124681d3c81118130bde79"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2494a00a-04e6-4e77-bef0-d209c086028b/68a2f2faccfabc9a51892d5c989f618c/go1.24.6-20250806.4.linux-arm64.tar.gz", SHA256: "947f5a4f7cd3a05aa2f0d6cba2a51c9d8acc960aae37bf02a0713f93cf6a36ab"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2494a00a-04e6-4e77-bef0-d209c086028b/fd939574ffdfa938a8f239e4ea47addd/go1.24.6-20250806.4.windows-amd64.zip", SHA256: "6aa4499c7e9151362a143a3a8d41ff34ecff0169e7a7496ce184c85a7f411871"}
	default:
		return version.MicrosoftArchive{}
	}
}
