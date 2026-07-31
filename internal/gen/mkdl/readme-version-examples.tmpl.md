<!-- BEGIN GENERATED VERSION EXAMPLES: DO NOT EDIT -->

### Stable release tracking

`msgo` currently downloads Microsoft build of Go {{currentVersion}}.

`msgo-previous` follows the previous stable Go release series and currently downloads Microsoft build of Go {{previousVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo-previous@latest
```

### Release series tracking

`msgo{{currentSeries}}` stays on Go {{currentSeries}} and currently downloads Microsoft build of Go {{currentVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo{{currentSeries}}@latest
```

`msgo{{previousSeries}}` stays on Go {{previousSeries}} and currently downloads Microsoft build of Go {{previousVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo{{previousSeries}}@latest
```

### Exact release

Install a wrapper from `static` to select one exact Microsoft build of Go:

```console
go install github.com/microsoft/go-dl/static/msgo{{currentVersion}}@latest
```

## Update a wrapper

A wrapper installed with `go install` does not update itself when its package points to a newer Microsoft build of Go.
Run the same `go install ...@latest` command again to replace the installed wrapper.

For a project-managed tool, you can record the wrapper package as a tool dependency in `go.mod` and run it without an `@version` suffix:

```console
go run github.com/microsoft/go-dl/latest/msgo{{currentSeries}} download
```

Without an `@version` suffix, Go uses the go-dl module version selected by `go.mod`.
Dependency update tools such as Dependabot can then propose go-dl updates through the usual `go.mod` and `go.sum` changes.
If your project's CI uses the wrapper to build and test, the dependency-update pull request exercises the new Microsoft build of Go before merge.
This can catch toolset update breaks and prevent them from affecting the target branch.

<!-- END GENERATED VERSION EXAMPLES -->
