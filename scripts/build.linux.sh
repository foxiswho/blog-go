#!/bin/bash
# 当前文件目录
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )


cd $SCRIPT_DIR/../

rm -rf blogGo
version="v0.0.1"
GitCommit=hhhhhh
COMPACT_TS=$(date +"%Y%m%d%H%M%S")

ldflags='-s -w -extldflags "-static" -X "main.UserName=" -X "github.com/foxiswho/blog-go/cmd.BuildVersion='${version}'" '
ldflags=${ldflags}' -X "github.com/foxiswho/blog-go/cmd.BuildGitCommit='${GitCommit}'" '
ldflags=${ldflags}' -X "github.com/foxiswho/blog-go/cmd.BuildTime='${COMPACT_TS}'" '

#echo ${ldflags}

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
go build -a -ldflags "${ldflags}" \
-gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -trimpath \
-o blogGo .



echo "build success"