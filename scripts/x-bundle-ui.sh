#!/bin/sh

# threat this script as if it's in the project root dir

set -e

cd ui
npm run-script build
find ./build/static -name "*.map" -delete
cd ..

statik -src=./ui/build -dest=./impl -f -p ui
