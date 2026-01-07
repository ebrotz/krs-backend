# Start by building the application.
FROM golang:1.25 AS build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY api api/
COPY cmd cmd/
COPY internal internal/

RUN CGO_ENABLED=0 go build -C cmd -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /
CMD ["/app"]
