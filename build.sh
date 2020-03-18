#!/usr/bin/env bash
set -e
rm -rf ./bin

GOOS="darwin" GOARCH="amd64" go build  -ldflags="-s -w" -o "bin/ipsec-exporter_darwin_amd64"
GOOS="windows" GOARCH="amd64" go build  -ldflags="-s -w" -o "bin/ipsec-exporter_windows_amd64.exe"
GOOS="windows" GOARCH="386" go build  -ldflags="-s -w" -o "bin/ipsec-exporter_windows_386.exe"
GOOS="linux" GOARCH="amd64" go build  -ldflags="-s -w" -o "bin/ipsec-exporter_linux_amd64"
GOOS="linux" GOARCH="386" go build  -ldflags="-s -w" -o "bin/ipsec-exporter_linux_386"

chmod +x ./bin/*