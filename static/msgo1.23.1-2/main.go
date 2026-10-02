// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.1-2 command runs the go command from Microsoft build of Go 1.23.1-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.1-2@latest
//	$ msgo1.23.1-2 download
//
// And then use the msgo1.23.1-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.1-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.1-2", "msgo1.23.1-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0b41dea8-d112-441d-8680-d46a7eaa903f/f8c6ecf7fa7dd100196cd1b3566d8ec9/go1.23.1-20240910.2.linux-amd64.tar.gz", SHA256: "9d96d73d914c253877df03e90289dec59524f476ef3eda392d6d023bac580f9b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0b41dea8-d112-441d-8680-d46a7eaa903f/b0c4fbc45fd3657c4dac50d627c12f96/go1.23.1-20240910.2.linux-armv6l.tar.gz", SHA256: "24cbf24f63d5b616cd68807bee9902efad98b0198cd5c5ca2857c3e63439cb1c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0b41dea8-d112-441d-8680-d46a7eaa903f/c3350a7610d97ab68410ba5fc6e46cd1/go1.23.1-20240910.2.linux-arm64.tar.gz", SHA256: "d88366194f835f1ce25737877c7a0950b2491712279b2b694bd3ef356f9dc28a"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/0b41dea8-d112-441d-8680-d46a7eaa903f/fb21febba63b3735cc73dff1fc81f7b1/go1.23.1-20240910.2.windows-amd64.zip", SHA256: "045e72af08cdc9ce363a1bf6992b6b20d7930a8a62574413615ea53bfd091de9"}
	default:
		return version.MicrosoftArchive{}
	}
}
