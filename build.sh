#!/bin/bash
GOOS=windows GOARCH=amd64 go build $(ls -1 *.go | grep -v _test.go)
mv app.exe launcher.exe
upx launcher.exe
echo "Built launcher.exe"
