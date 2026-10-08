// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.22.2-1 command runs the go command from Microsoft build of Go 1.22.2-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.22.2-1@latest
//	$ msgo1.22.2-1 download
//
// And then use the msgo1.22.2-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.22.2-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.22.2-1", "msgo1.22.2-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d249bdfb-9343-4782-ab58-57b5cb22a825/489c44c8de2324fb6745ad5a80235d0c/go1.22.2-20240403.7.linux-amd64.tar.gz", SHA256: "f4780456c13e8e616ed4613515c9a462f45118cd5fd71a90eab5235aaa2428e0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d249bdfb-9343-4782-ab58-57b5cb22a825/741881d3e0b4e4b821145d5cb5593cff/go1.22.2-20240403.7.linux-armv6l.tar.gz", SHA256: "f3a2bb2da5c03e1a9043f85f57a5b5b696add2f5013e1120adffc211e12d9ee7"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d249bdfb-9343-4782-ab58-57b5cb22a825/afd73e5036831cfcedc3d29d5daa8eac/go1.22.2-20240403.7.linux-arm64.tar.gz", SHA256: "a753b5a9db4e0fc88e12b3b6d8f8d27fd6bd15edf222f9e96314a338d7c6d55d"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/d249bdfb-9343-4782-ab58-57b5cb22a825/bf13f4e7bb760ed79416ba603a8857b4/go1.22.2-20240403.7.windows-amd64.zip", SHA256: "6d059cb6be79673ec82d44dd48fa3afc3cd4dc20473bcb8ab8bd3a98a4771696"}
	default:
		return version.MicrosoftArchive{}
	}
}
