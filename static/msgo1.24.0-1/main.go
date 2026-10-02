// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.0-1 command runs the go command from Microsoft build of Go 1.24.0-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.0-1@latest
//	$ msgo1.24.0-1 download
//
// And then use the msgo1.24.0-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.0-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.0-1", "msgo1.24.0-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9fb78595-7514-47a4-8e49-b644c67c371d/1197ec01c47358c81b2baeb2557659f5/go1.24.0-20250212.5.darwin-amd64.tar.gz", SHA256: "1d604ddccb9ac5a91f15b7a1afd8850da54f6b71aad929e30e2f3c65fef5bbda"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9fb78595-7514-47a4-8e49-b644c67c371d/1f04b47105a476c894bb6df01a0a1ab6/go1.24.0-20250212.5.darwin-arm64.tar.gz", SHA256: "434313f09c1ca12eff6650128dfc1de6d2596edbab3a1e47cf2cf459788e9599"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9fb78595-7514-47a4-8e49-b644c67c371d/921cbb2606d808226357fec1d8b699f6/go1.24.0-20250212.5.linux-amd64.tar.gz", SHA256: "95eeca67708c56ab96831957d28315f3ae381711d21d0f3b3492fdfb3e182df8"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9fb78595-7514-47a4-8e49-b644c67c371d/28d2cd432167bce2bb7bc5d84c9b80f2/go1.24.0-20250212.5.linux-armv6l.tar.gz", SHA256: "3450e57566660c370d9a94c44a03b214dd90a6bccab8062447326f02707acc66"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9fb78595-7514-47a4-8e49-b644c67c371d/4ba54854dc5d7176dbf9deee807d300e/go1.24.0-20250212.5.linux-arm64.tar.gz", SHA256: "5c58f2a182d9cf9d48e3fb1cc81136980ca9a5b17b0a8e2ea95844fc72408154"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9fb78595-7514-47a4-8e49-b644c67c371d/e71fc2828b4431667977977f5b44dbbe/go1.24.0-20250212.5.windows-amd64.zip", SHA256: "386043e46944a4cb884cadee89112fb9fd7a5d24a86b8a2d4196a362eb65c9f3"}
	default:
		return version.MicrosoftArchive{}
	}
}
