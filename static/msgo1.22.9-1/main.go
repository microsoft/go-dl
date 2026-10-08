// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.9-1 command runs the go command from Microsoft build of Go 1.22.9-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.9-1@latest
//	$ msgo1.22.9-1 download
//
// And then use the msgo1.22.9-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.9-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.9-1", "msgo1.22.9-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/65739ec9-36f1-4d7e-909a-77a2bdcb9451/4eb8f95ad73e5decb0b3c52c2529b85e/go1.22.9-20241107.6.linux-amd64.tar.gz", SHA256: "50259081a4ecdcd271ad878a07de6271f24e0e0f8c88e8b917f3ed3484d05da4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/65739ec9-36f1-4d7e-909a-77a2bdcb9451/933d421187961ae10db29746d7fc1181/go1.22.9-20241107.6.linux-armv6l.tar.gz", SHA256: "6110832dd26f9579eab0fc056ee8c8b37edce1ad2fa1f644e5d70b1403d57231"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/65739ec9-36f1-4d7e-909a-77a2bdcb9451/32e6aa0a9d1ff03d2dd6f12d6142c8a4/go1.22.9-20241107.6.linux-arm64.tar.gz", SHA256: "e7bb171aff384516c94c0ad367faacb02a92c9918c1c9f44d48e3b2c71a4327b"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/65739ec9-36f1-4d7e-909a-77a2bdcb9451/6592b1d1d3a93a01cfd09ef5cc182eaf/go1.22.9-20241107.6.windows-amd64.zip", SHA256: "6863225222fae97ce63905d7ee6fbdccc0bf3222f8afc72dd6c2f47998d21d00"}
	default:
		return version.MicrosoftArchive{}
	}
}
