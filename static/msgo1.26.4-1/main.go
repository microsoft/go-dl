// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.4-1 command runs the go command from Microsoft build of Go 1.26.4-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.4-1@latest
//	$ msgo1.26.4-1 download
//
// And then use the msgo1.26.4-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.4-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.4-1", "msgo1.26.4-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/88c9f950-d0f1-46ac-90d0-8e8b4095e743/b4fd14d59150fa3fff40bec803473350/go1.26.4-20260602.8.darwin-amd64.tar.gz", SHA256: "2ae7f888d38d04e891cc5512622da48cadb3c24afc96a3634723ee71ff35042c"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/88c9f950-d0f1-46ac-90d0-8e8b4095e743/c8a11ca9a3d0828665e27739dce2a60a/go1.26.4-20260602.8.darwin-arm64.tar.gz", SHA256: "396705207e67a387f9052d789ac9d0bd2ddfaf5fe57f4bc720ba9881567f8e42"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/88c9f950-d0f1-46ac-90d0-8e8b4095e743/30d6f3c908b93744d77f0cfcc29a8d0b/go1.26.4-20260602.8.linux-amd64.tar.gz", SHA256: "fe7d72e1e83e15633d0a99bd67ce5ffbb67d8a5f71dcea309ff274457ef3c927"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/88c9f950-d0f1-46ac-90d0-8e8b4095e743/0661020e65a0ca7812a1c63367eed92b/go1.26.4-20260602.8.linux-armv6l.tar.gz", SHA256: "594e5c8c552ea19ac232ca124147f926296b36b7cb916ab21b3cf56ef6d724e7"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/88c9f950-d0f1-46ac-90d0-8e8b4095e743/4c37d72249bf2963ca1588c288cb547b/go1.26.4-20260602.8.linux-arm64.tar.gz", SHA256: "9c1c8660dd3cec69415c476232ffb32d529e158d2b2deca4038804f1766f2f78"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/88c9f950-d0f1-46ac-90d0-8e8b4095e743/8761fc8d0a8bc095eb0da86ee4454e0e/go1.26.4-20260602.8.windows-amd64.zip", SHA256: "e90a437d8abba14c9094f17a74fd22124c9c5af11919ca4e51e3da7988d1ceda"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/88c9f950-d0f1-46ac-90d0-8e8b4095e743/7f6b116858edb7fd5814caf7a9e240ca/go1.26.4-20260602.8.windows-arm64.zip", SHA256: "9db0361213b4a6f2e83c8c0d9db0d5029172ab78006d2a708727c890bcd3de27"}
	default:
		return version.MicrosoftArchive{}
	}
}
