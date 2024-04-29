#!/usr/bin/env bash

set -euCo pipefail

# Run docker run --rm -it --entrypoint='' oryd/kratos:v1.0.0 id to find the default uid:gid
sudo chown -R 10000:101 kratos-sqlite

# Run docker run --rm -it --entrypoint='' oryd/hydra:v2.2.0 id to find the default uid:gid
sudo chown -R 10000:101 hydra-sqlite

docker compose up --build --force-recreate
