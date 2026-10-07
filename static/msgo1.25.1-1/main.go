// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.1-1 command runs the go command from Microsoft build of Go 1.25.1-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.1-1@latest
//	$ msgo1.25.1-1 download
//
// And then use the msgo1.25.1-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.1-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.1-1", "msgo1.25.1-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2aca50c5-e32f-4c8b-80fd-3ff5a35f2a8d/9448a46e36f279801408fab0ead385ff/go1.25.1-20250903.7.darwin-amd64.tar.gz", SHA256: "e6f1e96d3af27f5810aa0fb870bcd5b2023465f3baf9f1f2d7d2943636b9cbd5"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2aca50c5-e32f-4c8b-80fd-3ff5a35f2a8d/c81e94c67addcaaf3f8791e0ffe37596/go1.25.1-20250903.7.darwin-arm64.tar.gz", SHA256: "71ef8e9fcb73bb44b79ca287fe4b36811666ed74d50e9cb3740a4a597f991275"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2aca50c5-e32f-4c8b-80fd-3ff5a35f2a8d/923c5727dfd0c9115b9933959d6b700e/go1.25.1-20250903.7.linux-amd64.tar.gz", SHA256: "d99f6457a5ebdd2822d8b7660b6afd9f7a01d969d8984df8e156617ebb6b7cd0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2aca50c5-e32f-4c8b-80fd-3ff5a35f2a8d/f16ccee0bdc85dd874607b7f1b4fc106/go1.25.1-20250903.7.linux-armv6l.tar.gz", SHA256: "832ffb443523bf8e5ea49aed4bce0f6e1a05c494aca05c54a60c4b822473c0bd"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2aca50c5-e32f-4c8b-80fd-3ff5a35f2a8d/49cfc8f6f09707eac76763e35494d4f8/go1.25.1-20250903.7.linux-arm64.tar.gz", SHA256: "fe440299a943d97175860a4b8f1913625ffd663abf2081f551b75b40323434b0"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/2aca50c5-e32f-4c8b-80fd-3ff5a35f2a8d/cf70bc40550866f6b224b085377105ec/go1.25.1-20250903.7.windows-amd64.zip", SHA256: "2c22c35c5064208552a0d66ea200a157d853e97de8f48eec8758c9e515dda2d0"}
	default:
		return version.MicrosoftArchive{}
	}
}
