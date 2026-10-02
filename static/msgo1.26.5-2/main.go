// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.5-2 command runs the go command from Microsoft build of Go 1.26.5-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.5-2@latest
//	$ msgo1.26.5-2 download
//
// And then use the msgo1.26.5-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.5-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.5-2", "msgo1.26.5-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a38f1e38-3cd4-48b2-b032-ca38793ec901/d45676061dfccb09b72f2579a6d8927f/go1.26.5-20260709.6.darwin-amd64.tar.gz", SHA256: "7a4f0bd26ccf1ce912bc509fccb473ee39bdd149b804fa581e59a72b7708e78e"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a38f1e38-3cd4-48b2-b032-ca38793ec901/f8da58db7007fa456d93c4f56bf6c337/go1.26.5-20260709.6.darwin-arm64.tar.gz", SHA256: "a6d2124035c18602926af78d1e3ae523abc613e36b2086da9b451991ce4288c4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a38f1e38-3cd4-48b2-b032-ca38793ec901/96d1f4b28a670f8e9b3259194508ae2f/go1.26.5-20260709.6.linux-amd64.tar.gz", SHA256: "df9f45167cf1cd825aa8555b76c5b24cb89dfbd771f48cc0007458236c414715"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a38f1e38-3cd4-48b2-b032-ca38793ec901/7245edbeea69e24e2c01980b6d2073ff/go1.26.5-20260709.6.linux-armv6l.tar.gz", SHA256: "d50dd47aa9ed76d2772b438d53a890a22a8e673e0167d3c953146e0827685e56"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a38f1e38-3cd4-48b2-b032-ca38793ec901/af17683a03acab86aec9ba37afbd5ff1/go1.26.5-20260709.6.linux-arm64.tar.gz", SHA256: "9990f8fcfefcbd15fb888dd777c7a72f1745889c9e94836ffe26f7e46e43328d"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a38f1e38-3cd4-48b2-b032-ca38793ec901/5b622136aa64e16797ca1cb765025d02/go1.26.5-20260709.6.windows-amd64.zip", SHA256: "4eea6203e774651c999da356bce298af2bfceeb06677c0f7f2a2f43da800b983"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a38f1e38-3cd4-48b2-b032-ca38793ec901/61a7f51dfc4218b1309c692cea862826/go1.26.5-20260709.6.windows-arm64.zip", SHA256: "875d5b4df1331d421f3c68aad7a0a8a77566f9bbc4f9413b4d4f3be99f75b2a4"}
	default:
		return version.MicrosoftArchive{}
	}
}
