#!/bin/bash

set -euCo pipefail

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --proto_path=/usr/local/include/:. \
    proto/api.proto
