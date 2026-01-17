########## BUILD STAGE ##########
FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/chat ./cmd/chat

########## RUNTIME STAGE ##########
FROM alpine:3.20

RUN apk add --no-cache tzdata

ENV TZ=Europe/Moscow
RUN ln -sf /usr/share/zoneinfo/Europe/Moscow /etc/localtime

WORKDIR /app

COPY .env /app/
COPY config/* /app/config/
COPY --from=builder /out/chat /app/

ENTRYPOINT ["/app/chat"]
