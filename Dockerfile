FROM alpine:3.14.2
COPY bin/webapp-sidecar.linux.amd64 /app/webapp-sidecar
COPY configs/webappsidecar.toml /configs/webappsidecar.toml
ENTRYPOINT ["/app/webapp-sidecar", "start"]