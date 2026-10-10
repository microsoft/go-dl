// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The msgo1.26.0-1 command runs the go command from Microsoft build of Go 1.26.0-1.
//
// To install, run:
//
//	$ go install github.com/microsoft/go-dl/static/msgo1.26.0-1@latest
//	$ msgo1.26.0-1 download
//
// And then use the msgo1.26.0-1 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/microsoft/go/releases/tag/v1.26.0-1.
//
// File bugs at https://github.com/microsoft/go/issues/new.
package main

import (
	"runtime"

	"github.com/microsoft/go-dl/internal/version"
)

func main() {
	version.RunMicrosoft("msgo1.26.0-1", "msgo1.26.0-1", microsoftArchive())
}

func microsoftArchive() version.MicrosoftArchive {
	switch {
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1213b9cd-fed8-4540-99ef-5bcb4232a5ce/e95f52e24345d0bb91e31ee96bb83088/go1.26.0-20260210.6.darwin-amd64.tar.gz", SHA256: "d41a8d6eb580f150b4dd1080bc22b61e8b01c581800bc56d828ef412e903a8de"}
	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1213b9cd-fed8-4540-99ef-5bcb4232a5ce/5d6f05df5cb22c41b5470faa5ec93dbf/go1.26.0-20260210.6.darwin-arm64.tar.gz", SHA256: "91ee13d9fdcc25f29bfc974f76e1f95f03f93a35b738df9fbf3d949d969a6d10"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1213b9cd-fed8-4540-99ef-5bcb4232a5ce/2b419b1c79f298267a6f714d5d7a7f86/go1.26.0-20260210.6.linux-amd64.tar.gz", SHA256: "54490264014e939bea4df08db83139c016bf7860db6ccbe996c6106b865a286a"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1213b9cd-fed8-4540-99ef-5bcb4232a5ce/28722388bc71307dfd0b76ddb52a7ff8/go1.26.0-20260210.6.linux-armv6l.tar.gz", SHA256: "7b36e42a578deec098b2b77bee9598992090d0848d582a2985639529fc38408e"}
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1213b9cd-fed8-4540-99ef-5bcb4232a5ce/1ccafce81600e7fbacde66bf901ace4f/go1.26.0-20260210.6.linux-arm64.tar.gz", SHA256: "889dda151f8be8fb13b9beb164ba49b320c091bfb665e64b04280f35bd7df0c9"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "amd64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1213b9cd-fed8-4540-99ef-5bcb4232a5ce/95c159cc4cb803dd461e18c47c53d18f/go1.26.0-20260210.6.windows-amd64.zip", SHA256: "20d5b2e0811bac8670abdb6ea3425925c39060988f0ed103a7b11a0c3aa11c93"}
	case runtime.GOOS == "windows" && runtime.GOARCH == "arm64":
		return version.MicrosoftArchive{URL: "https://download.visualstudio.microsoft.com/download/pr/1213b9cd-fed8-4540-99ef-5bcb4232a5ce/34d4db8ba8f9e3e91ed3ba86b5deab4f/go1.26.0-20260210.6.windows-arm64.zip", SHA256: "4781aa99a18fb9a7e47053dbdc91a8bcc9c56ecb79bb85a3b243c0c1e09bca4d"}
	default:
		return version.MicrosoftArchive{}
	}
}
