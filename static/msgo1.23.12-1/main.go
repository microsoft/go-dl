// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.12-1 command runs the go command from Microsoft build of Go 1.23.12-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.12-1@latest
//	$ msgo1.23.12-1 download
//
// And then use the msgo1.23.12-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.12-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.12-1", "msgo1.23.12-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4561a92e-5ea8-4523-9985-279a84947032/a5f41bbc8c45b6f38c1a2377237484d3/go1.23.12-20250806.6.linux-amd64.tar.gz", SHA256: "3090e3e16cb95801205562c794d27cfa4bcaf8c6e763a5b0c7de90f886715fb0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4561a92e-5ea8-4523-9985-279a84947032/7618ca6f14431d33290a81201eb55e8d/go1.23.12-20250806.6.linux-armv6l.tar.gz", SHA256: "8ef53400cd95432204d7bd1cca4e80cd5dfaed550c4478fbbad63ed112d89d1b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4561a92e-5ea8-4523-9985-279a84947032/15ca41df53d10b71fe387811f7cc9075/go1.23.12-20250806.6.linux-arm64.tar.gz", SHA256: "87276beccb47a4f57deb57e2ce17935d11b458e1e569be581f24457873f82708"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4561a92e-5ea8-4523-9985-279a84947032/351263d7ea5527cc0c11e7d39793b305/go1.23.12-20250806.6.windows-amd64.zip", SHA256: "926628fc8cff3db5e80c6f1c5063545420b1adfc742cd68bc1227a1352ab00c1"}
	default:
		return version.MicrosoftArchive{}
	}
}
