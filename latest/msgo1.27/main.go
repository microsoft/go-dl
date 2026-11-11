// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.27 command runs the go command from Microsoft build of Go 1.27.4-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/latest/msgo1.27@latest
//	$ msgo1.27 download
//
// And then use the msgo1.27 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.27.4-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.27", "msgo1.27.4-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b32709ee-0545-4d3f-913f-bf51b626be52/43ccd825778df800bdfe4e643b21c993/go1.27.4-20261110.2.darwin-amd64.tar.gz", SHA256: "67a3de1ff011455e7b2695b9d9784660d2ced50d820d4a4651a67d1307f44ff8"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b32709ee-0545-4d3f-913f-bf51b626be52/78d35f2d102e66366832eced0595c03c/go1.27.4-20261110.2.darwin-arm64.tar.gz", SHA256: "795f9410a78103483f5442d85616ce36a6137527f9dd18e3d695269e44029e71"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b32709ee-0545-4d3f-913f-bf51b626be52/33fbba165da669d1cb19d3799e623805/go1.27.4-20261110.2.linux-amd64.tar.gz", SHA256: "423eee29e71d3da6b8d59f7fa080ca21f28015857af58f3da5349c264a5f3f2f"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b32709ee-0545-4d3f-913f-bf51b626be52/b06b03f79d14ea3068c9c533713c6999/go1.27.4-20261110.2.linux-armv6l.tar.gz", SHA256: "d68af464b9e935b933218de4262e98262b2d481112b82ac8e4a3b635e49bccaa"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b32709ee-0545-4d3f-913f-bf51b626be52/a5caf67eaca4942d973c08bb1e54fa40/go1.27.4-20261110.2.linux-arm64.tar.gz", SHA256: "8bd381cdc3a3dc932117a3c32ac49ac8787d8a32ca4fba13ecbac890265afd3a"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b32709ee-0545-4d3f-913f-bf51b626be52/047873dd94d8bdb8e5298a6f465c1d82/go1.27.4-20261110.2.windows-amd64.zip", SHA256: "4f8e74152881060f4f5173c27fb43d1c94287cae305c54548fcccb6e2d9f628f"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b32709ee-0545-4d3f-913f-bf51b626be52/287ae53a80f46f3b2844dfbacb58ff89/go1.27.4-20261110.2.windows-arm64.zip", SHA256: "6b5c3176b7444d0b00de7b2490b487103bf0d970bfc4fc95f0c95dfb030f95d6"}
	default:
		return version.MicrosoftArchive{}
	}
}
