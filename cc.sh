#!/bin/bash

# 定义目录
PROTO_DIR="./proto"
PB_DIR="./pb"

# 创建 pb 目录（如果不存在）
mkdir -p $PB_DIR

# 查找所有 .proto 文件并编译
for proto_file in $PROTO_DIR/*.proto; do
    if [ -f "$proto_file" ]; then
        echo "Compiling $proto_file..."
        protoc --go_out=$PB_DIR "$proto_file"
    fi
done

echo "Compilation completed."
