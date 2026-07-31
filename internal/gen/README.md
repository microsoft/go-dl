# Package generation

The build asset manifests in [`db`](./db) are the source data for every package under [`static`](../../static) and [`latest`](../../latest).
An external release process adds each manifest at a path such as `db/go1.26/go1.26.3-1.assets.json`.
Keep the downloaded file unchanged so the database remains an exact record of the published manifest.

To regenerate every package after adding a manifest or changing the template:

```console
cd internal/gen
go run ./cmd/generate-dl-packages
```

The command validates all embedded manifests, deletes the existing `static` and `latest` trees, recreates them completely, and updates the marked version examples in the root [`README.md`](../../README.md).
Do not edit files under those directories by hand.

Each manifest produces one version-specific package under `static`.
Each generated package contains a compile-time platform switch with the archive URL and checksum from its manifest.
The newest version in each Go release series produces its corresponding `latest/msgo1.x` package.
The two highest release series produce `latest/msgo` and `latest/msgo-previous`.

## Scraping published manifests

The release process normally adds one manifest directly to `db`.
To discover and archive all published manifests instead:

```console
cd internal/gen
go run ./cmd/scrape-db
```

The command starts at Go 1.22 and follows published Microsoft revisions, Go patch releases, and Go release series.
It validates each manifest before adding it to `db`.
If a checked-in manifest differs from the published bytes at the same URL, the command fails rather than replacing it.
