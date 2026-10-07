# Wrapper versions

Use these commands to install wrappers for other release channels or one exact Microsoft build of Go.

## Stable release tracking

`msgo` currently downloads Microsoft build of Go {{.CurrentVersion}}.

`msgo-previous` follows the previous stable Go release series and currently downloads Microsoft build of Go {{.PreviousVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo-previous@latest
```

## Release series tracking

`msgo{{.CurrentSeries}}` stays on Go {{.CurrentSeries}} and currently downloads Microsoft build of Go {{.CurrentVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo{{.CurrentSeries}}@latest
```

`msgo{{.PreviousSeries}}` stays on Go {{.PreviousSeries}} and currently downloads Microsoft build of Go {{.PreviousVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo{{.PreviousSeries}}@latest
```

## Exact release

Install a wrapper from `static` to select one exact Microsoft build of Go:

```console
go install github.com/microsoft/go-dl/static/msgo{{.CurrentVersion}}@latest
```
