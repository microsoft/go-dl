// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.9-1 command runs the go command from Microsoft build of Go 1.26.9-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.9-1@latest
//	$ msgo1.26.9-1 download
//
// And then use the msgo1.26.9-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.9-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.9-1", "msgo1.26.9-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b1abaa90-2e57-464f-8c7a-70edc0b126ce/17cfc771cdd2e01f7dee40cd1e570f52/go1.26.9-20261008.11.darwin-amd64.tar.gz", SHA256: "15694aa8ce1c554ae9f176344bff7bb6cec29f11165ceb93db3da97db9626e68"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b1abaa90-2e57-464f-8c7a-70edc0b126ce/dd96766d7e7010fa228e31b55dee4e16/go1.26.9-20261008.11.darwin-arm64.tar.gz", SHA256: "276c92d162a9c967e718ef693728f76c4a7518c4f516e1bc51099c95f5d9515e"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b1abaa90-2e57-464f-8c7a-70edc0b126ce/726a4603a7c498e5444fbeec66050176/go1.26.9-20261008.11.linux-amd64.tar.gz", SHA256: "4bec2b6e49d6b51a16010db3e8099155af3157545a7725d073e328e15173a6fa"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b1abaa90-2e57-464f-8c7a-70edc0b126ce/66d2de769db850d9b8e300627ebccc34/go1.26.9-20261008.11.linux-armv6l.tar.gz", SHA256: "1dca695721efb64a3e75cf2d3ae3cfa335cca697bf33b82bb6ce36ec17324a70"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b1abaa90-2e57-464f-8c7a-70edc0b126ce/4b90ac8a5d9b89a88d2e7e9fee2762ac/go1.26.9-20261008.11.linux-arm64.tar.gz", SHA256: "9af7a777e26da9a1f032fd8efaff838d8971085c47900c3bea3ecd4d83cceb9a"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b1abaa90-2e57-464f-8c7a-70edc0b126ce/6161cc70a785bb2874d7f387fcef9377/go1.26.9-20261008.11.windows-amd64.zip", SHA256: "cfa4f140bdcabd1cb7574b64fa512594b836b4dda9bfe7730863bac820f49480"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/b1abaa90-2e57-464f-8c7a-70edc0b126ce/dbe76937cabf65bc1a6b04fb953e5bd9/go1.26.9-20261008.11.windows-arm64.zip", SHA256: "71a8595168801548ae3ef5b34b52f61c98eed45f5075928a84ae9c77e09c88c4"}
	default:
		return version.MicrosoftArchive{}
	}
}
