#!/bin/bash

set -eux
set -o errexit
set -o pipefail
set -o nounset

# CGO_ENABLED=1 GO111MODULE=on go mod tidy
# CGO_ENABLED=1 GO111MODULE=on go mod vendor
rm -fr ./build/plugin.so
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
  heroiclabs/nakama-pluginbuilder:3.11.0 \
    build \
      -a \
      -buildmode plugin \
      -mod vendor \
      -trimpath \
      -o ./build/plugin.so
