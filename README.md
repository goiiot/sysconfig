# Config server

[![demo container status](https://quay.io/repository/goiiot/sysconfig/status)](https://quay.io/repository/goiiot/sysconfig)
 [![GoReportCard](https://goreportcard.com/badge/goiiot/sysconfig)](https://goreportcard.com/report/github.com/goiiot/sysconfig)

- [Config server](#config-server)
    - [Features](#features)
    - [Supported Platforms](#supported-platforms)
    - [Usage](#usage)
    - [Build](#build)
        - [Prerequisite](#prerequisite)
        - [Steps](#steps)

## Features

- Shell execution (via `xterm` and `websocket`)
- File upload and download
- Device monitoring
- TODO Configuration
- TODO Power management

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
6. [nodejs](https://nodejs.org) for web app build

__NOTE__: You can install all these dependencies by executing the script `x-install-deps.sh`

### Steps

1. `cd $THIS_PROJECT_DIR`
2. `./x-build.sh`
3. Find output in `dist` directory
