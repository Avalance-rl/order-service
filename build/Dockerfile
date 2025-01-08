FROM golang:1.23-alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY internal/ ./internal

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -o /order-service cmd/main.go

FROM alpine:3.20

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser


WORKDIR /app

RUN mkdir /app/config
RUN mkdir /app/logs

COPY --from=build /order-service .

EXPOSE 8080

CMD ["./order-service"]
