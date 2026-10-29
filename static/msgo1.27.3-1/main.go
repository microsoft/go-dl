// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.27.3-1 command runs the go command from Microsoft build of Go 1.27.3-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.27.3-1@latest
//	$ msgo1.27.3-1 download
//
// And then use the msgo1.27.3-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.27.3-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.27.3-1", "msgo1.27.3-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/6c530770-3a15-4be3-91b9-b6cf38ba108a/cdbf5d354df90319bfb1dbc3e56c83c3/go1.27.3-20261020.4.darwin-amd64.tar.gz", SHA256: "919b784671fcb803d02358b4bf8d77ab3a989a6dca16d113510be7ac9e4c4138"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/6c530770-3a15-4be3-91b9-b6cf38ba108a/6696ddd7174d00fec784b2065433e525/go1.27.3-20261020.4.darwin-arm64.tar.gz", SHA256: "d9bdb77a38f0cd725dde8560b0d10e66300eeae101307eb7a9cac47a6184f610"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/6c530770-3a15-4be3-91b9-b6cf38ba108a/757c3dc0ac97389868d125c160c3bf83/go1.27.3-20261020.4.linux-amd64.tar.gz", SHA256: "6fa66bc2b0d907ec4583b5704c51a8b07299ef2fee0e5b7e5737a0446e9e0145"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/6c530770-3a15-4be3-91b9-b6cf38ba108a/89432bb3e4ef10e159876c6b0efe52b7/go1.27.3-20261020.4.linux-armv6l.tar.gz", SHA256: "5a3d4b80f89efa95c3a6e9813db4be1c71ba9fed5a9e4c764ff0ee699a2502ad"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/6c530770-3a15-4be3-91b9-b6cf38ba108a/3254ac009a24e0eeef12ff4519ccdf9a/go1.27.3-20261020.4.linux-arm64.tar.gz", SHA256: "133b5020784b7b93d472d2a94107ca17a7d032dbdc7fbdb2fd6ffaa1166c86da"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/6c530770-3a15-4be3-91b9-b6cf38ba108a/37818a1fe24e8ea2dad4477dd33b4693/go1.27.3-20261020.4.windows-amd64.zip", SHA256: "9a764d71fe11a5f887e2bfd7d201b1c570dff678b5c271b6dff101efe728f431"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/6c530770-3a15-4be3-91b9-b6cf38ba108a/b25478818d834f6a1af813a0ae97f356/go1.27.3-20261020.4.windows-arm64.zip", SHA256: "ee91f0ac07cdfb118a5ce600bd4cf51ed4a9aa5d0345fe2659d3d3c5ab65dce3"}
	default:
		return version.MicrosoftArchive{}
	}
}
