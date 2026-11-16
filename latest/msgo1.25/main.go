// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25 command runs the go command from Microsoft build of Go 1.25.14-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/latest/msgo1.25@latest
//	$ msgo1.25 download
//
// And then use the msgo1.25 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.14-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25", "msgo1.25.14-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0e586b74-a55b-4a5b-bdc1-faf2f85bd042/f0ba5963da0f5346b9d315fc0c06e816/go1.25.14-20261106.4.darwin-amd64.tar.gz", SHA256: "25fa2ad7a768e64eaedb3a8afed403e96ac6a2cc726cfc3c036854d98c74d5f6"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0e586b74-a55b-4a5b-bdc1-faf2f85bd042/ca343634672a356d777912b21390bead/go1.25.14-20261106.4.darwin-arm64.tar.gz", SHA256: "a0f5841f2631d48434a1071d33385547c683e09c327f9d3364fc21f0bc5b0808"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0e586b74-a55b-4a5b-bdc1-faf2f85bd042/60ef1ce47eeb4481681985514c8e171a/go1.25.14-20261106.4.linux-amd64.tar.gz", SHA256: "4603159f08faf898c1f82d59376e1a2661c0e771e3846a354f5779df2eda5944"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0e586b74-a55b-4a5b-bdc1-faf2f85bd042/9680c88335822bc1c802b1cfb12c613f/go1.25.14-20261106.4.linux-armv6l.tar.gz", SHA256: "f2814b4fbb6df00010a470be92643ccabeefbb46a0ae89cf39caadf5eb8486e4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0e586b74-a55b-4a5b-bdc1-faf2f85bd042/50fabe0c9e1250f45d4bf0b5133349e5/go1.25.14-20261106.4.linux-arm64.tar.gz", SHA256: "c9955fa5a0cdc6d7a07e9553831396cb543281495690ccccf57d4ca8b818bcdd"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0e586b74-a55b-4a5b-bdc1-faf2f85bd042/cf1ca70b0673f4689ed28cc91b7cd079/go1.25.14-20261106.4.windows-amd64.zip", SHA256: "a713f038a4694542b546fda234ecbabc1478879fdbfce6c7f9c0408a3f260fad"}
	default:
		return version.MicrosoftArchive{}
	}
}
