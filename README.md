# github.com/microsoft/go-dl

[![Go Reference](https://pkg.go.dev/badge/github.com/microsoft/go-dl.svg)](https://pkg.go.dev/github.com/microsoft/go-dl)

Use this project to install and run a specific version of the [Microsoft build of Go](https://github.com/microsoft/go) alongside an existing Go installation.
It follows the wrapper and installation model established by [golang.org/dl](https://pkg.go.dev/golang.org/dl), using release metadata and downloads for the Microsoft build of Go.

## Install

You need an existing Go installation to install this wrapper.
The wrapper then downloads and runs the selected Microsoft build of Go.
The directory where `go install` places executables must be on your `PATH`; check your Go environment's `GOBIN` and `GOPATH` settings if the wrapper command is not found.

Most users should install `msgo`, which follows the current Go release series:

```console
go install github.com/microsoft/go-dl/latest/msgo@latest
msgo download
```

Use the installed wrapper as you would use the regular `go` command:

```console
msgo version
msgo test ./...
```

You may want to use an alternative package path depending on your scenario.
See [What version should I use?](https://github.com/microsoft/go/blob/microsoft/main/eng/doc/MigrationGuide.md#what-version-should-i-use) for help choosing, then use one of the alternatives below.

<!-- BEGIN GENERATED VERSION EXAMPLES: DO NOT EDIT -->

### Stable release tracking

`msgo` currently downloads Microsoft build of Go 1.27.1-2.

`msgo-previous` follows the previous stable Go release series and currently downloads Microsoft build of Go 1.26.8-1.

```console
go install github.com/microsoft/go-dl/latest/msgo-previous@latest
```

### Release series tracking

`msgo1.27` stays on Go 1.27 and currently downloads Microsoft build of Go 1.27.1-2.

```console
go install github.com/microsoft/go-dl/latest/msgo1.27@latest
```

`msgo1.26` stays on Go 1.26 and currently downloads Microsoft build of Go 1.26.8-1.

```console
go install github.com/microsoft/go-dl/latest/msgo1.26@latest
```

### Exact release

Install a wrapper from `static` to select one exact Microsoft build of Go:

```console
go install github.com/microsoft/go-dl/static/msgo1.27.1-2@latest
```

## Update a wrapper

A wrapper installed with `go install` does not update itself when its package points to a newer Microsoft build of Go.
Run the same `go install ...@latest` command again to replace the installed wrapper.

For a project-managed tool, you can record the wrapper package as a tool dependency in `go.mod` and run it without an `@version` suffix:

```console
go run github.com/microsoft/go-dl/latest/msgo1.27 download
```

Without an `@version` suffix, Go uses the go-dl module version selected by `go.mod`.
Dependency update tools such as Dependabot can then propose go-dl updates through the usual `go.mod` and `go.sum` changes.
If your project's CI uses the wrapper to build and test, the dependency-update pull request exercises the new Microsoft build of Go before merge.
This can catch toolset update breaks and prevent them from affecting the target branch.

<!-- END GENERATED VERSION EXAMPLES -->

## Support

See [SUPPORT.md](SUPPORT.md) for more information about reporting bugs, requesting features, and asking questions.

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us the rights to use your contribution.
For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide a CLA and decorate the PR appropriately (e.g., status check, comment).
Simply follow the instructions provided by the bot.
You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

Please read the [internal generation module's README](./internal/gen/README.md) for more information about contributing to this project.

### Tests and CI

Run the downloader and generator tests separately:

```console
go test ./...
(cd internal/gen && go test ./...)
```

Live tests download and run actual SDKs and are disabled by default.
To test selected wrappers locally, pass their package paths as a space-separated list:

```console
GO_DL_LIVE_TEST=1 GO_DL_LIVE_PACKAGES="./latest/msgo ./latest/msgo-previous" go test -v ./internal/integration
```

The [CI workflow](./.github/workflows/ci.yml) tests each directly changed wrapper in a separate live-test job at both the minimum Go version declared in [go.mod](./go.mod) and the project default toolchain.
Changes to shared downloader or live-test code also test `latest/msgo` and `latest/msgo-previous`.
Each wrapper's downloaded SDK is cleaned up after its subtest to prevent disk usage accumulating when multiple wrappers are tested together.
Live-test subprocesses use an isolated Go telemetry configuration with collection disabled so background telemetry processes cannot race with SDK cleanup.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft 
trademarks or logos is subject to and must follow 
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
