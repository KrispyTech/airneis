FROM golang:1.21 AS base

WORKDIR /build

COPY . .
RUN go mod download && go mod verify

FROM base as dev

WORKDIR /build

RUN go install github.com/cosmtrek/air@v1.49.0
CMD ["air", "-c", ".air.toml"]

FROM base as builder

WORKDIR /build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o airneis .

FROM scratch as prod

COPY --from=builder ["/build/pkg/config", "pkg/config"]
COPY --from=builder ["/build/.env", "/"]
COPY --from=builder ["/build/airneis", "/"]
CMD ["/airneis"]