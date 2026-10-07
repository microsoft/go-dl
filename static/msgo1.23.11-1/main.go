// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.11-1 command runs the go command from Microsoft build of Go 1.23.11-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.11-1@latest
//	$ msgo1.23.11-1 download
//
// And then use the msgo1.23.11-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.11-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.11-1", "msgo1.23.11-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4907e3a-3349-49ba-b03e-8289633429d6/e430267b97aed9ab7315212a18f4cae7/go1.23.11-20250708.6.linux-amd64.tar.gz", SHA256: "0dc3746aac6e59e30759035a726142c04e81f0aaa1f4f07fec789de11c0c11b5"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4907e3a-3349-49ba-b03e-8289633429d6/d7b92173158347c521757f0b65af4ad0/go1.23.11-20250708.6.linux-armv6l.tar.gz", SHA256: "665263ee2726c9bb05533bab491234e40eb71578a97b6276a5924c8eb82f39c2"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4907e3a-3349-49ba-b03e-8289633429d6/1c5d0d33b46c0539b23b625f832a325e/go1.23.11-20250708.6.linux-arm64.tar.gz", SHA256: "6359ec5ad5a641e73e6e256f9571d4a446eb90b8755ef4baacee9695ff6656f3"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/c4907e3a-3349-49ba-b03e-8289633429d6/ac3b28027629286235ecb9f126237b69/go1.23.11-20250708.6.windows-amd64.zip", SHA256: "ce399fd055117203da88cbb5fabcd572ae931667f8ef79184650548413a3320e"}
	default:
		return version.MicrosoftArchive{}
	}
}
