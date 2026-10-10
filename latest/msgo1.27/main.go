// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.27 command runs the go command from Microsoft build of Go 1.27.2-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/latest/msgo1.27@latest
//	$ msgo1.27 download
//
// And then use the msgo1.27 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.27.2-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.27", "msgo1.27.2-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/bbd872a5-5748-4062-ad96-a1f25f786c9c/8bdc1999dd384471ab011ce13c53fad9/go1.27.2-20261008.10.darwin-amd64.tar.gz", SHA256: "ddacf2f4ed1da8ccf285bd359c63d4b2badf0fbc54f5a7b8d272959da5d92aa1"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/bbd872a5-5748-4062-ad96-a1f25f786c9c/3239679110bf10ee37eb81caeb18dc0a/go1.27.2-20261008.10.darwin-arm64.tar.gz", SHA256: "c39f5085a8204fdcf341c86dd49ccb8f52bb22a5c95a514bf23deae34ae2070e"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/bbd872a5-5748-4062-ad96-a1f25f786c9c/edf21dedcef3a8c6c76dd2daba02fb5f/go1.27.2-20261008.10.linux-amd64.tar.gz", SHA256: "892a2ad3d66159b6537a11ea7bb5c68e925655ab5f48c811e8856f463ba8df17"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/bbd872a5-5748-4062-ad96-a1f25f786c9c/a3cb1c282fa7ccc3f007fbbba5747d3d/go1.27.2-20261008.10.linux-armv6l.tar.gz", SHA256: "bcc384687fce25dad862308b90c0e636eb56dbf7092c80c9bfa608f47868079a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/bbd872a5-5748-4062-ad96-a1f25f786c9c/5a7946725289cd376c7e6b2fc50b8dc1/go1.27.2-20261008.10.linux-arm64.tar.gz", SHA256: "30592532b9069b89bdd84ec642bc6d93f918e927869ac1beb1db64f114923b64"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/bbd872a5-5748-4062-ad96-a1f25f786c9c/f0f5b507f611b332aa0d88d7c6bc2826/go1.27.2-20261008.10.windows-amd64.zip", SHA256: "3b5ef22477a9c0ae7d845e4b0c7d1639440f3b4073c91a714e3025381ef79653"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/bbd872a5-5748-4062-ad96-a1f25f786c9c/992e0963e74cbbbff4031a384f15c4a0/go1.27.2-20261008.10.windows-arm64.zip", SHA256: "ba459ac1e677bf41bc55183f772e19cd4672fb8e3cf660d6493dbc480c967440"}
	default:
		return version.MicrosoftArchive{}
	}
}
