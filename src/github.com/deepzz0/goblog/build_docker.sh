#!/bin/bash
echo "go build ..."
sh ./version.sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build && \
mkdir tmp && \
# cp -a $GAMEDATA/config/json  tmp/
cp -a views tmp/views && \
cp -a conf tmp/conf && \
cp -a static tmp/static && \
cp Dockerfile tmp/ && \
cp goblog tmp/ && \
cp version tmp/ && \
cd tmp && \

# docker -H tcp://127.0.0.1:2375 build -t gameserver .
docker build -t registry.cn-hangzhou.aliyuncs.com/deepzz/goblog . 
cd .. 
rm -rf tmp
