// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.7-1 command runs the go command from Microsoft build of Go 1.26.7-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.7-1@latest
//	$ msgo1.26.7-1 download
//
// And then use the msgo1.26.7-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.7-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.7-1", "msgo1.26.7-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/41b3c5b3-4064-4059-b8ee-7ef7bd749531/598ae6821354892970c521c293d97d6b/go1.26.7-20260819.2.darwin-amd64.tar.gz", SHA256: "e559e3e989b1d6e1178de2c531404ef16dd5dfa0c00d3260d0405f193e62cb22"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/41b3c5b3-4064-4059-b8ee-7ef7bd749531/9fbc9d3e5940f0d5f3a3897b0072683e/go1.26.7-20260819.2.darwin-arm64.tar.gz", SHA256: "5dc3042e4c2ac09e0b316807dc2d03ff936f69384ef1190021f30eac526b2a7b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/41b3c5b3-4064-4059-b8ee-7ef7bd749531/2213a63f5db48361095a0100894e06c1/go1.26.7-20260819.2.linux-amd64.tar.gz", SHA256: "8c7c598971462ff2c7b304ee874a8939b68c52a36036c1fd55d1ed0404452711"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/41b3c5b3-4064-4059-b8ee-7ef7bd749531/3cd1b2dc5269cdc1e573136c7cdcf6e7/go1.26.7-20260819.2.linux-armv6l.tar.gz", SHA256: "7e674ff53d8545dae31539d85072989fe108cf991702a622984ecf512476078b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/41b3c5b3-4064-4059-b8ee-7ef7bd749531/5f5e830f5817adb1ba1e27eb393c70c6/go1.26.7-20260819.2.linux-arm64.tar.gz", SHA256: "17f2001fa8d2ad1a73f68452f472f5aa66db483474a4eaeb34efcbc6f8c7c8c3"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/41b3c5b3-4064-4059-b8ee-7ef7bd749531/d6ee592d71e5854f0a103d03ecf9bf77/go1.26.7-20260819.2.windows-amd64.zip", SHA256: "763bc4f90ab616ed7b13e0f9079b4344434a25b83abc72dff4fe81f9153b9adb"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/41b3c5b3-4064-4059-b8ee-7ef7bd749531/b73d38c944214f742251df315e2c6d25/go1.26.7-20260819.2.windows-arm64.zip", SHA256: "eef4d1265ffda5d204fff13be2052589036b2f34a1a8dcf1913ef3a88a9cf5d1"}
	default:
		return version.MicrosoftArchive{}
	}
}
