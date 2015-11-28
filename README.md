droot
========

Dochroot is a CLI tool for chrooting a docker image.

## Overview

[Docker](https://www.docker.com) has a powerfull concept about application deployment process, that is Build, Ship, Run. But there are many cases that docker runtime is beyond our current capabilities. I supporse simpler container runtime by chrooting into docker image. `droot` helps you chrooting an image built by `docker build` command and managing exported images on Amazon S3 or a local filesystem.

## Requirements

- Docker (`droot push` only depends on it)

## Installation

## Usage

```bash
$ droot push --to s3://example.com/dockerfiles/app.tar.gz dockerfiles/app
```

```bash
$ droot pull --dest /var/containers/app --src s3://example.com/dockerfiles/app.tar.gz
```

```bash
# droot run --bind /var/log/ --root /var/containers/app -- command
```

```bash
# droot umount --root /var/containers/app
```

## Roodmap

- `rm` command for cleaning container environment
- `rmi` command for cleaning image on S3
- `pull` command with the rsync option
- `push/pull` other compression algorithms
- image versioning
- `pull` from docker registry
- drivers except Amazon S3

