// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.3-2 command runs the go command from Microsoft build of Go 1.23.3-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.3-2@latest
//	$ msgo1.23.3-2 download
//
// And then use the msgo1.23.3-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.3-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.3-2", "msgo1.23.3-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e600fd19-c2ed-459f-9aa3-ebf561226613/8989686c7dd6a0e51452200378025cb6/go1.23.3-20241202.3.linux-amd64.tar.gz", SHA256: "24e0df0bcd9d9442a8d6c77bd88a29bbad99de6eeb4d142b08a28a0d396041dd"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e600fd19-c2ed-459f-9aa3-ebf561226613/0a38b348e4bf867d789f25f709afa555/go1.23.3-20241202.3.linux-armv6l.tar.gz", SHA256: "90fca16ca309d75897032bbdb3a7abf4983122d764081a222113ae3d862d2e3d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e600fd19-c2ed-459f-9aa3-ebf561226613/04dd0e5c4431e6f0da2997bdcbba0f5e/go1.23.3-20241202.3.linux-arm64.tar.gz", SHA256: "47764247cbedbcaba25b7b37dfa957de33956e595f114ad34eacf25e1cb91abd"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e600fd19-c2ed-459f-9aa3-ebf561226613/addd15bc2c5768be4d7243b479991460/go1.23.3-20241202.3.windows-amd64.zip", SHA256: "68730ac4e2b3c232b4aa016f95ea56d8b2dc56ee4d5b28529e9286c951f259b7"}
	default:
		return version.MicrosoftArchive{}
	}
}
