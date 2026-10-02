// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.2-2 command runs the go command from Microsoft build of Go 1.24.2-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.2-2@latest
//	$ msgo1.24.2-2 download
//
// And then use the msgo1.24.2-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.2-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.2-2", "msgo1.24.2-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/292b8087-c136-4c12-ba06-d54395a069b6/c589f6f4ab978406e2cefb87d0a55322/go1.24.2-20250417.2.darwin-amd64.tar.gz", SHA256: "067a0aeec1fecda5337bab9f06438daa0ae8add61e6d0168f7c31fafa30925ba"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/292b8087-c136-4c12-ba06-d54395a069b6/9beb34757334d08d673109037d499574/go1.24.2-20250417.2.darwin-arm64.tar.gz", SHA256: "1da3875148713b6f9d797ff76729a869c2ad79f68fce8ec140d5336334b0046d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/292b8087-c136-4c12-ba06-d54395a069b6/e7f5ac59079473e696c7cb1f40628590/go1.24.2-20250417.2.linux-amd64.tar.gz", SHA256: "7844f426e83fb66b3d109357214abcb7cce4ed87178179ef413efca3bce0303a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/292b8087-c136-4c12-ba06-d54395a069b6/a7f8c654153eb499611b1ec79d13af24/go1.24.2-20250417.2.linux-armv6l.tar.gz", SHA256: "3d9b518e07b4bf8db8c6c5454a9043d61b97de9cf95b57b9465f4fd4adc37002"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/292b8087-c136-4c12-ba06-d54395a069b6/1b7ade28110bad2222850f0994c2e5e5/go1.24.2-20250417.2.linux-arm64.tar.gz", SHA256: "26531b0ff1add9922b05b0dc8535dce5a9ca09517b6320d21d0dd77dc7f1f536"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/292b8087-c136-4c12-ba06-d54395a069b6/82db032eeee6e408803e7f91ece6c949/go1.24.2-20250417.2.windows-amd64.zip", SHA256: "d019a9b2b2bc5fb28af3ead545727a287cf4151c1922b475052f70559492a610"}
	default:
		return version.MicrosoftArchive{}
	}
}
