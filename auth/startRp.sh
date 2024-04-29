#!/usr/bin/env bash

set -euCo pipefail

docker compose -f compose-rp.yml up --build --force-recreate
