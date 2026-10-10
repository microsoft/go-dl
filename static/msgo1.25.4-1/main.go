// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.4-1 command runs the go command from Microsoft build of Go 1.25.4-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.4-1@latest
//	$ msgo1.25.4-1 download
//
// And then use the msgo1.25.4-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.4-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.4-1", "msgo1.25.4-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/db7adc7a-2fc2-4613-88de-a0a0b59cc636/7d5c3c0feb991de34fb17d8f8ba119eb/go1.25.4-20251105.5.darwin-amd64.tar.gz", SHA256: "7251e2dfdb5a3989fa513bae24be506610280a9c929cbc658d0e5583aac65660"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/db7adc7a-2fc2-4613-88de-a0a0b59cc636/dee7a3f48d7636c02d3948c700735f49/go1.25.4-20251105.5.darwin-arm64.tar.gz", SHA256: "d9231c639b90e84cc3ceed736b77bdcc7e55a2c3f97d0b3b3d9729285839a5df"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/db7adc7a-2fc2-4613-88de-a0a0b59cc636/76553284b63a7a29a54db5dc4160e2be/go1.25.4-20251105.5.linux-amd64.tar.gz", SHA256: "c728dc46225d7a73e90414097015eeecce8fa63643d0c28be6a98d511e04e60d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/db7adc7a-2fc2-4613-88de-a0a0b59cc636/5eb61cb895a92c2ec2185fe30a9a5158/go1.25.4-20251105.5.linux-armv6l.tar.gz", SHA256: "eb9675861ecd6f83b9de294bd2f9702bb3a0dce264328314a63b5b04ca33d6b2"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/db7adc7a-2fc2-4613-88de-a0a0b59cc636/0a27d8f2dddb2ba580ff977ae42a6a8b/go1.25.4-20251105.5.linux-arm64.tar.gz", SHA256: "dcba9ab5fb79f25d4da993dc9b8116b7c0f9e10fc243623637e84b4c8d6d03d0"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/db7adc7a-2fc2-4613-88de-a0a0b59cc636/7bf019389baf2e28bb046bc8356f0b08/go1.25.4-20251105.5.windows-amd64.zip", SHA256: "2b90deb4dc0dd990d812fc88fbb51c6c33e346371be30bdb9e864e2c810e8bfc"}
	default:
		return version.MicrosoftArchive{}
	}
}
