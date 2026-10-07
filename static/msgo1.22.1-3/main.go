// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.1-3 command runs the go command from Microsoft build of Go 1.22.1-3.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.1-3@latest
//	$ msgo1.22.1-3 download
//
// And then use the msgo1.22.1-3 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.1-3.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.1-3", "msgo1.22.1-3", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8274b236-39f7-48f6-919f-d57f4897aa66/4ef6e8c9b263e52a9f3eb6a608e1dd3a/go1.22.1-20240321.8.linux-amd64.tar.gz", SHA256: "2928adebf5139a4ed9365145357b5d4cd72c4de68b773bea76a6013e27589229"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8274b236-39f7-48f6-919f-d57f4897aa66/bcd974f192fd061237f5e4dc7e06447a/go1.22.1-20240321.8.linux-armv6l.tar.gz", SHA256: "9ab24e136fcf71d32f9024be641f6bb4dc95f7e81ddd29a3134aa431d32aea42"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8274b236-39f7-48f6-919f-d57f4897aa66/b7c54318087b4d84c72d81892d883b57/go1.22.1-20240321.8.linux-arm64.tar.gz", SHA256: "b74664c8d5466ca3c9e3ed460c773608b532bc623643e4916c5a8bc790084bd8"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8274b236-39f7-48f6-919f-d57f4897aa66/2656b701af0428a24c94488eca82fd40/go1.22.1-20240321.8.windows-amd64.zip", SHA256: "fc768bebca774e482687ef1d9c3f19bdd3841fcb60be7dff0b9708ecc3fc97f5"}
	default:
		return version.MicrosoftArchive{}
	}
}
