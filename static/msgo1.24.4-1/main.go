// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.4-1 command runs the go command from Microsoft build of Go 1.24.4-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.4-1@latest
//	$ msgo1.24.4-1 download
//
// And then use the msgo1.24.4-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.4-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.4-1", "msgo1.24.4-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a9a3f50-f9f9-4a06-a5b6-30e4922ba01d/5de771e656d722c317af6a348e9e8c75/go1.24.4-20250605.5.darwin-amd64.tar.gz", SHA256: "b42f576f72267f7bb30cf995a6a2f1566d6d802d5ba4e84f32d46298084890d8"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a9a3f50-f9f9-4a06-a5b6-30e4922ba01d/5c56e8303a66ae0f31d7c2f9ed8f2316/go1.24.4-20250605.5.darwin-arm64.tar.gz", SHA256: "3d325c7f8e05dcdc8fe3371551315b897cfe98a3a93801bb6d199a3a8b87c914"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a9a3f50-f9f9-4a06-a5b6-30e4922ba01d/b20104f34e94fab7ff1bcde2c2ecd5f0/go1.24.4-20250605.5.linux-amd64.tar.gz", SHA256: "f21ddf7c3b029cf67615f6b2ca68e5aa83a50b2e2ea5952a5ea7bf9fa17e3358"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a9a3f50-f9f9-4a06-a5b6-30e4922ba01d/905f3fd3607fca8f9f8bf4f8dc67a1ff/go1.24.4-20250605.5.linux-armv6l.tar.gz", SHA256: "51cdaa2fbf6fce79a69c66a0ecfb9b837d56d634545230dde88f144248342769"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a9a3f50-f9f9-4a06-a5b6-30e4922ba01d/a33f0da754f734626e4e84a409e21ca2/go1.24.4-20250605.5.linux-arm64.tar.gz", SHA256: "44b769dd398f8fc5f381812c2eb323b4bc6cb11010d59efdba3244184e29b3fd"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1a9a3f50-f9f9-4a06-a5b6-30e4922ba01d/30134e16ba015eb9a0ec00a52896b703/go1.24.4-20250605.5.windows-amd64.zip", SHA256: "b41e6f47e4fb672f2f526c2415620e7702e5378139931fdc136c56cfc92f36f0"}
	default:
		return version.MicrosoftArchive{}
	}
}
