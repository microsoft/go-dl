// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.2-1 command runs the go command from Microsoft build of Go 1.23.2-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.2-1@latest
//	$ msgo1.23.2-1 download
//
// And then use the msgo1.23.2-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.2-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.2-1", "msgo1.23.2-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e6c8c6fa-e120-4af2-9866-3c0e5a161bc6/2ada0cfb430890c2db4b50477559ecc0/go1.23.2-20241001.7.linux-amd64.tar.gz", SHA256: "af972fade8be4fe3a780c0c6bd7294ebe994fa519066caf475dcbf7f59090cfa"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e6c8c6fa-e120-4af2-9866-3c0e5a161bc6/fa532edcd5d2e9182478e303f2fd1781/go1.23.2-20241001.7.linux-armv6l.tar.gz", SHA256: "319f8bfb575edff89223500ea4dd27b74ef02ddb4491754a9678b43831f91072"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e6c8c6fa-e120-4af2-9866-3c0e5a161bc6/127727d7e594dcef3e2c645dbdbbfa73/go1.23.2-20241001.7.linux-arm64.tar.gz", SHA256: "3710cfffe7179a06549e4acf3131ab10a7eb039636fb38f7ceaaae8a8858cc7c"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/e6c8c6fa-e120-4af2-9866-3c0e5a161bc6/6ab5a6efeb8e1ddf484b13fff7a69195/go1.23.2-20241001.7.windows-amd64.zip", SHA256: "03a59113b0d9adb1659fb0f0cfa08222e0adbb267fe57c0370054587f59f6d88"}
	default:
		return version.MicrosoftArchive{}
	}
}
