// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.7-1 command runs the go command from Microsoft build of Go 1.22.7-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.7-1@latest
//	$ msgo1.22.7-1 download
//
// And then use the msgo1.22.7-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.7-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.7-1", "msgo1.22.7-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b07458ed-d49f-4b93-8be3-0cf6708967ce/450ce587ac0c88b59e7321655d0cdcb6/go1.22.7-20240905.3.linux-amd64.tar.gz", SHA256: "b2a2b676844ed7586fad1df8dda96527da3e0bf8de2a25916c2394633f6f914d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b07458ed-d49f-4b93-8be3-0cf6708967ce/5c5c3894dfe2f12d4eb6c4deec83a769/go1.22.7-20240905.3.linux-armv6l.tar.gz", SHA256: "829fb8e6050f7d6a33c1b9879a5ff51535212399f6962bc72652b5743dd827c0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b07458ed-d49f-4b93-8be3-0cf6708967ce/89d7a5e19d7026d8f5f979ca34596b29/go1.22.7-20240905.3.linux-arm64.tar.gz", SHA256: "071530b11bc734e4e35a336ad59f643fb4ac8050b0b366aaeffdf4408c8ded60"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b07458ed-d49f-4b93-8be3-0cf6708967ce/5e6ca8fd563e2d8b56ea063cef0a0768/go1.22.7-20240905.3.windows-amd64.zip", SHA256: "93d7e398701a0b898cf4904b157acabd8bb4028bf7e6f10ac54e5bce8cc9ccca"}
	default:
		return version.MicrosoftArchive{}
	}
}
