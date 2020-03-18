#!/usr/bin/env bash
set -e
rm -rf ./bin

#GOOS="darwin" GOARCH="amd64" go build -o "bin/digibee_darwin_amd64" main.go
#GOOS="windows" GOARCH="amd64" go build -o "bin/digibee_windows_amd64.exe" main.go
#GOOS="windows" GOARCH="386" go build -o "bin/digibee_windows_386.exe" main.go
#GOOS="linux" GOARCH="amd64" go build -o "bin/digibee_linux_amd64" main.go
#GOOS="linux" GOARCH="386" go build -o "bin/digibee_linux_386" main.go

GOOS="linux" GOARCH="amd64" go build -ldflags="-s -w" -o "bin/digibee_linux_amd64"

chmod +x ./bin/*

git add --all &&\
git commit -m "add new binaries" &&\
git push origin master
