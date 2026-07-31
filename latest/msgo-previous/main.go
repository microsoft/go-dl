// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo-previous command runs the go command from Microsoft build of Go 1.26.8-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/latest/msgo-previous@latest
//	$ msgo-previous download
//
// And then use the msgo-previous command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.8-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo-previous", "msgo1.26.8-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/37bcb825-5f60-48f1-920d-8fc9799540c2/d478ae81714a98fc11f695c0855b68dc/go1.26.8-20260901.2.darwin-amd64.tar.gz", SHA256: "485c1cb540de9b07b5a3f121cc4ac2c55b075b8192eaa7642126eb7eb4b2640c"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/37bcb825-5f60-48f1-920d-8fc9799540c2/72b37c1b44bcbaaa68f54ad474ba96f5/go1.26.8-20260901.2.darwin-arm64.tar.gz", SHA256: "61ef73d207de274f528d88999f9433feb51bd60e5f2157f3853e86836f0d321c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/37bcb825-5f60-48f1-920d-8fc9799540c2/0b0bda7ee5028e64e9753e951c93ab45/go1.26.8-20260901.2.linux-amd64.tar.gz", SHA256: "2aeca690395bd5c8a6f00ac6eeed4350d204daa3952b9d7f3335c954dc6231b4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/37bcb825-5f60-48f1-920d-8fc9799540c2/14a52ef486ace8634fc139c4ec7d0a91/go1.26.8-20260901.2.linux-armv6l.tar.gz", SHA256: "265f8f1828d2f67e5531dd8de954a652143ee052b786f482167c176b8b913a4a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/37bcb825-5f60-48f1-920d-8fc9799540c2/f06d531995fc69df3a3b2c17a14c65b2/go1.26.8-20260901.2.linux-arm64.tar.gz", SHA256: "04d749b59f3c445ed51456cec95cfc0b7e42dbcfdd3f9f7b9e2e9ae0c34e6bec"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/37bcb825-5f60-48f1-920d-8fc9799540c2/500f53524f092cf19397602b9e0dc40b/go1.26.8-20260901.2.windows-amd64.zip", SHA256: "b628a8061d3f3d64641e9999fb336d2503ea2c5bc721f0b1b4aca1989fa962c1"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/37bcb825-5f60-48f1-920d-8fc9799540c2/cb34f522a98778e5a671d297d67c9c87/go1.26.8-20260901.2.windows-arm64.zip", SHA256: "5f4d719873e575af6f7654740c7d0237c5c78b9521e58b6d5cc78067f5605a61"}
	default:
		return version.MicrosoftArchive{}
	}
}
