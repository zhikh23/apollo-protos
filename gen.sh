#!/bin/sh

D_OUT="gen"
D_GO_OUT="$D_OUT"/go
D_PY_OUT="$D_OUT"/py


fetch_service_name() {
    name="${1#*proto\/}"
    name="${name%\.proto}"
    echo "$name"
}

for file in proto/*; do
    name="$(fetch_service_name "$file")"
    echo Compiling "$name"
    [ ! -d "$D_GO_OUT/$name" ] && mkdir "$D_GO_OUT/$name"
    [ ! -d "$D_PY_OUT/$name" ] && mkdir "$D_PY_OUT/$name"
    protoc -I proto "proto/${name}.proto" \
		--go_out="$D_GO_OUT/${name}" \
		--go_opt=paths=source_relative \
		--go-grpc_out="$D_GO_OUT/${name}" \
		--go-grpc_opt=paths=source_relative
    python -m grpc_tools.protoc \
        --proto_path=proto "${name}.proto" \
        --python_out="$D_PY_OUT/$name" \
        --grpc_python_out="$D_PY_OUT/${name}"
done
