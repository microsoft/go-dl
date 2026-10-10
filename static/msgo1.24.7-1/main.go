// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.7-1 command runs the go command from Microsoft build of Go 1.24.7-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.7-1@latest
//	$ msgo1.24.7-1 download
//
// And then use the msgo1.24.7-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.7-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.7-1", "msgo1.24.7-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8bd8c979-3ada-4de9-ac35-867919b98621/0faa562f7504e64836b802e12ee0d7e7/go1.24.7-20250903.8.darwin-amd64.tar.gz", SHA256: "d7db16275adff3ecaea4db8d7677f8e67e8cea8d2091d7907dacec00d5de3360"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8bd8c979-3ada-4de9-ac35-867919b98621/e586577e8113b65d91bdf15d07f2b255/go1.24.7-20250903.8.darwin-arm64.tar.gz", SHA256: "73a2cf0e6df56dc56e406fb936adb8e7ca9a2fb00102ea52a60c1905930a8e38"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8bd8c979-3ada-4de9-ac35-867919b98621/b3ad9c947374c49a202d0a42ce3b3e2c/go1.24.7-20250903.8.linux-amd64.tar.gz", SHA256: "a283380559f8cfc7860bd4528ad7346904552cde9d0fb0255fac34e64f61a4cb"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8bd8c979-3ada-4de9-ac35-867919b98621/c44860f526766fe4da9fb09b30f97bea/go1.24.7-20250903.8.linux-armv6l.tar.gz", SHA256: "f02859a8be12e7f2dbb16c690a5f8dddaab226c4123f826adfbc6e4c53cf6925"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8bd8c979-3ada-4de9-ac35-867919b98621/69bd78eab380aee26d3ee302ce984061/go1.24.7-20250903.8.linux-arm64.tar.gz", SHA256: "27c5659c76f2e9fdc2ad3ea031442d969f45ec60d2156351256e6c281cb10bd8"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/8bd8c979-3ada-4de9-ac35-867919b98621/b7023d1e06caf2a38a96789ba6e16034/go1.24.7-20250903.8.windows-amd64.zip", SHA256: "16bb9f06b6c93b4895c737b8fa00e55d66e741589310a07fa5f76427965b765c"}
	default:
		return version.MicrosoftArchive{}
	}
}
