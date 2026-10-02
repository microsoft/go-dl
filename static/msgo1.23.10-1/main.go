// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.10-1 command runs the go command from Microsoft build of Go 1.23.10-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.10-1@latest
//	$ msgo1.23.10-1 download
//
// And then use the msgo1.23.10-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.10-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.10-1", "msgo1.23.10-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a32ea32f-4963-4971-ba28-b37e2855a6c7/11ab4cd86102ebe50cd2b11c0ba98ed1/go1.23.10-20250605.3.linux-amd64.tar.gz", SHA256: "c4ee22c0de9fd2ff2c0d0d1a446e5aac75c9e3a3bfdc16e25e32d2461c5347ed"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a32ea32f-4963-4971-ba28-b37e2855a6c7/5307727ce24461dc83787c3a80b90ab7/go1.23.10-20250605.3.linux-armv6l.tar.gz", SHA256: "38db1b6c242d1d5cf8baeb7229c40f227351f5b198f2a7deb82c820a33771f89"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a32ea32f-4963-4971-ba28-b37e2855a6c7/7ad7991a18bd6147f8da65f1ba1d216b/go1.23.10-20250605.3.linux-arm64.tar.gz", SHA256: "901f023a7ad01b743f4f1f7c6e5dd12c03c2aaaebedd70f7eef415caae79510d"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a32ea32f-4963-4971-ba28-b37e2855a6c7/e0c46f5aa57373ddb27a717d351d542e/go1.23.10-20250605.3.windows-amd64.zip", SHA256: "eb3a9b30afabca33a8f4d127b58d050b5d35417fef4d3a173c75a451a2211ccf"}
	default:
		return version.MicrosoftArchive{}
	}
}
