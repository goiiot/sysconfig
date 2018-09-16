#!/bin/sh

set -e

# delete .DS_Store
find . -name ".DS_Store" -delete
# delete tarballs and zip files
find . -name "*.tar.*" -delete
find . -name "*.zip" -delete

MISC=(./dist ./ui/build ./server)
for toDel in ${MISC[@]}
do
    rm -rf $toDel
done
