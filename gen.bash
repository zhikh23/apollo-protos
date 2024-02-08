#!/bin/bash

fetch_service_name() {
    name=${file#*proto\/}
    name=${name%\.proto}
}

for file in proto/*; do
    fetch_service_name
    echo Compiling "$name"
    protoc -I proto proto/${name}.proto \
		--go_out=./gen/go/${name}/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./gen/go/${name}/ \
		--go-grpc_opt=paths=source_relative \
        --python_out=./gen/py/${name}/
done
