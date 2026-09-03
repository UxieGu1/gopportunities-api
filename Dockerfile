FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -trimpath -ldflags="-s -w" -o /app/gopportunities-api ./cmd/api

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates \
	&& addgroup -S app \
	&& adduser -S -G app app

COPY --from=builder /app/gopportunities-api .
RUN chown app:app /app/gopportunities-api

USER app

EXPOSE 8081

CMD ["./gopportunities-api"]