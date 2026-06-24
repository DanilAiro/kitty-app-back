ARG SERVICE_NAME=kitty-app-back
ARG BUILD_DATE=$(date +%Y%m%d)

# ---------- build ----------
FROM golang:1-alpine AS builder

WORKDIR /app/kitty-app-back

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o kitty-app-back ./cmd/app

# ---------- runtime ----------
FROM alpine:3.20

LABEL org.opencontainers.image.source="https://github.com/DanilAiro/kitty-app-back"
LABEL org.opencontainers.image.title="kitty-app-back"
LABEL app.service="kitty-app-back"

RUN apk add --no-cache ca-certificates

# копирую .env и приложение в итоговый контейнер
COPY .env .
COPY --from=builder /app/kitty-app-back .

CMD ["./kitty-app-back"]