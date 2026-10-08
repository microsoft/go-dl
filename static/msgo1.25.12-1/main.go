// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.12-1 command runs the go command from Microsoft build of Go 1.25.12-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.12-1@latest
//	$ msgo1.25.12-1 download
//
// And then use the msgo1.25.12-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.12-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.12-1", "msgo1.25.12-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0839b8e8-6fb1-4b0e-af81-831c2e243817/b3678c5d5452e778b442538430900163/go1.25.12-20260707.5.darwin-amd64.tar.gz", SHA256: "7bc2c51b29d07b0c64d17d8c5a008907da853c4f941773ac539ab7ae7ddabd8a"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0839b8e8-6fb1-4b0e-af81-831c2e243817/b229f7e97f982e8e77247989687a5f01/go1.25.12-20260707.5.darwin-arm64.tar.gz", SHA256: "96294b2128ec556b4737e36dca4defbaea938bd6e871cab19705f92fbace4d2c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0839b8e8-6fb1-4b0e-af81-831c2e243817/43a71737ac471c200bfbc00fc4cb9f33/go1.25.12-20260707.5.linux-amd64.tar.gz", SHA256: "c373e19623c7756a5d1f75fce0ee53b9651fe2244d8f78fd4fb6611dcba86915"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0839b8e8-6fb1-4b0e-af81-831c2e243817/d03d9163946c2906768c2b000d26d8de/go1.25.12-20260707.5.linux-armv6l.tar.gz", SHA256: "9139126272fbcb81aa4be40b8c1ef68f3d75a5e41d38ebff9d35dbdc21a11260"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0839b8e8-6fb1-4b0e-af81-831c2e243817/e3ddd8401aab3eb012c17814dbebe7a7/go1.25.12-20260707.5.linux-arm64.tar.gz", SHA256: "5458d27bdc7b4714ffb63a963cf1a9c16f8dbd02108a718fb0f694179a1b100c"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0839b8e8-6fb1-4b0e-af81-831c2e243817/d4794c9c2f17ea0946e0b1c359bb2254/go1.25.12-20260707.5.windows-amd64.zip", SHA256: "e7b0cba30e156ff873164c96d54fb38860b80bbc7624bea51334657f6a8bf1b1"}
	default:
		return version.MicrosoftArchive{}
	}
}
