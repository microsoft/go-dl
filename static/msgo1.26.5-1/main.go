// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.5-1 command runs the go command from Microsoft build of Go 1.26.5-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.5-1@latest
//	$ msgo1.26.5-1 download
//
// And then use the msgo1.26.5-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.5-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.5-1", "msgo1.26.5-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/fd2354b5-c939-4c73-83c0-5077786da0b6/c2df7635b2ac88bce64468842bb54153/go1.26.5-20260707.4.darwin-amd64.tar.gz", SHA256: "d2a409239c9e9a08f9368c57526d35a841b7c782635e70dcfea628326cf7713d"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/fd2354b5-c939-4c73-83c0-5077786da0b6/d97ac77838ee3b8f38d8b3e9f6caeab3/go1.26.5-20260707.4.darwin-arm64.tar.gz", SHA256: "33b11bd9794d53b8e799ebf2b75f9ddb7f80246884c81d0a188b0083f22c9456"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/fd2354b5-c939-4c73-83c0-5077786da0b6/e56b298b1140d2a70f6533d8c296c168/go1.26.5-20260707.4.linux-amd64.tar.gz", SHA256: "4fac685e9ff2c09922a5a6cd9bd14e4170927b4a478e8e44caf4c45d47876425"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/fd2354b5-c939-4c73-83c0-5077786da0b6/1da8f659b118232a344428e8591a1ecb/go1.26.5-20260707.4.linux-armv6l.tar.gz", SHA256: "ae5a14f01160db0dc779689a39e9453072490398597ab07113c1097b41cc6957"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/fd2354b5-c939-4c73-83c0-5077786da0b6/5d186536f3fd1025ecd73faba4c3ef19/go1.26.5-20260707.4.linux-arm64.tar.gz", SHA256: "f8934ec86c8e80117c3864fdd4498d22afd3a2fa9e1d924bdac5fe7417adc963"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/fd2354b5-c939-4c73-83c0-5077786da0b6/f0d11e497b929467d6c0692932e836b4/go1.26.5-20260707.4.windows-amd64.zip", SHA256: "8abe65c07f0b44ff70f8b9dc4db86e7d91652ecde5a79f3a5f9ad3950722f408"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/fd2354b5-c939-4c73-83c0-5077786da0b6/49980349b4e814bc381fe02dda6678e5/go1.26.5-20260707.4.windows-arm64.zip", SHA256: "57b5d84c42b6efed57dfa81b3ff9b495fb5ac42ed33253ae7bc883068137e66b"}
	default:
		return version.MicrosoftArchive{}
	}
}
