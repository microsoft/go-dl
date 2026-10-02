// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.27.0-1 command runs the go command from Microsoft build of Go 1.27.0-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.27.0-1@latest
//	$ msgo1.27.0-1 download
//
// And then use the msgo1.27.0-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.27.0-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.27.0-1", "msgo1.27.0-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1054bb-3e0f-4b3c-b83d-1a07ec46b204/4692f7ef405e09839e9f70234e055fac/go1.27.0-20260819.4.darwin-amd64.tar.gz", SHA256: "34308667fa8a75a72b10cdf51f739acba49a22ba6ec68c7a76d6ed5803a14add"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1054bb-3e0f-4b3c-b83d-1a07ec46b204/2aa0e45f18296b0d96630e004cb3c4e9/go1.27.0-20260819.4.darwin-arm64.tar.gz", SHA256: "601f89342267b81161a5e56f275ab80df3da04a60155dc65a26e98b2d5683eca"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1054bb-3e0f-4b3c-b83d-1a07ec46b204/43f5fdf39b6031eaedcf36dba5db4d73/go1.27.0-20260819.4.linux-amd64.tar.gz", SHA256: "785f37e8655eefdb0d6a240e1a13a4931f2870a64aebf88328cbb8903ef47a01"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1054bb-3e0f-4b3c-b83d-1a07ec46b204/ff2427a4a65ccaab1cc0a2cb5e690413/go1.27.0-20260819.4.linux-armv6l.tar.gz", SHA256: "96def51f72119788a89bf11ec376fd8d4aadfba60b80651fa768a93a756c6353"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1054bb-3e0f-4b3c-b83d-1a07ec46b204/886c60c80bde9ba0d47e7aa8f9ea87db/go1.27.0-20260819.4.linux-arm64.tar.gz", SHA256: "23fa6053262886fa504d6a4e5d9b50ba4f3cb352c85e8aa5d6d1ef5506b8558e"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1054bb-3e0f-4b3c-b83d-1a07ec46b204/c3d7da16e2977e3ec9e44f564f955010/go1.27.0-20260819.4.windows-amd64.zip", SHA256: "291748b70f99ec0d2d3045895e7624dc23b14fbe9396c547e2a94dd55b9eadd1"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1054bb-3e0f-4b3c-b83d-1a07ec46b204/d99324517f53044b6459e06ad5405e9b/go1.27.0-20260819.4.windows-arm64.zip", SHA256: "8a7e7b48fcf05980a9a4890bd3cc645c4ea2cf096e8713bccfc0a77d049f1613"}
	default:
		return version.MicrosoftArchive{}
	}
}
