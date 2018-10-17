#!/bin/bash -x

set -e

BIN_DEPS=(goreleaser statik upx)

for dep in ${BIN_DEPS[@]}
do
    if ! [ -x "$(command -v ${dep})" ]; then
        echo "bin dependency $dep not found"
        exit 1
    fi
done

GOHOSTOS=$(go env GOHOSTOS) \
GOHOSTARCH=$(go env GOHOSTARCH) \
GOVERSION=$(go version | awk '{print $3;}') \
goreleaser --snapshot --rm-dist
