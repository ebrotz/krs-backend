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

When using Postgres, first bring up the database:

```shell
docker compose up
```

Then start the backend:

```shell
make run
```

## Build Image

```shell
docker image build -t "krs-backend:$(git rev-parse --short HEAD)" .
```

