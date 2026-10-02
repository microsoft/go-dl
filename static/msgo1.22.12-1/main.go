// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.12-1 command runs the go command from Microsoft build of Go 1.22.12-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.12-1@latest
//	$ msgo1.22.12-1 download
//
// And then use the msgo1.22.12-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.12-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.12-1", "msgo1.22.12-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f93a0b32-76c4-4b69-b045-9b8d544e9e24/701869de8c95d74f94be6640c36318dd/go1.22.12-20250205.4.linux-amd64.tar.gz", SHA256: "c1415b09acf0856fa0f8e1ee0c0a3928e639c075d0c3c35c64a51b582cc96388"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f93a0b32-76c4-4b69-b045-9b8d544e9e24/033e00cfeae14a09f828d7b79ae6c8e3/go1.22.12-20250205.4.linux-armv6l.tar.gz", SHA256: "5c9741b4c27a5c771f2fed64a7f463b033b279f9a8961376e2d5107c23461dbd"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f93a0b32-76c4-4b69-b045-9b8d544e9e24/0cc6f023f342d02bce9ec4ad1a946c52/go1.22.12-20250205.4.linux-arm64.tar.gz", SHA256: "acb78fa7ee0b33a345a368964b802199415d9364c99ec0c314417c22d71dd5ff"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f93a0b32-76c4-4b69-b045-9b8d544e9e24/5a57d6bb689c5e1403b4275ed61b0a6d/go1.22.12-20250205.4.windows-amd64.zip", SHA256: "918c68b158082e1c0ed217923410895dadd346b08316c900f8d3241bf7580179"}
	default:
		return version.MicrosoftArchive{}
	}
}
