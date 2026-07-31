#!/bin/bash
# 当前文件目录
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )


cd $SCRIPT_DIR/../

rm -rf xianFuBlogGo

make build

echo "build success"