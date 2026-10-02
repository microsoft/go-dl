// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.24.3-1 command runs the go command from Microsoft build of Go 1.24.3-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.24.3-1@latest
//	$ msgo1.24.3-1 download
//
// And then use the msgo1.24.3-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.24.3-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.24.3-1", "msgo1.24.3-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/18982fcf-1155-40cc-bf0e-4b7564f22251/a6a4d462a5d8a7be72e3222c112d2101/go1.24.3-20250506.3.darwin-amd64.tar.gz", SHA256: "1f1d8e1537c4a8049e6f049a861c05d9709f7627178df36cb176872a2546fb7c"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/18982fcf-1155-40cc-bf0e-4b7564f22251/02acaf96922416da68a7b15764e56e3c/go1.24.3-20250506.3.darwin-arm64.tar.gz", SHA256: "330cb681a69cd32465b9e5381e1127eba3184a19d29e9216a9f36383f1bedff0"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/18982fcf-1155-40cc-bf0e-4b7564f22251/dfb7c8bd6b64c9fa9debd6efa724c609/go1.24.3-20250506.3.linux-amd64.tar.gz", SHA256: "d1e9e4465951816b556f648dfebc1cf20fdef4832d7d3f01f1ef35a259375286"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/18982fcf-1155-40cc-bf0e-4b7564f22251/3d149b264ff7d7718ee3d27a04518d70/go1.24.3-20250506.3.linux-armv6l.tar.gz", SHA256: "5c20018efd13d7c66f2e49ce78f4e35d6277cf61a95de7d8a46b7a91e2642b86"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/18982fcf-1155-40cc-bf0e-4b7564f22251/e4d43c57efd1a31cf8673bf43aa69412/go1.24.3-20250506.3.linux-arm64.tar.gz", SHA256: "6d4232da7ca7d2bd19e0b6b69524448de8e5480f49e5d39f7a5507cef682a9ea"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/18982fcf-1155-40cc-bf0e-4b7564f22251/db73ecc4766736dfc60201f64d96fb4c/go1.24.3-20250506.3.windows-amd64.zip", SHA256: "4369e3f6563060445ee19f83170755cc307d8bdb5ad56f42fe4bc93f4de5319d"}
	default:
		return version.MicrosoftArchive{}
	}
}
