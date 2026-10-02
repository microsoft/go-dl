// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.13-1 command runs the go command from Microsoft build of Go 1.24.13-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.13-1@latest
//	$ msgo1.24.13-1 download
//
// And then use the msgo1.24.13-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.13-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.13-1", "msgo1.24.13-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afbe31b0-e28e-4753-b212-1dabc02e90e9/6fa428c0993354d1ff8e69cf85d35cf9/go1.24.13-20260204.5.darwin-amd64.tar.gz", SHA256: "a313083f4b3e23ab037a617fcb72b1f15f0c969dec1f6e3c7326c04f317b6ff3"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afbe31b0-e28e-4753-b212-1dabc02e90e9/2857b4ff3c8c98d90c0201586bf25fec/go1.24.13-20260204.5.darwin-arm64.tar.gz", SHA256: "efe122af27011a935843df83ac7097b2371282ea26e33a030ad85287fce174fd"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afbe31b0-e28e-4753-b212-1dabc02e90e9/08e84dd28f4d4caadffbc69257f1bff3/go1.24.13-20260204.5.linux-amd64.tar.gz", SHA256: "12298e9f9ac99fedfbac6d5c12cac1204102972023952c94602205e7893903c8"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afbe31b0-e28e-4753-b212-1dabc02e90e9/53ead640781478c3575935427017a590/go1.24.13-20260204.5.linux-armv6l.tar.gz", SHA256: "b7a2b8fce672a346c93c6e28046d2bd7d11fd5698f0a9b0349b7673a4dd08723"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afbe31b0-e28e-4753-b212-1dabc02e90e9/e74fd223f2d0884f6b711602ca5e7148/go1.24.13-20260204.5.linux-arm64.tar.gz", SHA256: "64231be88e1dfdcdb8cc94fa96fdeecacd6265088bac13f83a90d754bc4209f3"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afbe31b0-e28e-4753-b212-1dabc02e90e9/13b9e90864d76f06dfaf8da78aa87799/go1.24.13-20260204.5.windows-amd64.zip", SHA256: "681971df418347fb84fa63e0866822e683d6ee6c17fc308bd54448096deec32e"}
	default:
		return version.MicrosoftArchive{}
	}
}
