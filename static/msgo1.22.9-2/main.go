// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.9-2 command runs the go command from Microsoft build of Go 1.22.9-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.9-2@latest
//	$ msgo1.22.9-2 download
//
// And then use the msgo1.22.9-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.9-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.9-2", "msgo1.22.9-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a844e92-bf42-4d3a-9a4b-853f4f37b022/b671ceecb432858e57eeadb9f478c803/go1.22.9-20241202.5.linux-amd64.tar.gz", SHA256: "de930bd3e0816af77b1bbd199f0d09390ff24855a2bb96e2705d0b17e3954b7b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a844e92-bf42-4d3a-9a4b-853f4f37b022/fae8cdfc52adfdad823e2f780f5fe376/go1.22.9-20241202.5.linux-armv6l.tar.gz", SHA256: "aefbab3c9019112e1e934d96c5edb25477c38b3ef03ea1cdaf05250bcc2b907e"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a844e92-bf42-4d3a-9a4b-853f4f37b022/605bcc8a27de32b2c71cd5bb6dd0a296/go1.22.9-20241202.5.linux-arm64.tar.gz", SHA256: "a3409994020f0638aee7430749b973e77e07604e3659c2625ab2a880db2251cf"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a844e92-bf42-4d3a-9a4b-853f4f37b022/43347ed2335265a0d348a94be918644e/go1.22.9-20241202.5.windows-amd64.zip", SHA256: "264bb6cc4db288be0f35b7e9d7a62aa281b869d805e87510b877f3f44323e86c"}
	default:
		return version.MicrosoftArchive{}
	}
}
