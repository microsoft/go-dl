// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.10-1 command runs the go command from Microsoft build of Go 1.24.10-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.10-1@latest
//	$ msgo1.24.10-1 download
//
// And then use the msgo1.24.10-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.10-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.10-1", "msgo1.24.10-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a3e237d5-f73a-4a88-85ee-c03ec600cba5/32881b43964b9b20fb01b4c0fe04ec67/go1.24.10-20251105.4.darwin-amd64.tar.gz", SHA256: "fc0b35438337b8c3600cbf740132e0c156add4aa536ed021a3bced18a288e2c4"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a3e237d5-f73a-4a88-85ee-c03ec600cba5/3649a582d4990b810c81a4e5d34a474c/go1.24.10-20251105.4.darwin-arm64.tar.gz", SHA256: "3e58df57ac6e221f311101e6b6a7bb22dfcee3a0e41266bec7b77c51b74b60df"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a3e237d5-f73a-4a88-85ee-c03ec600cba5/27388c96ea1aa07b82cb3bdc5f561979/go1.24.10-20251105.4.linux-amd64.tar.gz", SHA256: "d4515bbf4d177f7d070ff7ce0f2548c9f73b659e25224b6f9195845a4641e333"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a3e237d5-f73a-4a88-85ee-c03ec600cba5/1f0bbc1b20cfe791c833bd89d913b532/go1.24.10-20251105.4.linux-armv6l.tar.gz", SHA256: "86968a260e51b67abde3857a886c0e91a0e0934788a2950dd73f1afd187c3efd"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a3e237d5-f73a-4a88-85ee-c03ec600cba5/bfd02e38ec499fb90373c5cd36a26e13/go1.24.10-20251105.4.linux-arm64.tar.gz", SHA256: "81cf3b585667522dc7fe749fc274fcaf650a74b36f91fe64fcdaf64d4be4dc33"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a3e237d5-f73a-4a88-85ee-c03ec600cba5/2a9063cf4b7590a54fa71a9f88ad6c21/go1.24.10-20251105.4.windows-amd64.zip", SHA256: "a0a7c92d3ae9e0919e632e06d85479e25676ce1c42739d3665f0e554d8a0db1d"}
	default:
		return version.MicrosoftArchive{}
	}
}
