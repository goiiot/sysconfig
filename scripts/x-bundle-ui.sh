#!/bin/bash -x

# threat this script as if it's in the project root dir

set -e

pushd ui
npm run-script build
find ./build/static -name "*.map" -delete
popd

statik -src=./ui/build -dest=./impl -f -p ui