// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.13-1 command runs the go command from Microsoft build of Go 1.25.13-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.13-1@latest
//	$ msgo1.25.13-1 download
//
// And then use the msgo1.25.13-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.13-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.13-1", "msgo1.25.13-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c7ed68fb-651a-4f44-96f3-cd63ee69d54a/3a00144de0178c3777123bc6a349fe10/go1.25.13-20260813.2.darwin-amd64.tar.gz", SHA256: "4d6cb013d6e9bd7406559f028b0c53f85469647530746709fc4ce8ff47aee5db"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c7ed68fb-651a-4f44-96f3-cd63ee69d54a/32b9488a6f826261dd7d9c3b799f7d92/go1.25.13-20260813.2.darwin-arm64.tar.gz", SHA256: "5157f776f9d028f4506304cea7556ce0bdbdfa98a69436f6a42684bf27956e6d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c7ed68fb-651a-4f44-96f3-cd63ee69d54a/f3cf2fb6c22a9aa8496ef05277d78a30/go1.25.13-20260813.2.linux-amd64.tar.gz", SHA256: "d982b5b991aebd17cd56af1fd4693c43a5de40e45c39851532c4199b02e55bfd"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c7ed68fb-651a-4f44-96f3-cd63ee69d54a/812b0c964385655d8f97d18689104a86/go1.25.13-20260813.2.linux-armv6l.tar.gz", SHA256: "093ff15469666c243b404b77cd5694e647cb856fe31501ba590d72e2d8c02160"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c7ed68fb-651a-4f44-96f3-cd63ee69d54a/ac1b4c37936100c57838b7000a25166e/go1.25.13-20260813.2.linux-arm64.tar.gz", SHA256: "61f653c24cda975522fc178371e54a54af5d7b1a30baff761d4a75ad2c802362"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c7ed68fb-651a-4f44-96f3-cd63ee69d54a/6782f28f73c005627aa1687456fc4825/go1.25.13-20260813.2.windows-amd64.zip", SHA256: "3debdd347c33553cc295b8ab395423259cebf12424c7d8676f7c4cea9b8108cb"}
	default:
		return version.MicrosoftArchive{}
	}
}
