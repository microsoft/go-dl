// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.27 command runs the go command from Microsoft build of Go 1.27.1-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/latest/msgo1.27@latest
//	$ msgo1.27 download
//
// And then use the msgo1.27 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.27.1-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.27", "msgo1.27.1-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f70fa462-cb6b-40dc-9240-b24e98dd2a80/4234262ceee01840d405f82a2b155372/go1.27.1-20260904.9.darwin-amd64.tar.gz", SHA256: "c3ae8412df8cb61fc24d22718b0339926fc326b5c10208b7102f17c9b90ba450"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f70fa462-cb6b-40dc-9240-b24e98dd2a80/5c2737db2c276b38168ef715257f8095/go1.27.1-20260904.9.darwin-arm64.tar.gz", SHA256: "12c68a572a0e8ef4ed540fa05c48b21f3d94f487539c3c7711018c371f7c8fd1"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f70fa462-cb6b-40dc-9240-b24e98dd2a80/7426f37648a3551e951452efc30fc8ed/go1.27.1-20260904.9.linux-amd64.tar.gz", SHA256: "fcaec387bae73ff6241a92b2ef45e2f86fe22583c4183f9008b236b168c3a464"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f70fa462-cb6b-40dc-9240-b24e98dd2a80/a7f423ad5eff8236a683c83aa87975e1/go1.27.1-20260904.9.linux-armv6l.tar.gz", SHA256: "750f236ab86e87bbc8286bd7d63cb66abf1b8ec30a090820b62e0a6d7b850bc3"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f70fa462-cb6b-40dc-9240-b24e98dd2a80/eedbd98072794baba3bfd34c6cb0e2d0/go1.27.1-20260904.9.linux-arm64.tar.gz", SHA256: "5ec3b63260359f3b636c6892c87bafed494ad7f431b762c6ede109c7029dddcc"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f70fa462-cb6b-40dc-9240-b24e98dd2a80/5320098654b2d4e814c4342e972c25f5/go1.27.1-20260904.9.windows-amd64.zip", SHA256: "22a8bc3d8af5e0fe0515c9ba7bb6324aa9d743c739c8fc670028eca0d9ef1c3b"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/f70fa462-cb6b-40dc-9240-b24e98dd2a80/cee88e982d7fc8c4d39ca9fa55e75c79/go1.27.1-20260904.9.windows-arm64.zip", SHA256: "0ad7f7cfb29f9b7993fb9119159b3e89de0c05814cda7cc9cb6ccb534b34700a"}
	default:
		return version.MicrosoftArchive{}
	}
}
