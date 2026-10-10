// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.9-1 command runs the go command from Microsoft build of Go 1.25.9-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.9-1@latest
//	$ msgo1.25.9-1 download
//
// And then use the msgo1.25.9-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.9-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.9-1", "msgo1.25.9-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3689f0db-0362-41b9-bb70-26d5d9a92eb8/e481476d45849a40f071364ef7649a0e/go1.25.9-20260407.1.darwin-amd64.tar.gz", SHA256: "af835ac5478ae0971b7286c3ef937c901719df5a0a6d505a42b28fab7a7afcf9"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3689f0db-0362-41b9-bb70-26d5d9a92eb8/12792196575c95d3f870c648cc214e3b/go1.25.9-20260407.1.darwin-arm64.tar.gz", SHA256: "ff032c1c1fc4cd82777ddab571f6e3ca57ef212a715a7a58d3fea0bf1f57c9cc"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3689f0db-0362-41b9-bb70-26d5d9a92eb8/12b7f7ecc402cd82a14139a0e8898790/go1.25.9-20260407.1.linux-amd64.tar.gz", SHA256: "812813ecc7a26ad29eb44a73e9cbc7e705ee5ced37996098a6bce1b2ebd9f568"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3689f0db-0362-41b9-bb70-26d5d9a92eb8/25fc295aa239095d741192c22776e01f/go1.25.9-20260407.1.linux-armv6l.tar.gz", SHA256: "7a1788b52e1f5c70c3e77717af2caedae7ba9ff8f414d06f87d537acca78bb5b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3689f0db-0362-41b9-bb70-26d5d9a92eb8/1a71c9718afdc946a0e005b99bcecc57/go1.25.9-20260407.1.linux-arm64.tar.gz", SHA256: "f92d1b76f61402905eead2c651b7ccf11a97d93bb69b0e413489593870f8b220"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3689f0db-0362-41b9-bb70-26d5d9a92eb8/bc65eeb206909dc7bfd1c9bbe71c9570/go1.25.9-20260407.1.windows-amd64.zip", SHA256: "c1ed3128b63e1c366b03de1e40856edbbbf74d3049b74a2096768b023459478f"}
	default:
		return version.MicrosoftArchive{}
	}
}
