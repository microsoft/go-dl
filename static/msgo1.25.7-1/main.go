// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.7-1 command runs the go command from Microsoft build of Go 1.25.7-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.7-1@latest
//	$ msgo1.25.7-1 download
//
// And then use the msgo1.25.7-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.7-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.7-1", "msgo1.25.7-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/39e12277-1463-47e4-9f91-f71f8c450c61/16503ff2f5ff30261b0602f647300571/go1.25.7-20260204.4.darwin-amd64.tar.gz", SHA256: "3cca680b5619449e436a15e4d799d6ae3c4e79616637b4c6fadc0b186df8e61a"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/39e12277-1463-47e4-9f91-f71f8c450c61/cda77c95a10a9e64d75094d1cdc79d03/go1.25.7-20260204.4.darwin-arm64.tar.gz", SHA256: "e8ebdd66099a1e216adb643355f040416cd27077c2215681f9e7419d98266679"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/39e12277-1463-47e4-9f91-f71f8c450c61/b6ae036e2c246d52bd15d5300994c8ee/go1.25.7-20260204.4.linux-amd64.tar.gz", SHA256: "328b235ebd36ddc92c852b6d25b01d8e9478c2779bbbac3c8964890118a8aa6f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/39e12277-1463-47e4-9f91-f71f8c450c61/eb2a4cf3b0aa15eb2d83f1bf5908e739/go1.25.7-20260204.4.linux-armv6l.tar.gz", SHA256: "55c53b2e1f9bc2f3bbdcb40733249813238be74367855b76c67560d85f696a24"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/39e12277-1463-47e4-9f91-f71f8c450c61/218094e73bfc123b8065805487bd52fa/go1.25.7-20260204.4.linux-arm64.tar.gz", SHA256: "73c5ffdedf6cf5ec57eb2478a64ab50a363e17cb1f53e37049d4a42b2886f939"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/39e12277-1463-47e4-9f91-f71f8c450c61/e75fa6e15650a4cb8ba231404e6a4949/go1.25.7-20260204.4.windows-amd64.zip", SHA256: "d9fca6ce53fbe1d4741ff61e7290dedc8500dad1a7167edc59baedf91553b54e"}
	default:
		return version.MicrosoftArchive{}
	}
}
