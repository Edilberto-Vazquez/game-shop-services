ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS builder
RUN go env -w GOPROXY=https://goproxy.io
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /game-shop-services

FROM alpine:latest AS runner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=builder /game-shop-services /game-shop-services
EXPOSE 8000
ENTRYPOINT [ "/game-shop-services" ]