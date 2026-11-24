# Wrapper versions

Use these command pairs to install a wrapper and download the Microsoft build of Go it selects.

## Stable release tracking

`msgo` currently downloads Microsoft build of Go {{.CurrentVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo@latest
msgo download
```

`msgo-previous` follows the previous stable Go release series and currently downloads Microsoft build of Go {{.PreviousVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo-previous@latest
msgo-previous download
```

## Release series tracking

`msgo{{.CurrentSeries}}` stays on Go {{.CurrentSeries}} and currently downloads Microsoft build of Go {{.CurrentVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo{{.CurrentSeries}}@latest
msgo{{.CurrentSeries}} download
```

`msgo{{.PreviousSeries}}` stays on Go {{.PreviousSeries}} and currently downloads Microsoft build of Go {{.PreviousVersion}}.

```console
go install github.com/microsoft/go-dl/latest/msgo{{.PreviousSeries}}@latest
msgo{{.PreviousSeries}} download
```

## Exact release

Install a wrapper from `static` to select one exact Microsoft build of Go:

```console
go install github.com/microsoft/go-dl/static/msgo{{.CurrentVersion}}@latest
msgo{{.CurrentVersion}} download
```
