// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.23.6-1 command runs the go command from Microsoft build of Go 1.23.6-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.23.6-1@latest
//	$ msgo1.23.6-1 download
//
// And then use the msgo1.23.6-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.23.6-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.23.6-1", "msgo1.23.6-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/eb97f523-9d90-4dad-8942-f0937ef28aec/6151a9136701940c39caa2b15de5351d/go1.23.6-20250205.5.linux-amd64.tar.gz", SHA256: "6b6c3cea5a0469f10ba90c568c298d7f83717c6feb240005bbc41041a659400c"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/eb97f523-9d90-4dad-8942-f0937ef28aec/37ac92549596cc44b9e793f82727899d/go1.23.6-20250205.5.linux-armv6l.tar.gz", SHA256: "22eba798e6ec106dc4519b824b40ff7730d830f996f68dab72fcb0e141930a3e"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/eb97f523-9d90-4dad-8942-f0937ef28aec/b0c08e6798232c6cfdab4d4b8aaf4a8a/go1.23.6-20250205.5.linux-arm64.tar.gz", SHA256: "959c3bed26f259f0af600fcba6c49643abde8823a775dec99fc9c68b2391bb3a"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/eb97f523-9d90-4dad-8942-f0937ef28aec/58fa4fa06a186727f16df1acdf2db0a6/go1.23.6-20250205.5.windows-amd64.zip", SHA256: "f61820f022dce524060bf61e2557138513795b1e1132e2237dad7d119bbf01e5"}
	default:
		return version.MicrosoftArchive{}
	}
}
