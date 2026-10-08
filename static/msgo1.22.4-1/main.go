// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.4-1 command runs the go command from Microsoft build of Go 1.22.4-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.4-1@latest
//	$ msgo1.22.4-1 download
//
// And then use the msgo1.22.4-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.4-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.4-1", "msgo1.22.4-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b09809b-5dda-4fa5-9088-e89676a3bbde/ba511b096cf3e8364bb1b67c6aa8d47a/go1.22.4-20240604.2.linux-amd64.tar.gz", SHA256: "2ea3fa766af6a2bf286c8c37d44fbdf3114a5933892f18ec314302a7ad9aad66"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b09809b-5dda-4fa5-9088-e89676a3bbde/c0792a956df47a23e166ad6569cc395b/go1.22.4-20240604.2.linux-armv6l.tar.gz", SHA256: "6f6206db3720fb2d24cb58cb581192a06c2fdb3b68e37f12cbdfbfb3563e4adc"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b09809b-5dda-4fa5-9088-e89676a3bbde/521cf6b5f936a4e6069d72b6ceb43519/go1.22.4-20240604.2.linux-arm64.tar.gz", SHA256: "e825691e5c9210dfa2f3c676a6204d8ff730aaa5356332dd0d95b3ba4c5225f5"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2b09809b-5dda-4fa5-9088-e89676a3bbde/4a8154ff3efa9ab0cddae8e3f1dead91/go1.22.4-20240604.2.windows-amd64.zip", SHA256: "466202685c5f0c601f5efcab0ce3c693beec9183cb5ee03c8e2e4140476396c7"}
	default:
		return version.MicrosoftArchive{}
	}
}
