// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.25.0-1 command runs the go command from Microsoft build of Go 1.25.0-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.25.0-1@latest
//	$ msgo1.25.0-1 download
//
// And then use the msgo1.25.0-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.25.0-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.25.0-1", "msgo1.25.0-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/74fad862-16c4-45b1-b1de-9a1cc83e5c84/8633018b4fa86b3d9eb80c4341a56306/go1.25.0-20250812.10.darwin-amd64.tar.gz", SHA256: "465eeb806ae54d5b8588d77163cc3e65911d76e5d335e71d7c5a3f2a0e9c1bbe"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/74fad862-16c4-45b1-b1de-9a1cc83e5c84/31ef4849d05fd2a5f9fcec8d705a47f4/go1.25.0-20250812.10.darwin-arm64.tar.gz", SHA256: "6c2ece846ca81b22981daa376555f17a6d135dcc67fc8ce5f332b6cb433c036d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/74fad862-16c4-45b1-b1de-9a1cc83e5c84/0889a3db50c0eedc86fe67ddef7e7498/go1.25.0-20250812.10.linux-amd64.tar.gz", SHA256: "ee4f04a8b1fa955bb78ff9302613e147ddece4bbe894cb942b89fcae19c4fa81"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/74fad862-16c4-45b1-b1de-9a1cc83e5c84/6aa2cecde0c9174687a87ef9c4ef9a29/go1.25.0-20250812.10.linux-armv6l.tar.gz", SHA256: "45dd1ef1a28ed5a240a46e70179086c771915cce50bf5fac57eeddcb3e5df0b7"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/74fad862-16c4-45b1-b1de-9a1cc83e5c84/01cf1a20a8ceeb111c26123bf9f25aa9/go1.25.0-20250812.10.linux-arm64.tar.gz", SHA256: "ebc1ea3f1262989d0e377f0a838f2ff0248eeed1224d8999c31b39b0f5e7a2f2"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/74fad862-16c4-45b1-b1de-9a1cc83e5c84/8c43e2416ae762346c7649a242f9bae7/go1.25.0-20250812.10.windows-amd64.zip", SHA256: "9358dede5fdd41b15eaa2a072efb8815cfaba90db3321e69eaddcd401fa717fb"}
	default:
		return version.MicrosoftArchive{}
	}
}
