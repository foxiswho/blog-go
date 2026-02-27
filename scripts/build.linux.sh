


cd ../

rm -rf blogGo

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -ldflags '-extldflags "-static"' -o blogGo .



echo "build success"