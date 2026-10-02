// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.1-3 command runs the go command from Microsoft build of Go 1.23.1-3.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.1-3@latest
//	$ msgo1.23.1-3 download
//
// And then use the msgo1.23.1-3 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.1-3.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.1-3", "msgo1.23.1-3", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4a8fa8ad-2eaa-4f6a-bb60-30f0cba0c700/7d67ce8d03fbf51346a195c66a591861/go1.23.1-20240925.6.linux-amd64.tar.gz", SHA256: "57567261b4f5a15a2aa65d8ec196502afa912fe971047f720bc194cb50238194"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4a8fa8ad-2eaa-4f6a-bb60-30f0cba0c700/e260e1e4bb509c03ca15995f04f1d3fb/go1.23.1-20240925.6.linux-armv6l.tar.gz", SHA256: "c5800c4105a99c93236d3aac2c9b6372e939dfcf48b6e35bf34bf997b922c6d5"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4a8fa8ad-2eaa-4f6a-bb60-30f0cba0c700/7868ffa77c98ac8279c0c8a1a766e889/go1.23.1-20240925.6.linux-arm64.tar.gz", SHA256: "69b67133e4951df8866a579b09fd1bd28056c3e9db973522b51b7768e6f294f2"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4a8fa8ad-2eaa-4f6a-bb60-30f0cba0c700/17b256bcb0d52dd4ad300a61cd9e3626/go1.23.1-20240925.6.windows-amd64.zip", SHA256: "6ec0dd5184eb7383413ebed4093f7d129017ca52583c52ce09c3a12eee9a9919"}
	default:
		return version.MicrosoftArchive{}
	}
}
