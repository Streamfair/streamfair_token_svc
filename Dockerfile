# Build Stage
FROM golang:1.22.0-alpine3.19 AS build
WORKDIR /streamfair_token_service
COPY . .
RUN go build -o streamfair_token_service main.go

# Run Stage
FROM alpine:3.19
WORKDIR /streamfair_token_service
COPY --from=build /streamfair_token_service/streamfair_token_service .
COPY app.env .

EXPOSE 8082
EXPOSE 9092

CMD ["./streamfair_token_service"]