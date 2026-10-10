// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.11-1 command runs the go command from Microsoft build of Go 1.25.11-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.11-1@latest
//	$ msgo1.25.11-1 download
//
// And then use the msgo1.25.11-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.11-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.11-1", "msgo1.25.11-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2e0e3c7d-1d8b-4df5-a4c2-5155a0cf1f0d/92183f93ab4a47f982a18268db9d4cb2/go1.25.11-20260602.7.darwin-amd64.tar.gz", SHA256: "1f61b372134096ec02d1d3a657ae0587ebde26a1c62687fa83e3da0b2be37e94"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2e0e3c7d-1d8b-4df5-a4c2-5155a0cf1f0d/d655a44b247e6b2f32906e4aa0ee3faf/go1.25.11-20260602.7.darwin-arm64.tar.gz", SHA256: "4813ee327b393437eca1ea97ad766608ed61cdc03cfd8be9a6d515a8198522bb"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2e0e3c7d-1d8b-4df5-a4c2-5155a0cf1f0d/42f3a47ae62b2ce3ba49d46952aba281/go1.25.11-20260602.7.linux-amd64.tar.gz", SHA256: "180ac77cb448e0f7f356e79838c04c1a5fc1564879af0c2ba9aa392c56692ab6"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2e0e3c7d-1d8b-4df5-a4c2-5155a0cf1f0d/503b6bd64dc607e527d8f211c6b83520/go1.25.11-20260602.7.linux-armv6l.tar.gz", SHA256: "839cc2f41e6bf334ebe2d7e28b8dae07d7be12d50ae38102920790fdcdebc953"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2e0e3c7d-1d8b-4df5-a4c2-5155a0cf1f0d/9b90cfc869ea3cae5d41f5fe281f32bb/go1.25.11-20260602.7.linux-arm64.tar.gz", SHA256: "52c2bdbcdd2d808966375b117ad63b4ee04550ffbdbbed00762decde3a8c9730"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2e0e3c7d-1d8b-4df5-a4c2-5155a0cf1f0d/c7566805723ceb7b901e2f699ac22df3/go1.25.11-20260602.7.windows-amd64.zip", SHA256: "c5fe3aa61971d5b23d50efe8f2c8c6379031b4ab470b0abadfd5593a2cc6ba0d"}
	default:
		return version.MicrosoftArchive{}
	}
}
