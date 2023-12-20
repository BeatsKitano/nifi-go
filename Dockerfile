FROM golang:1.21 AS builder

ARG APP_RELATIVE_PATH
ARG APP_VERSION

COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}
RUN mkdir -p bin/ && go build -ldflags "-X main.Version=${APP_VERSION}" -o ./bin/ ./...

FROM debian:stable-slim

ARG APP_RELATIVE_PATH

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app

WORKDIR /app

EXPOSE 18000
VOLUME /configs

CMD ["./server", "-conf", "/configs"]