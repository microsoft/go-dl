// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.2-1 command runs the go command from Microsoft build of Go 1.26.2-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.2-1@latest
//	$ msgo1.26.2-1 download
//
// And then use the msgo1.26.2-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.2-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.2-1", "msgo1.26.2-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a29a63c1-f82c-449f-baff-9c086f732ce6/0c2f246caa635f2270f7fb31e0bff430/go1.26.2-20260407.2.darwin-amd64.tar.gz", SHA256: "46291a8b87fb39dc499333d9657d590a7d38fefbd88d80da387f6d01660be446"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a29a63c1-f82c-449f-baff-9c086f732ce6/6c465361732842d8ffa7df6086354d80/go1.26.2-20260407.2.darwin-arm64.tar.gz", SHA256: "411616b1e261235c700220ff4aae3773d0d09a8697077400f67ebbd46b9e0003"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a29a63c1-f82c-449f-baff-9c086f732ce6/606a9119cca2b06097017987a66db36a/go1.26.2-20260407.2.linux-amd64.tar.gz", SHA256: "68bcd46d095165b37f4773450a8239ae5aa8d7e5be371eb69aea0510526ced5a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a29a63c1-f82c-449f-baff-9c086f732ce6/d6032b39256fa2819bfcc5681472e7c6/go1.26.2-20260407.2.linux-armv6l.tar.gz", SHA256: "ff4289aa9a6cafa64ce2dfbbe0c7ac8dd24fdc0948eebcbca606a7fcc9265414"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a29a63c1-f82c-449f-baff-9c086f732ce6/a5234a67f99d71512680ddea0881fe04/go1.26.2-20260407.2.linux-arm64.tar.gz", SHA256: "0274f18451d73bf234798b1d1212e60b59ed1cd86d1b982701df3df76c563571"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a29a63c1-f82c-449f-baff-9c086f732ce6/d415a62033e765d020f51145ff49da4f/go1.26.2-20260407.2.windows-amd64.zip", SHA256: "784849d8cd688ceaba0eba679959b50a1ff872ba73b5c6df7894c1cfe4f0b154"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/a29a63c1-f82c-449f-baff-9c086f732ce6/ed68084491395c15e0a3d613d28293a6/go1.26.2-20260407.2.windows-arm64.zip", SHA256: "60e5dc13ccb7186266dfad67d5a4a75c8537c510dd6c6fee2a4ad51f87c1647c"}
	default:
		return version.MicrosoftArchive{}
	}
}
