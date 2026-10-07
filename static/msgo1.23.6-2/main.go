// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.6-2 command runs the go command from Microsoft build of Go 1.23.6-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.6-2@latest
//	$ msgo1.23.6-2 download
//
// And then use the msgo1.23.6-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.6-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.6-2", "msgo1.23.6-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f39f7aff-d083-4003-bb85-3ef0eb73e0c3/dd0df19ca59970ed0bfacd1126b44ae0/go1.23.6-20250211.6.linux-amd64.tar.gz", SHA256: "75911454278623a05181671709ae6845e0e3155a8e91b992ca3546851735c5a4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f39f7aff-d083-4003-bb85-3ef0eb73e0c3/29df14e0e3ae4d9e1f5b9ecf1ceebe79/go1.23.6-20250211.6.linux-armv6l.tar.gz", SHA256: "607d8436ce87ee4d0d378a93bf0178916d5150f28b8e67aca9563ee63f7eae7a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f39f7aff-d083-4003-bb85-3ef0eb73e0c3/9a4b954acbbb5537f6ef5bd695e44b75/go1.23.6-20250211.6.linux-arm64.tar.gz", SHA256: "45e7c01610774cbec5f56b51b466bb4c0c713fa463897d547076e303eaba3a00"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f39f7aff-d083-4003-bb85-3ef0eb73e0c3/ff8a54b03113fc4077f7119a448a693a/go1.23.6-20250211.6.windows-amd64.zip", SHA256: "ba38a09d33ae224ee28e563d1f3215a38888d13f7461980c1e4028e23364aacd"}
	default:
		return version.MicrosoftArchive{}
	}
}
