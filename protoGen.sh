#!/usr/bin/env bash

set -euCo pipefail

declare -r SCRIPT_DIR=$(cd $(dirname $0) && pwd)
declare -r PROTO_DIR="${SCRIPT_DIR}/proto"
declare -r JS_GEN_OUT=${SCRIPT_DIR}/frontend-simple/my-app/proto

function generateGo() {
  protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      --proto_path=/usr/local/include/:. \
      proto/api.proto proto/weight.proto
}

function generateJs() {
    if [[ ! -d ${JS_GEN_OUT} ]]
      then
        mkdir -p ${JS_GEN_OUT}
    fi

    protoc --proto_path=/usr/local/include/:"${PROTO_DIR}" \
      api.proto \
      --js_out=import_style=commonjs:${JS_GEN_OUT} \
      --grpc-web_out=import_style=commonjs,mode=grpcwebtext:${JS_GEN_OUT}
}

#generateGo
generateJs
