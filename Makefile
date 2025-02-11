.PHONY: build
build:
	GOARCH=amd64 go build -o ./bin/pad-amd64 cmd/main.go

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
        --proto_path=/usr/local/include/:. \
      	proto/api.proto proto/weight.proto