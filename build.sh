#!/usr/bin/env bash
set -eu

cd "$(dirname "${BASH_SOURCE[0]}")"

ROOT_DIR="$(realpath ".")"
export GOPATH="${ROOT_DIR}/workspace"
./lib/goyacc -o ./src/grammar/y.go -v ./src/grammar/y.output ./src/grammar/parser.y
(cd src && go build -o "${ROOT_DIR}/bin/compiler" mit.edu/compilers/compiler)
