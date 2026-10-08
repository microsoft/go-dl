// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.3-1 command runs the go command from Microsoft build of Go 1.25.3-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.3-1@latest
//	$ msgo1.25.3-1 download
//
// And then use the msgo1.25.3-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.3-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.3-1", "msgo1.25.3-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d87093ec-184f-45db-8a05-25cc47cdf3b5/d4190e6aa87e3ac003e4c314f59330a2/go1.25.3-20251013.2.darwin-amd64.tar.gz", SHA256: "200489b757371e27527c4987a07497cc0927678cc53e0bb44345857ee20cd308"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d87093ec-184f-45db-8a05-25cc47cdf3b5/9c0a2bd58eb34f65eb1b3bf089e58555/go1.25.3-20251013.2.darwin-arm64.tar.gz", SHA256: "67dd94d7a71c72d814592b850c0658087ae5d9efb3779dd4cd6598ce080975a4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d87093ec-184f-45db-8a05-25cc47cdf3b5/119bb0e18622204bfa873fa0e6eec679/go1.25.3-20251013.2.linux-amd64.tar.gz", SHA256: "304d4ce3aeee2d5ce3cca73aebeb97ea238d0dadf5c2cc7e59aeab9ff08ab7e9"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d87093ec-184f-45db-8a05-25cc47cdf3b5/7d430bba90d7edc80beb8777466d4405/go1.25.3-20251013.2.linux-armv6l.tar.gz", SHA256: "4648174ab39421c9ee7283048e6ce4e3d53c89cfdf926eab25bae2f75052408c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d87093ec-184f-45db-8a05-25cc47cdf3b5/b74b54e20f351b8a08be31aeec737bd5/go1.25.3-20251013.2.linux-arm64.tar.gz", SHA256: "e43a7f7a806cf64ceba013e78a98e92d55d2a8e146d367ff9c22966d60aa25ce"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d87093ec-184f-45db-8a05-25cc47cdf3b5/6d91b29e82af91a365fb1f6134bf67b4/go1.25.3-20251013.2.windows-amd64.zip", SHA256: "7c44c55744fcdd8eb486e3c156f6b2357a087c9bb2e5eb7eb8808ee327d14d11"}
	default:
		return version.MicrosoftArchive{}
	}
}
