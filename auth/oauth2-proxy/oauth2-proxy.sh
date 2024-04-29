#!/bin/bash
# See
# - https://github.com/grpc/grpc-web
# - https://github.com/grpc/grpc-web/tree/master/net/grpc/gateway/examples/helloworld#configure-the-proxy

set -euCo pipefail

docker run -v "$(pwd)"/oauth2-proxy.cfg:/oauth2-proxy.cfg:ro \
    --network=host quay.io/oauth2-proxy/oauth2-proxy:v7.6.0 \
    --config /oauth2-proxy.cfg