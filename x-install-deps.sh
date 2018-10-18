#!/bin/sh

set -e

go mod download
cd ui
npm install
cd ..
