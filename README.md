# krs-backend

Backend for KRS proof of concept.

## Generate Code

```shell
go tool oapi-codegen -config api/cfg.yaml api/openapi.yaml
```

## Build Server

```shell
go build -C cmd -o '../krs-backend'
```

## Run Server Locally

```shell
go run -C cmd server.go
```

## Build Image

```shell
docker image build -t "krs-backend:$(git rev-parse --short HEAD)" .
```
