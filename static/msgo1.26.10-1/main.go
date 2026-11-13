// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.10-1 command runs the go command from Microsoft build of Go 1.26.10-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.10-1@latest
//	$ msgo1.26.10-1 download
//
// And then use the msgo1.26.10-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.10-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.10-1", "msgo1.26.10-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/696ef58e-ac4f-43aa-a62d-8bb99b8f8658/c3d2ff784b2f0c16f1baddb4702b80a6/go1.26.10-20261020.3.darwin-amd64.tar.gz", SHA256: "75377fd9ccfc5f2d17240b00248602457e770ad9876611dfb23880c278e22892"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/696ef58e-ac4f-43aa-a62d-8bb99b8f8658/02d0857a871bffc1e05735107f7e85b9/go1.26.10-20261020.3.darwin-arm64.tar.gz", SHA256: "978f85e5fc9a3568aeb0a836027c497744fcee58dc8b22811c5ab24e28a22977"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/696ef58e-ac4f-43aa-a62d-8bb99b8f8658/f62a2fdda4300e98cfc1724952953dd7/go1.26.10-20261020.3.linux-amd64.tar.gz", SHA256: "595acf8ccb6506a43cb3b1d93f97d5222fd7f6ecd928eba77a751b7e1f6a6f05"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/696ef58e-ac4f-43aa-a62d-8bb99b8f8658/d70649a04b5119b9f82510cd0320977d/go1.26.10-20261020.3.linux-armv6l.tar.gz", SHA256: "23116ba22a652adb042df1c420f344d55d89bb5ced14fca91d294455b2cc3840"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/696ef58e-ac4f-43aa-a62d-8bb99b8f8658/17cb14a506610b525cd60a744a975871/go1.26.10-20261020.3.linux-arm64.tar.gz", SHA256: "fbb8c3db00dbd9f2d41a1bb29afc86311ae55dbe9ba9f38510cb119ec9e2f7b3"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/696ef58e-ac4f-43aa-a62d-8bb99b8f8658/86486b3a2945b90c162eeba5e9ec2091/go1.26.10-20261020.3.windows-amd64.zip", SHA256: "2e3663270c50c69e3875cab0069f682fa76d2168e8556760cbf0cb9674f3a819"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/696ef58e-ac4f-43aa-a62d-8bb99b8f8658/5e3a531221c826e93436038a6b0fcf24/go1.26.10-20261020.3.windows-arm64.zip", SHA256: "da180d3a9bf34f64496ce5d82be03d8f762b7b9ab900f33ac30f0da4bb0b1b65"}
	default:
		return version.MicrosoftArchive{}
	}
}
