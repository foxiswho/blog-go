#!/bin/sh

curl -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: 55371481847297741127:1ec2c4d971831ba980c57846222b3a27368cc2dbd1048f8e1724d88c40ff6ce0" \
  -d '{"url":"http://localhost:9981/api/collect/push","title":"测试测试","description":"北京","tags":["测试TAG"]}' \
  http://localhost:9981/api/collect/push

