// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.27.2-2 command runs the go command from Microsoft build of Go 1.27.2-2.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.27.2-2@latest
//	$ msgo1.27.2-2 download
//
// And then use the msgo1.27.2-2 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.27.2-2.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.27.2-2", "msgo1.27.2-2", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afc02740-5d8f-4ab9-b887-baaf80df63a0/6abf0f17334a68aeca0271719b5af651/go1.27.2-20261019.3.darwin-amd64.tar.gz", SHA256: "01850c1dc73e53dd79b025cd68ceb98b370f7d4a57fc3ee0ffc3aa2ade81bef5"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afc02740-5d8f-4ab9-b887-baaf80df63a0/a0568a6ddf28f12968cc3726f4622acf/go1.27.2-20261019.3.darwin-arm64.tar.gz", SHA256: "418bbc321d2ceeab9f968f68b9eb2bb7e628edafca3872f579e5ef8627f1a0b0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afc02740-5d8f-4ab9-b887-baaf80df63a0/eaea5a5719b481c50b14565a1f44fa80/go1.27.2-20261019.3.linux-amd64.tar.gz", SHA256: "37dc7dd8b0bbf881d80e25d7211319fe10b0bf86c29b321e9c7e3147110c5fc4"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afc02740-5d8f-4ab9-b887-baaf80df63a0/8c199577a1802acc75013f6fb308b1ab/go1.27.2-20261019.3.linux-armv6l.tar.gz", SHA256: "a6a2b97beedc5caa966f888657613fbe9b62d12c9fe42ab001c7db659c8f599d"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afc02740-5d8f-4ab9-b887-baaf80df63a0/8c5ca1ce3cec8f8c9e085d1e0c584270/go1.27.2-20261019.3.linux-arm64.tar.gz", SHA256: "1e87af32aa7fe1b55d481f1d710b7f4ea7c5561d799c8f0c9634626e80edb002"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afc02740-5d8f-4ab9-b887-baaf80df63a0/477bdd5e6654f7502164c550ce5d5706/go1.27.2-20261019.3.windows-amd64.zip", SHA256: "d5a980fecc3c38aa25acdfff94a65519d496c0aac88ae749ce2325a318ef4d19"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/afc02740-5d8f-4ab9-b887-baaf80df63a0/e8ebbb41883a499656ea6da960aeb2d4/go1.27.2-20261019.3.windows-arm64.zip", SHA256: "964d368bacb883d415175d4d7bdeb9a925eabd9563501c0e2dae213cd4f348af"}
	default:
		return version.MicrosoftArchive{}
	}
}
