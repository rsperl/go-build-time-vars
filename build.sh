#!/usr/bin/env bash
set -eu

BINARY=main
DIST_DIR="./dist"
ACTION="${1:-}"

if [[ $ACTION == "build" ]]; then
  set -x
  go generate -v
  go build -o "$DIST_DIR/$BINARY" .
  set +x
elif [[ $ACTION == "clean" ]]; then
  rm -rf ./tmp
  rm -rf "$DIST_DIR"
else
  echo "Usage: $0 <build|clean>"
  exit 1
fi
