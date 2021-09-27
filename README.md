# WebApp Sidecar

Sidecar service on Go for webapp which expose some metrics in Prometheus format.

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