#!/bin/bash

set -eux
set -o errexit
set -o pipefail
set -o nounset

export CGO_ENABLED=1
export GO111MODULE=on
export GOARCH=amd64
export GOOS=linux

DOCKER_COMPOSE=${1:-false}
OUTPUT_BIN=${2:-"plugin.so"}
OUTPUT_DIR=${3:-"build"}
OUTPUT=./${OUTPUT_DIR}/${OUTPUT_BIN}

rm -fr ${OUTPUT}

go mod tidy
go mod vendor

if ${DOCKER_COMPOSE}; then

go build \
  -a \
  -buildmode plugin \
  -mod vendor \
  -trimpath \
  -o ${OUTPUT}

else

COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker run \
  --env CGO_ENABLED=1 \
  --env GO111MODULE=on \
  --env GOARCH=amd64 \
  --env GOOS=linux \
  --interactive \
  --platform linux/amd64 \
  --rm \
  --tty \
  --volume $(pwd):/workspace \
  --workdir /workspace \
  heroiclabs/nakama-pluginbuilder:3.12.0 build \
    -a \
    -buildmode plugin \
    -mod vendor \
    -trimpath \
    -o ${OUTPUT}

fi