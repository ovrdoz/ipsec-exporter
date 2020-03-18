
# Ipsec Exporter
Prometheus exporter for ipsec metrics, written in Go.

## How to build 

Clone repository
````
git clone https://github.com/ovrdoz/ipsec-exporter.git
cd ipsec-exporter
````
Make a local build and run
````
GOOS="<so>" GOARCH="<arch>" go build  -ldflags="-s -w" -o "bin/ipsec-exporter_darwin_amd64"
````
The following table shows the possible combinations of GOOS and GOARCH you can use:
|GOOS - Target Operating System|GOARCH - Target Platform|
|--- |--- |
|android|arm|
|darwin|386|
|darwin|amd64|
|darwin|arm|
|darwin|arm64|
|dragonfly|amd64|
|freebsd|386|
|freebsd|amd64|
|freebsd|arm|
|linux|386|
|linux|amd64|
|linux|arm|
|linux|arm64|
|linux|ppc64|
|linux|ppc64le|
|linux|mips|
|linux|mipsle|
|linux|mips64|
|linux|mips64le|
|netbsd|386|
|netbsd|amd64|
|netbsd|arm|
|openbsd|386|
|openbsd|amd64|
|openbsd|arm|
|plan9|386|
|plan9|amd64|
|solaris|amd64|
|windows|386|
|windows|amd64|


## Functionality
The IPsec exporter is determining the state of the configured IPsec tunnels via the following procedure.
1. Starting up the `ipsec.conf` is read. All tunnels configured via the `conn` keyword are observed.
1. If the `/metrics` endpoint is queried, the exporter calls `ipsec status <tunnel name>` for each configured
connection. The output is parsed.
    * If the output contains `ESTABLISHED`, we assume that only the connection is up.
    * If the output contains `INSTALLED`, we assume that the tunnel is up and running.
    * If the output contains `no match`, we assume that the connection is down.

## Value Definition
| Metric | Value | Description |
|--------|-------|-------------|
| ipsec_status | 0 | The connection is established and tunnel is installed. The tunnel is up and running. |
| ipsec_status | 1 | The connection is established, but the tunnel is not up. |
| ipsec_status | 2 | The tunnel is down. |
| ipsec_status | 3 | The tunnel is in an unknown state. |
| ipsec_status | 4 | The tunnel is ignored. |# tmp

This project is based in dennisstritzke project

