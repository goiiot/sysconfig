#!/bin/sh

# use upx to compress executables
# this script will be executed after `goreleaser` finishing building

# threat this script as if it's in the project root dir

set -e

if ! [ -x "$(command -v upx)" ]; then
  echo "upx not found, not going to compress"
  exit 1
fi

DIST_DIR="./dist"
BIN_NAME="sysconfig"
GOOS_LIST="linux darwin"
GOARCH_LIST="amd64 386 arm64 arm_7 arm_6 arm_5"
BIN_DIR_LIST=""

for OS in ${GOOS_LIST}
do
  for ARCH in ${GOARCH_LIST}
  do
    BIN_DIR_LIST="${OS}_${ARCH} ${BIN_DIR_LIST}"
  done
done

upx=$(which upx)

for target in ${BIN_DIR_LIST[@]}
do
  b="$DIST_DIR/$target/$BIN_NAME"
  if [ -f $b ]; then
    $upx --brute "$DIST_DIR/$target/$BIN_NAME"
  fi
done
