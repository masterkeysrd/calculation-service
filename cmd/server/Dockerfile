FROM golang:1.20.4-alpine as builder
WORKDIR /app
RUN apk add --no-cache bash
COPY ./go.mod ./go.sum ./
COPY ./vendor ./vendor
COPY ./Makefile ./
COPY ./cmd/server ./cmd/server
COPY ./internal ./internal
COPY ./scripts ./scripts
RUN /bin/bash ./scripts/build.sh

FROM alpine:3.18
WORKDIR /app
ENV APP_ENV=production
ENV APP_CONFIG_PATH=/app/config
COPY --from=builder /app/bin/server ./server
COPY ./config ./config
CMD ["/app/server"]