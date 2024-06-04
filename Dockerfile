FROM golang:1.21-alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]