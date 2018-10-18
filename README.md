# sysconfig

[![demo container status](https://quay.io/repository/goiiot/sysconfig/status)](https://quay.io/repository/goiiot/sysconfig)
 [![GoReportCard](https://goreportcard.com/badge/goiiot/sysconfig)](https://goreportcard.com/report/github.com/goiiot/sysconfig)

Configure your embedded system with ease

## Table of Contents

- [Features](#features)
- [Demo](#demo)
- [Supported Platforms](#supported-platforms)
- [Usage](#usage)
- [Build](#build)
    - [Prerequisite](#prerequisite)
    - [Steps](#steps)

## Features

- Shell execution (via `xterm` and `websocket`)
- File upload and download
- Device monitoring
- Device configuration
- Power management

## Demo

Try sysconfig demo with docker

```bash
$ docker run -d \
    --name sysconfig-demo \
    -p 8080:8080 \
    -p 8443:8443 \
    goiiot/sysconfig:demo

# or use image at quay.io
#
# $ docker run -d --name sysconfig-demo -p 8080:8080 -p 8443:8443 quay.io/goiiot/sysconfig:latest
```

## Supported Platforms

- Darwin
- Linux

## Usage

1. Make a copy of `config.yaml` by `cp config.example.yaml config.yaml`
2. Modify `config.yaml` according to your system and needs
3. Start server with `./config-server --config ./config.yaml`

## Build

### Prerequisite

1. [Git](https://git-scm.com/)
2. [Go](https://golang.org/) 1.11+ (Go modules required)
3. [gorealser](https://goreleaser.com) for build ease
4. [statik](https://github.com/rakyll/statik) for bundling web app
5. [upx](https://github.com/upx/upx) for binary compression
6. [nodejs and npm](https://nodejs.org) for web app build

### Steps

1. `cd $THIS_PROJECT_DIR`
2. `./x-install-deps.sh`
3. `./x-build.sh`
4. Find output in `dist` directory

## LICENSE

```text
Copyright Go-IIoT (https://github.com/goiiot)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```