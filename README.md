# vul-db 

![Build DB](https://github.com/khulnasoft-lab/vul-db/workflows/Vul%20DB/badge.svg)
[![GitHub Release][release-img]][release]
![Downloads][download]
[![Go Report Card][report-card-img]][report-card]
[![Go Doc][go-doc-img]][go-doc]
[![License][license-img]][license]

[download]: https://img.shields.io/github/downloads/khulnasoft-lab/vul-db/total?logo=github
[release-img]: https://img.shields.io/github/release/khulnasoft-lab/vul-db.svg?logo=github
[release]: https://github.com/khulnasoft-lab/vul-db/releases
[report-card-img]: https://goreportcard.com/badge/github.com/khulnasoft-lab/vul-db
[report-card]: https://goreportcard.com/report/github.com/khulnasoft-lab/vul-db
[go-doc-img]: https://godoc.org/github.com/khulnasoft-lab/vul-db?status.svg
[go-doc]: https://godoc.org/github.com/khulnasoft-lab/vul-db
[code-cov]: https://codecov.io/gh/khulnasoft-lab/vul-db/branch/main/graph/badge.svg
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[license]: https://github.com/khulnasoft-lab/vul-db/blob/main/LICENSE

## Overview
`vul-db` is a CLI tool and a library to manipulate Vul DB.

### Library
Vul uses `vul-db` internally to manipulate vulnerability DB. This DB has vulnerability information from NVD, Red Hat, Debian, etc.

### CLI
The `vul-db` CLI tool builds vulnerability DBs. A [GitHub Actions workflow](.github/workflows/cron.yml)
periodically builds a fresh version of the vulnerability DB using `vul-db` and uploads it to the GitHub
Container Registry (see [Download the vulnerability database](#download-the-vulnerability-database) below).

```
NAME:
   vul-db - Vul DB builder

USAGE:
   main [global options] command [command options] image_name

VERSION:
   0.0.1

COMMANDS:
     build    build a database file
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Building the DB
You can utilize `make db-all` to build the database, the DB artifact is outputted to the assets folder.

Alternatively Docker is supported, you can run `docker build . -t vul-db`.

If you want to build a trivy integration test DB, please run `make create-test-db`

## Update interval
Every 6 hours

## Download the vulnerability database
### version 1 (deprecated)
Vul DB v1 reached the end of support on February 2023. Please upgrade Vul to v0.23.0 or later.

Read more about the Vul DB v1 deprecation in [the discussion](https://github.com/khulnasoft-lab/trivy/discussions/1653).

### version 2
Vul DB v2 is hosted on [GHCR](https://github.com/orgs/khulnasoft-lab/packages/container/package/vul-db).
Although GitHub displays the `docker pull` command by default, please note that it cannot be downloaded using `docker pull` as it is not a container image.

You can download the actual compiled database via [Vul](https://khulnasoft-lab.github.io/trivy/) or [Oras CLI](https://oras.land/cli/).

Vul:
```sh
VUL_TEMP_DIR=$(mktemp -d)
trivy --cache-dir $VUL_TEMP_DIR image --download-db-only
tar -cf ./db.tar.gz -C $VUL_TEMP_DIR/db metadata.json trivy.db
rm -rf $VUL_TEMP_DIR
```
oras >= v0.13.0:
```sh
$ oras pull ghcr.io/khulnasoft-lab/vul-db:2
```

oras < v0.13.0:
```sh
$ oras pull -a ghcr.io/khulnasoft-lab/vul-db:2
```
The database can be used for [Air-Gapped Environment](https://khulnasoft-lab.github.io/trivy/latest/docs/advanced/air-gap/).
