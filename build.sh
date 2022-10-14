#!/bin/bash
echo "Building Windows amd64..."
GOOS=windows GOARCH=amd64 go build -o build/jsonform-win-amd64.exe src/main.go
echo "Building Windows arm64..."
GOOS=windows GOARCH=arm64 go build -o build/jsonform-win-arm64.exe src/main.go
echo "Building Linux amd64..."
GOOS=linux GOARCH=amd64 go build -o build/jsonform-linux-amd64 src/main.go
echo "Building Linux arm64..."
GOOS=linux GOARCH=arm64 go build -o build/jsonform-linux-arm64 src/main.go
echo "Building MacOS amd64..."
GOOS=darwin GOARCH=amd64 go build -o build/jsonform-mac-amd64 src/main.go
echo "Building MacOS arm64..."
GOOS=darwin GOARCH=arm64 go build -o build/jsonform-mac-arm64 src/main.go
echo "Done."
