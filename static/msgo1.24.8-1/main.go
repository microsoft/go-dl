// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.8-1 command runs the go command from Microsoft build of Go 1.24.8-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.8-1@latest
//	$ msgo1.24.8-1 download
//
// And then use the msgo1.24.8-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.8-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.8-1", "msgo1.24.8-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f178e56d-5f18-42a8-8299-52dfd56c651d/177d3586bfbc548c68c4b1f9ef059071/go1.24.8-20251008.4.darwin-amd64.tar.gz", SHA256: "ab0f7ce1219db8ab74e47d1e500c66e6d90d64db97f68d122a4a5e6782a645d7"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f178e56d-5f18-42a8-8299-52dfd56c651d/31ed6bab6bb369acc8731b526d750199/go1.24.8-20251008.4.darwin-arm64.tar.gz", SHA256: "a81e61053e8a8904cab6bb6a3070636063e2c79cb7d939dac5562fd5a6581bd2"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f178e56d-5f18-42a8-8299-52dfd56c651d/cb3aebedd6536e10cc595d4ccdd21bfe/go1.24.8-20251008.4.linux-amd64.tar.gz", SHA256: "63a1c7ddfde7bd358e20db6499e92cd8619d97622308c558b6bbc59b0fd5b53f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f178e56d-5f18-42a8-8299-52dfd56c651d/60aa63bd901da488c6dedac8fc8dd9d1/go1.24.8-20251008.4.linux-armv6l.tar.gz", SHA256: "26435711515dcb146c054b5d6db418212f5d2fc947931d8858ba657307de3798"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f178e56d-5f18-42a8-8299-52dfd56c651d/7e7efb99c62cfb81401ba990f045c805/go1.24.8-20251008.4.linux-arm64.tar.gz", SHA256: "00bceeb34482dd726a26e7e2508e7958920a49252d9b0ef665c370e5de10feba"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f178e56d-5f18-42a8-8299-52dfd56c651d/80a29b0df088e55ed03dbc6d61b939a3/go1.24.8-20251008.4.windows-amd64.zip", SHA256: "fe34b84278520798498eb1c7c62e26f79e92e474909cf4971ba0c53f30b468cd"}
	default:
		return version.MicrosoftArchive{}
	}
}
