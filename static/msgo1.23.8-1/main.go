// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.8-1 command runs the go command from Microsoft build of Go 1.23.8-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.8-1@latest
//	$ msgo1.23.8-1 download
//
// And then use the msgo1.23.8-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.8-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.8-1", "msgo1.23.8-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4e8cc2b3-53c0-4b0e-9fb6-5d62d5cca90c/1fc733d4dd0e63dedb1b516fb220ad28/go1.23.8-20250401.2.linux-amd64.tar.gz", SHA256: "e08de24661de3351bf216ab3f875377391a182ca0b1b3ee64d91fb64f8dcf10f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4e8cc2b3-53c0-4b0e-9fb6-5d62d5cca90c/60e58368cf39dad18b14257a01f41cd9/go1.23.8-20250401.2.linux-armv6l.tar.gz", SHA256: "4ac8b9ca0eb76c7fa17ed9f4683dfcd908eaeec5e2aca018efad69bab0d48e4a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4e8cc2b3-53c0-4b0e-9fb6-5d62d5cca90c/5924e75cdbf05090cb890932553ecfc6/go1.23.8-20250401.2.linux-arm64.tar.gz", SHA256: "6367764f5e2886e6e3a32c986d7d426f500423cb4c06ea80d844bd4013b96e62"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/4e8cc2b3-53c0-4b0e-9fb6-5d62d5cca90c/3b67efc8d2bd39a4c18db45267512233/go1.23.8-20250401.2.windows-amd64.zip", SHA256: "2a3340a260dcc389213d23985c7477646561df67bb056bb8f668388f47253532"}
	default:
		return version.MicrosoftArchive{}
	}
}
