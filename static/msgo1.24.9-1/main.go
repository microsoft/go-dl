// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.9-1 command runs the go command from Microsoft build of Go 1.24.9-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.9-1@latest
//	$ msgo1.24.9-1 download
//
// And then use the msgo1.24.9-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.9-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.9-1", "msgo1.24.9-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/060fa5c7-8cff-4ac0-b196-1c8f100e304c/f9e60c8c3082b5a8c0a17b1dfeb6273c/go1.24.9-20251013.3.darwin-amd64.tar.gz", SHA256: "e92df6d4e7da5d773be5decde9e63b744eba0ea28a0bee66b914be50b6834bc9"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/060fa5c7-8cff-4ac0-b196-1c8f100e304c/565fd319399b210b0fa1174428f3f3d7/go1.24.9-20251013.3.darwin-arm64.tar.gz", SHA256: "962fc5dfba2bb5fdd99a833bd60436a2388c8ad8bdb50eaacafdea545a21d049"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/060fa5c7-8cff-4ac0-b196-1c8f100e304c/1037667824cd76d4d3c281ad31ad91dd/go1.24.9-20251013.3.linux-amd64.tar.gz", SHA256: "1d54fded463be66d6ccfd95379de86f11fa9da7686ed3cfc05d59e87d9248f50"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/060fa5c7-8cff-4ac0-b196-1c8f100e304c/0b7c8defc6d75acceb5546924c8c75f6/go1.24.9-20251013.3.linux-armv6l.tar.gz", SHA256: "182896278572b5e6ad5552cbc6c4dc9a4c6dddf99fed7955cb4dc65b41ed1bd5"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/060fa5c7-8cff-4ac0-b196-1c8f100e304c/41f8798eeaa790ccc77a85d7e06796f3/go1.24.9-20251013.3.linux-arm64.tar.gz", SHA256: "62cbf3028a86e5fa75032c7ce6249cc002aa7ea70151789c9e25b34c64f4e329"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/060fa5c7-8cff-4ac0-b196-1c8f100e304c/754040fa2fbfc5f2be7ee4f957f00d02/go1.24.9-20251013.3.windows-amd64.zip", SHA256: "9f1feb3c64cd7709ac3f99036d0e59b61f5c1378b711cc81710c53adefa9bfc8"}
	default:
		return version.MicrosoftArchive{}
	}
}
