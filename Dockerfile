FROM golang:alpine3.19 AS builder

RUN apk update && apk upgrade && \
    apk add make build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -race -o main .

FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/env /app/env

EXPOSE 8080

CMD ["./main"]
