# Upstream version package

Most files in this directory come directly from [`golang/dl/internal/version`](https://github.com/golang/dl/tree/master/internal/version).
Microsoft build of Go behavior belongs in files prefixed with `ms_` so that the remaining files can stay identical to upstream.
This repository does not copy `gotip.go` because it does not provide a `gotip` wrapper.

Use [`fetch-upstream`](../gen/cmd/fetch-upstream) to update the upstream-owned files and their checksums:

```console
cd internal/gen
go run ./cmd/fetch-upstream
```

Review the upstream changes after running the command.
The tests in this package verify the checked-in files against [`upstream-files.sha256`](./upstream-files.sha256) without accessing the network.
