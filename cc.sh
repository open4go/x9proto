#!/bin/bash

# 定义目录
PROTO_DIR="./proto"
PB_DIR="./pb"

# 创建 pb 目录（如果不存在）
mkdir -p $PB_DIR

# 查找所有 .proto 文件并编译
for proto_file in $(find $PROTO_DIR -name "*.proto"); do
    if [ -f "$proto_file" ]; then
        echo "Compiling $proto_file..."

        # 使用绝对路径调用 protoc-gen-go 和 protoc-gen-go-grpc
        protoc --proto_path=$PROTO_DIR \
               --go_out=$PB_DIR \
               --go-grpc_out=$PB_DIR \
               "$proto_file"
    fi
done

echo "Compilation completed."
