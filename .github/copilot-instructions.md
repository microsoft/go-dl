# Repository instructions

## Upstream-owned version files

`upstreamFiles` in `internal/gen/cmd/fetch-upstream/main.go` is the source of truth for files copied from `golang.org/dl`.
Treat every file in that list as a vendored, byte-for-byte copy.

Never edit these files directly, even to fix a bug or security issue.
Never add tests that require these files to differ from upstream.
Do not edit `internal/version/upstream-files.sha256` to approve a local change.

In a directory that contains files listed in `upstreamFiles`, put Microsoft-specific behavior and tests in files whose names start with `ms_`.
If a change cannot be implemented without modifying an upstream-owned file, stop and explain the conflict rather than changing the file.

Update upstream-owned files and their checksums only by running:

```console
cd internal/gen
go run ./cmd/fetch-upstream
```

See `internal/version/README.md` for details.
