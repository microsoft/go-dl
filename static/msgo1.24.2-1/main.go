// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.2-1 command runs the go command from Microsoft build of Go 1.24.2-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.2-1@latest
//	$ msgo1.24.2-1 download
//
// And then use the msgo1.24.2-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.2-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.2-1", "msgo1.24.2-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/916e58f6-7b0b-4ea5-940a-68d169d0edb5/6702f84f439a62f63e0f49fb0ab0f04f/go1.24.2-20250401.4.darwin-amd64.tar.gz", SHA256: "a9710d74119f41a5d5717d6a0e35d97065b888a764e966f22e0fd98447c817d0"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/916e58f6-7b0b-4ea5-940a-68d169d0edb5/21e3327948f7a56489115b1e0e5a9916/go1.24.2-20250401.4.darwin-arm64.tar.gz", SHA256: "3fe64369cdb35aa17cd8a77de4766671f3842c912ae9dcffa75e24cf4e385911"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/916e58f6-7b0b-4ea5-940a-68d169d0edb5/21a7d72c9a1617cd21551f053de28626/go1.24.2-20250401.4.linux-amd64.tar.gz", SHA256: "7844f426e83fb66b3d109357214abcb7cce4ed87178179ef413efca3bce0303a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/916e58f6-7b0b-4ea5-940a-68d169d0edb5/9367f37ab4dd363e96f98cacbe4fa770/go1.24.2-20250401.4.linux-armv6l.tar.gz", SHA256: "3d9b518e07b4bf8db8c6c5454a9043d61b97de9cf95b57b9465f4fd4adc37002"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/916e58f6-7b0b-4ea5-940a-68d169d0edb5/0ef299f34d5024689021c7aaf1d98981/go1.24.2-20250401.4.linux-arm64.tar.gz", SHA256: "26531b0ff1add9922b05b0dc8535dce5a9ca09517b6320d21d0dd77dc7f1f536"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/916e58f6-7b0b-4ea5-940a-68d169d0edb5/c5850019d034d3d550c00b9fb71b3e0e/go1.24.2-20250401.4.windows-amd64.zip", SHA256: "d499ae352639f97620152a204abac347a3b3aa0ff4b5c8cc066b57a5738d8385"}
	default:
		return version.MicrosoftArchive{}
	}
}
