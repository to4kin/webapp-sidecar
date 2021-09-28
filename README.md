[![ci-build](https://github.com/to4kin/webapp-sidecar/actions/workflows/ci-build.yml/badge.svg?branch=master)](https://github.com/to4kin/webapp-sidecar/actions/workflows/ci-build.yml)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/to4kin/webapp-sidecar.svg)](https://github.com/to4kin/webapp-sidecar)
[![GitHub release](https://img.shields.io/github/release/to4kin/webapp-sidecar.svg)](https://GitHub.com/to4kin/webapp-sidecar/releases/)

# WebApp Sidecar

Sidecar service on Go for webapp which expose some metrics in Prometheus format.

[![GitHub](https://badgen.net/badge/icon/github?icon=github&label)](https://github.com/to4kin/webapp-sidecar)  
[![Docker](https://badgen.net/badge/icon/docker?icon=docker&label)](https://hub.docker.com/repository/docker/to4kin/webapp-sidecar)  

## Precondition

* Go 1.17+

### Start via Docker

Default location for custom file metrics is `/upload`

```bash
docker run -it --rm -p 3000:3000 -v `pwd`/upload:/upload --name webapp-sidecar to4kin/webapp-sidecar:latest
```

### Usage

```bash
Sidecar for WebApp to provide Prometheus metrics

Usage:
  webapp-sidecar [command]

Available Commands:
  help        Help about any command
  start       Start WebApp Sidecar
  version     Print version

Flags:
  -h, --help   help for webapp-sidecar

Use "webapp-sidecar [command] --help" for more information about a command.
```

### Start sidecar

```bash
Start WebApp Sidecar with config file
Simply execute webapp-sidecar start -c path/to/config/file.toml
or skip this flag to use default path

Usage:
  webapp-sidecar start [flags]

Flags:
  -c, --config-path string   path to config file (default "configs/webappsidecar.toml")
  -h, --help                 help for start
```

### Config file

```toml
bind_addr = ":3000"
upload_folder = "upload"
check_folder_interval = 2
metrics_path = "/metrics"
```

### Exposed custom metrics

```
# HELP files_count The number of files in upload folder
# TYPE files_count gauge
files_count 1
# HELP files_size_total The total size of files in the upload folder
# TYPE files_size_total gauge
files_size_total 0
```