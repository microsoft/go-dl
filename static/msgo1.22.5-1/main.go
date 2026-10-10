// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.5-1 command runs the go command from Microsoft build of Go 1.22.5-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.5-1@latest
//	$ msgo1.22.5-1 download
//
// And then use the msgo1.22.5-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.5-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.5-1", "msgo1.22.5-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/766eefd8-51c7-431c-8b58-5136273eced8/d6f0ed417acc7881cb620a7c1bdd0358/go1.22.5-20240702.3.linux-amd64.tar.gz", SHA256: "59963327dc86e8a5f48d9485ed800c26f80b9975c9331736c33005947a02686f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/766eefd8-51c7-431c-8b58-5136273eced8/4fb95c4fadc3278832af8f7184ac3d56/go1.22.5-20240702.3.linux-armv6l.tar.gz", SHA256: "c641ce883b1967fb5df679ee652ebd2482dcbc14e528b36fea942ab69b4ca2f2"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/766eefd8-51c7-431c-8b58-5136273eced8/29075dcd0471b2eb9980179446ad072c/go1.22.5-20240702.3.linux-arm64.tar.gz", SHA256: "92078bdb0a2cbe1a5c175b3046427eca30f137b57106157d0696099254bbe03d"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/766eefd8-51c7-431c-8b58-5136273eced8/dee47d15cb8f7b0cfce617cd7663e4d3/go1.22.5-20240702.3.windows-amd64.zip", SHA256: "8310ad1878f1420c36a67145e26ac3163c8be65fd93e61b09eca5544a21120d5"}
	default:
		return version.MicrosoftArchive{}
	}
}
