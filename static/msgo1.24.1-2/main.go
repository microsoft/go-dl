// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.1-2 command runs the go command from Microsoft build of Go 1.24.1-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.1-2@latest
//	$ msgo1.24.1-2 download
//
// And then use the msgo1.24.1-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.1-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.1-2", "msgo1.24.1-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/753e92d2-bc0b-409d-95dd-712fa177b254/dff08c3ab6867b1adeba2daa9ddf01a8/go1.24.1-20250306.2.darwin-amd64.tar.gz", SHA256: "9e0da2018ec0ef6801ffe7434697407c5d4a8f5c999f9252d3d80c2f8fa59916"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/753e92d2-bc0b-409d-95dd-712fa177b254/0b8f7b940f4805f984c3e7ec0e6c3ecf/go1.24.1-20250306.2.darwin-arm64.tar.gz", SHA256: "2e54f0450794c2be2c7350f409dc2a21af3cc296acf78f803e143cd9a1f877ff"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/753e92d2-bc0b-409d-95dd-712fa177b254/7f418d70e52f60abfec481381988492e/go1.24.1-20250306.2.linux-amd64.tar.gz", SHA256: "b0ca85ecc435a93e2ddb626dd9ef7fb6689700a0847b0392eb3e146345a8dea0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/753e92d2-bc0b-409d-95dd-712fa177b254/dc24acaff98372aa56698802c3f7ed90/go1.24.1-20250306.2.linux-armv6l.tar.gz", SHA256: "43d42297694b849f4266d6a91cfe50d3a59e958c2ffc1bb4527095bbeb6bd90b"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/753e92d2-bc0b-409d-95dd-712fa177b254/4919d5bd6bc9784de7dc9fad86966fd5/go1.24.1-20250306.2.linux-arm64.tar.gz", SHA256: "73b1befea457d5967632b2b0b93f8d2c0d899d5b6fbd1396c55d0a015292608b"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/753e92d2-bc0b-409d-95dd-712fa177b254/851968cdc9eaf855f6e9e412c92f9a8a/go1.24.1-20250306.2.windows-amd64.zip", SHA256: "07763e0fc6dabf06178c1b15d31660c236c1fd6310da59c1b112f617c8b0e9a1"}
	default:
		return version.MicrosoftArchive{}
	}
}
