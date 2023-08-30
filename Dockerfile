FROM golang:1.21.0 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /build/cmd/app cmd/main.go


# Production
FROM alpine:latest
WORKDIR /app
EXPOSE 8080
RUN apk add --no-cache bash
RUN adduser \
    --disabled-password \
    --home /app \
    --gecos '' app \
    && chown -R app /app
USER app
COPY --from=builder /build/cmd/app .
COPY /data ./data
ENTRYPOINT [ "/app/app" ]
