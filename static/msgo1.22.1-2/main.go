// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.1-2 command runs the go command from Microsoft build of Go 1.22.1-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.1-2@latest
//	$ msgo1.22.1-2 download
//
// And then use the msgo1.22.1-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.1-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.1-2", "msgo1.22.1-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/7b419f7e-0325-482d-9a99-38e8627fdf2b/fb791add4fb7fad94c768bf0c2251c22/go1.22.1-20240320.6.linux-amd64.tar.gz", SHA256: "2928adebf5139a4ed9365145357b5d4cd72c4de68b773bea76a6013e27589229"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/7b419f7e-0325-482d-9a99-38e8627fdf2b/18246250934ec1763781cde6e9734e04/go1.22.1-20240320.6.linux-armv6l.tar.gz", SHA256: "9ab24e136fcf71d32f9024be641f6bb4dc95f7e81ddd29a3134aa431d32aea42"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/7b419f7e-0325-482d-9a99-38e8627fdf2b/f486d7b69ca41e949d60b437e4ba1097/go1.22.1-20240320.6.linux-arm64.tar.gz", SHA256: "b74664c8d5466ca3c9e3ed460c773608b532bc623643e4916c5a8bc790084bd8"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/7b419f7e-0325-482d-9a99-38e8627fdf2b/2309380e918fed233b86293d000d75f9/go1.22.1-20240320.6.windows-amd64.zip", SHA256: "d328a7402fe49514b81032f2bb407a1571c70c0d9b2f59406182595b44d87ad6"}
	default:
		return version.MicrosoftArchive{}
	}
}
