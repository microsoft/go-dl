// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.12-2 command runs the go command from Microsoft build of Go 1.22.12-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.12-2@latest
//	$ msgo1.22.12-2 download
//
// And then use the msgo1.22.12-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.12-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.12-2", "msgo1.22.12-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/689e8305-5d7b-4a4b-b4e0-929d03f1aaeb/c07d05c0c4525452ad6c1b4f71d2d650/go1.22.12-20250211.4.linux-amd64.tar.gz", SHA256: "b133255677ed25bce1459b703280bfa353d6893a14fb9965d6442eb13ca577ec"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/689e8305-5d7b-4a4b-b4e0-929d03f1aaeb/ce79732779e13b4fb7009b6260f005f5/go1.22.12-20250211.4.linux-armv6l.tar.gz", SHA256: "2e8c03489f8d587b48d193f978708de0bde74a75fab6671eb1e978c911832f5c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/689e8305-5d7b-4a4b-b4e0-929d03f1aaeb/f804fa2d745b4affbe76bf23e112c6c4/go1.22.12-20250211.4.linux-arm64.tar.gz", SHA256: "81635f8e685ac253dbc43e7e12072435b4baf9b2efecc8236e74c5abe3a68608"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/689e8305-5d7b-4a4b-b4e0-929d03f1aaeb/75b7fed16bb93e84d24ed1067514d293/go1.22.12-20250211.4.windows-amd64.zip", SHA256: "0159d246327db5192f92f5f9fa89ecedf4c80a0beb0cfe1104ff28c0eeb1afee"}
	default:
		return version.MicrosoftArchive{}
	}
}
