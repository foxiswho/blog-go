#!/bin/bash
set -e  # 出错立即退出

# 配置参数
SCAN_DIR="./app,./infrastructure"
OUTPUT_FILE="auto_import.go"
# 模块名,不填写时，自动获取当前 go.mod 文件的模块名
MODULE_NAME=""

# 第一步：运行扫描工具
echo "===== 运行扫描工具 ====="
go run ./auto_tool/main.go -module=$MODULE_NAME -dir=$SCAN_DIR -output=$OUTPUT_FILE -concurrency=8 -ignore=**/vendor/**,**/.git/**,**/*_test.go,**/testdata/**,**/node_modules/**

# 第二步：运行项目
echo "===== 启动项目 ====="
go run .