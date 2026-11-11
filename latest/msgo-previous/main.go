// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo-previous command runs the go command from Microsoft build of Go 1.26.11-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/latest/msgo-previous@latest
//	$ msgo-previous download
//
// And then use the msgo-previous command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.11-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo-previous", "msgo1.26.11-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/54cd6c58-3476-493c-8076-398285e69f34/1e23967d2ecbed85f60568e05acbc493/go1.26.11-20261110.4.darwin-amd64.tar.gz", SHA256: "c129112bfa7b9d52e58f4689ca1ba6fb98a7ef4300ba3a71e2ae0b4b206d7ca0"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/54cd6c58-3476-493c-8076-398285e69f34/747f1b636324e209e55c19ba8a04e668/go1.26.11-20261110.4.darwin-arm64.tar.gz", SHA256: "692eb12ca2e2539a674ba7b174eaa796b6adf01594175f8e7a679ee7086f621a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/54cd6c58-3476-493c-8076-398285e69f34/ea9353fcf64a4653a95954f8aace37c5/go1.26.11-20261110.4.linux-amd64.tar.gz", SHA256: "6d91a8577474f1318d01a2d03e41bc5b43fedb4c4971a3d4967d7ddf052ccc0f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/54cd6c58-3476-493c-8076-398285e69f34/9f91cae3f0e27e31b1cb1170cedd10ee/go1.26.11-20261110.4.linux-armv6l.tar.gz", SHA256: "3a826585cccf0b98ad52f4225477339d3abe39539005ea761d561e1514720911"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/54cd6c58-3476-493c-8076-398285e69f34/f5772b3fb456c9e831f3cc1cc5486687/go1.26.11-20261110.4.linux-arm64.tar.gz", SHA256: "8a172f6279aca03e4a001262eee5e8d90ee79d9dcb8f55f539fcec11fee0029f"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/54cd6c58-3476-493c-8076-398285e69f34/6c5d46bb6574989355a329eb8cd7c3ca/go1.26.11-20261110.4.windows-amd64.zip", SHA256: "3c9631a1f934f44b17ec43ffafa8691cb3966206203b84e50fc67439979263a6"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/54cd6c58-3476-493c-8076-398285e69f34/6e74bc805dc7d4c89560223c88adcd76/go1.26.11-20261110.4.windows-arm64.zip", SHA256: "26ec73ea50073e8f54a4016fac2eb688f8821f486aaad944515cb0259bac7acc"}
	default:
		return version.MicrosoftArchive{}
	}
}
