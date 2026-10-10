// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.8-1 command runs the go command from Microsoft build of Go 1.25.8-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.8-1@latest
//	$ msgo1.25.8-1 download
//
// And then use the msgo1.25.8-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.8-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.8-1", "msgo1.25.8-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/51526d30-998a-4388-aa59-d3305064bc1f/0e464b74d26964a0715696c7415a8d6b/go1.25.8-20260306.2.darwin-amd64.tar.gz", SHA256: "a59a4554743179f7df752bc65ff910b34cface47705db8009489aa679de76925"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/51526d30-998a-4388-aa59-d3305064bc1f/d7a1a1f668a960244a143605fe2f5aea/go1.25.8-20260306.2.darwin-arm64.tar.gz", SHA256: "df0b34efa879b18127a13c414b6bac4b292e58f8fb84c14e45c9bec2ba30dd8c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/51526d30-998a-4388-aa59-d3305064bc1f/8b6c2b0df43124acef84edc2d5af438b/go1.25.8-20260306.2.linux-amd64.tar.gz", SHA256: "792d70b4942135891ba41626ca55cad89857b6c20fe03422fb88c542131e4550"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/51526d30-998a-4388-aa59-d3305064bc1f/a3c367b6e6ccc3e5df31ecdd4f9443d6/go1.25.8-20260306.2.linux-armv6l.tar.gz", SHA256: "d63ab9475e23956ce152b29ffeec6be1c64dd68c9ac56bf33bc8a88f840e7a45"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/51526d30-998a-4388-aa59-d3305064bc1f/36b21d38bb884677707e9b29b45f9ec0/go1.25.8-20260306.2.linux-arm64.tar.gz", SHA256: "94d9c555c725ee8b9aedc9ded8efd9aaf9b959750cc691b5a8572f73026219e6"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/51526d30-998a-4388-aa59-d3305064bc1f/01efcac05912cf5cfd46359eba906c32/go1.25.8-20260306.2.windows-amd64.zip", SHA256: "a285a13ef7373a2bdf47ce8f270eab96c083bf22d5d2104654d10dc513adccca"}
	default:
		return version.MicrosoftArchive{}
	}
}
