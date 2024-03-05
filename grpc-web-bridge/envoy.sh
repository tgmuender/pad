#!/bin/bash
# See
# - https://github.com/grpc/grpc-web
# - https://github.com/grpc/grpc-web/tree/master/net/grpc/gateway/examples/helloworld#configure-the-proxy

set -euCo pipefail

docker run -v "$(pwd)"/envoy.yaml:/etc/envoy/envoy.yaml:ro \
    --network=host \
    envoyproxy/envoy:distroless-v1.29.0
