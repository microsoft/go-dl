// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.6-1 command runs the go command from Microsoft build of Go 1.22.6-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.6-1@latest
//	$ msgo1.22.6-1 download
//
// And then use the msgo1.22.6-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.6-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.6-1", "msgo1.22.6-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f8a79935-ceae-4766-9ddc-c578a48189d3/1f280eb4eb6580d63fe44f02ec7e8b40/go1.22.6-20240806.4.linux-amd64.tar.gz", SHA256: "d49fda9363d5a4a0d0d7060447c25b197ebe57574d80892efef2e2768705a59d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f8a79935-ceae-4766-9ddc-c578a48189d3/9264aa4eadba6857b4ad0a3b9cb1d4de/go1.22.6-20240806.4.linux-armv6l.tar.gz", SHA256: "8f622cf1062095db07fb224471c75c8a8c8e7c138f29a53fd838a2a105bbbf39"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f8a79935-ceae-4766-9ddc-c578a48189d3/65716d7d39cac1f90895ba4d0048ca84/go1.22.6-20240806.4.linux-arm64.tar.gz", SHA256: "2e7164a1be4b90982568ac5bf94050a3384271bc0341c08783a668e74073f2fb"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f8a79935-ceae-4766-9ddc-c578a48189d3/07c3d1d1ae7a1188e57f89de31f8c7fb/go1.22.6-20240806.4.windows-amd64.zip", SHA256: "1067c9490e8914f783165665c8dab10f11018d3cc3df7f6e0baa10621e654626"}
	default:
		return version.MicrosoftArchive{}
	}
}
