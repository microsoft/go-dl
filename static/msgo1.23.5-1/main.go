// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.5-1 command runs the go command from Microsoft build of Go 1.23.5-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.5-1@latest
//	$ msgo1.23.5-1 download
//
// And then use the msgo1.23.5-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.5-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.5-1", "msgo1.23.5-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8276c735-5c96-40cf-89da-30fe80ec54b0/e27b0e8e3259e657e0df46301407f9ca/go1.23.5-20250116.6.linux-amd64.tar.gz", SHA256: "448182d88e40398ea532ec58a6c41d519d0746ee2b6bacb4102cd5b63118b084"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8276c735-5c96-40cf-89da-30fe80ec54b0/f38121fb1ef8092d61ca82e6c756fae4/go1.23.5-20250116.6.linux-armv6l.tar.gz", SHA256: "fb8d77f6bb7d6e157a5783ee518091c43de6947b8ecaf84e0d1805e6cd72ff5e"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8276c735-5c96-40cf-89da-30fe80ec54b0/2f0f259ea7c01f7692eeb8b675826c80/go1.23.5-20250116.6.linux-arm64.tar.gz", SHA256: "f0a601089c1a523f72f645b5d47cc6582dae495f70b41da614fe193d1931e649"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8276c735-5c96-40cf-89da-30fe80ec54b0/a889059793f00f24442b260a71079da0/go1.23.5-20250116.6.windows-amd64.zip", SHA256: "3c7f900619214a06c2eef176c1e44e263c7ff37a7024bac5d7946249b1aadf88"}
	default:
		return version.MicrosoftArchive{}
	}
}
