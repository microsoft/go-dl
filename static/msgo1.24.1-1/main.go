// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.1-1 command runs the go command from Microsoft build of Go 1.24.1-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.1-1@latest
//	$ msgo1.24.1-1 download
//
// And then use the msgo1.24.1-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.1-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.1-1", "msgo1.24.1-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3428561-0ec5-4554-8c3e-f77ccc696490/f139d53d3fce534c6deb63e58e5b7be1/go1.24.1-20250304.4.darwin-amd64.tar.gz", SHA256: "fa5a2d74488d99d75cd8cbd4fa8f2d892932078e92728b0b580468c95488a7c0"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3428561-0ec5-4554-8c3e-f77ccc696490/5798ae818c617e0f8bb8b4b221555997/go1.24.1-20250304.4.darwin-arm64.tar.gz", SHA256: "fe669bea5d0b4a45b135f00ff83e57fb3dc61bc9f96e9d9f001f918d938fbc54"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3428561-0ec5-4554-8c3e-f77ccc696490/22c15eae18e90a5a995dee497c3486b9/go1.24.1-20250304.4.linux-amd64.tar.gz", SHA256: "b0ca85ecc435a93e2ddb626dd9ef7fb6689700a0847b0392eb3e146345a8dea0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3428561-0ec5-4554-8c3e-f77ccc696490/69812aac073f39ac638f18c8bee61d1e/go1.24.1-20250304.4.linux-armv6l.tar.gz", SHA256: "43d42297694b849f4266d6a91cfe50d3a59e958c2ffc1bb4527095bbeb6bd90b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3428561-0ec5-4554-8c3e-f77ccc696490/e256ca0c885e677ad7b8fc233798b579/go1.24.1-20250304.4.linux-arm64.tar.gz", SHA256: "73b1befea457d5967632b2b0b93f8d2c0d899d5b6fbd1396c55d0a015292608b"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f3428561-0ec5-4554-8c3e-f77ccc696490/2f01ada1ec2ac21be6625bd67ffabf61/go1.24.1-20250304.4.windows-amd64.zip", SHA256: "48feae5afe4ff7f4068e0b6424a848de9ce013309b2e24a5102796aa8bd68b6f"}
	default:
		return version.MicrosoftArchive{}
	}
}
