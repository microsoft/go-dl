// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.7-3 command runs the go command from Microsoft build of Go 1.22.7-3.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.7-3@latest
//	$ msgo1.22.7-3 download
//
// And then use the msgo1.22.7-3 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.7-3.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.7-3", "msgo1.22.7-3", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1d80df-e337-462a-a79b-6cb1272f437b/bb8d680f4146e18aa09b4c26784c70c2/go1.22.7-20240925.5.linux-amd64.tar.gz", SHA256: "c3a53478b4448e4c967384cd4bc343887a0fc3bb6022aac13c5045f23cdb5b1a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1d80df-e337-462a-a79b-6cb1272f437b/d345f4d4499c2380e1dd2d4025aa6938/go1.22.7-20240925.5.linux-armv6l.tar.gz", SHA256: "d7b3c53fbf709c463329e17aa4c6ba980f02b59c002cea9f9cdb04519aaccaa7"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1d80df-e337-462a-a79b-6cb1272f437b/e4c24b3ec392111f7d934cbed555d9e7/go1.22.7-20240925.5.linux-arm64.tar.gz", SHA256: "974768cbde7de2587435e22f306604482ed8583f4677cd04cef87bf013febbbe"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/9d1d80df-e337-462a-a79b-6cb1272f437b/bfe9b8525a5d219bf89d8ccc26953b94/go1.22.7-20240925.5.windows-amd64.zip", SHA256: "795decabbaeeb18706130ab75e80a6cc13108233dd5072fd9daa9bf9e7a3f7bc"}
	default:
		return version.MicrosoftArchive{}
	}
}
