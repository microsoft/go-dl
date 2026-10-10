// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.5-1 command runs the go command from Microsoft build of Go 1.25.5-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.5-1@latest
//	$ msgo1.25.5-1 download
//
// And then use the msgo1.25.5-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.5-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.5-1", "msgo1.25.5-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0834b1f-21be-44ff-9a19-4c2adb01a258/237aa2c0c825367e2c0f5a18b72ef914/go1.25.5-20251202.2.darwin-amd64.tar.gz", SHA256: "8775dbfa5705d9cd635bb8fcd0d37898717157392f8bc823bf419bb5482b79fb"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0834b1f-21be-44ff-9a19-4c2adb01a258/f4cc0db37a1c9c3194adb65ccf04246e/go1.25.5-20251202.2.darwin-arm64.tar.gz", SHA256: "ded8b99c053745a797bf038f8d1e10a00e4d118c23060e27abd9fdcde5001b08"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0834b1f-21be-44ff-9a19-4c2adb01a258/f5607cc764efca24e89ed358efdcbd89/go1.25.5-20251202.2.linux-amd64.tar.gz", SHA256: "48e42b9e3554e106326d326624733fa7da9c4f315cd7e04d764adeb3f285c582"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0834b1f-21be-44ff-9a19-4c2adb01a258/529117d0dfc89d571cdc15c3269f9aeb/go1.25.5-20251202.2.linux-armv6l.tar.gz", SHA256: "d219799b19150cb2f0715f6b5cbd42fb7573263e8b545b124dfa064667897926"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0834b1f-21be-44ff-9a19-4c2adb01a258/fb712ade5018e3eeffc05875321ef99d/go1.25.5-20251202.2.linux-arm64.tar.gz", SHA256: "011bb6986c2fd978d5a74e70b5d3905ee6e35a24137732b78ec319dfdbdb077a"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d0834b1f-21be-44ff-9a19-4c2adb01a258/cd04b5375706deb87a2e339aa6a6fec4/go1.25.5-20251202.2.windows-amd64.zip", SHA256: "d2520eb0ba754f33278c07ecd6e0165ad83de1af6c6d9484b27a3cae2a595081"}
	default:
		return version.MicrosoftArchive{}
	}
}
