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
See [What version should I use?](https://github.com/microsoft/go/blob/microsoft/main/eng/doc/MigrationGuide.md#what-version-should-i-use) for help choosing.
Also see [Wrapper versions](./VERSIONS.md) for other release channels and exact-version wrappers.

## Update a wrapper

A wrapper installed with `go install` does not update itself when its package points to a newer Microsoft build of Go.
Run the same `go install ...@latest` command again to replace the installed wrapper.

For a project-managed tool, you can record the wrapper package as a tool dependency in `go.mod` and run it without an `@version` suffix:

```console
go run github.com/microsoft/go-dl/latest/msgo download
```

Without an `@version` suffix, Go uses the go-dl module version selected by `go.mod`.
Dependency update tools such as Dependabot can then propose go-dl updates through the usual `go.mod` and `go.sum` changes.
If your project's CI uses the wrapper to build and test, the dependency-update pull request exercises the new Microsoft build of Go before merge.
This can catch toolset update breaks and prevent them from affecting the target branch.

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

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft 
trademarks or logos is subject to and must follow 
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
