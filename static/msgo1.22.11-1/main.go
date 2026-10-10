// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.11-1 command runs the go command from Microsoft build of Go 1.22.11-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.11-1@latest
//	$ msgo1.22.11-1 download
//
// And then use the msgo1.22.11-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.11-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.11-1", "msgo1.22.11-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3a4d1a71-22da-43dc-84b4-d602111dfd56/7612a82260cbe747b78d3b4d96e3e170/go1.22.11-20250116.5.linux-amd64.tar.gz", SHA256: "cece3f061e0f933629fb4d0c4c80da7cfdade147c6aca8b537704c18193f33e0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3a4d1a71-22da-43dc-84b4-d602111dfd56/cd02195adfd03ed8981c61ec95ca22b2/go1.22.11-20250116.5.linux-armv6l.tar.gz", SHA256: "238daa19886d7131befc709ab93002e17032ba19b4f9fd5abbb6ce0cadc45b5a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3a4d1a71-22da-43dc-84b4-d602111dfd56/78df4b350e7530e8f00f53f23a0fd3fb/go1.22.11-20250116.5.linux-arm64.tar.gz", SHA256: "0af4dae0f7f884701bb15c46d7bfd1727ce1166c43d18c79b50cb6ed3d731673"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/3a4d1a71-22da-43dc-84b4-d602111dfd56/dcece299f4d8f4d09c3ae05100d519af/go1.22.11-20250116.5.windows-amd64.zip", SHA256: "e101a10ee94a1f4964ef39946c1d1ad06192c002e33460816258b916318f4e21"}
	default:
		return version.MicrosoftArchive{}
	}
}
