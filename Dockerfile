# Combined Build and Test Stage
FROM golang:1.22.0-alpine3.19 AS build
WORKDIR /streamfair_token_svc
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o token_svc main.go

# Run Stage
FROM alpine:3.19
WORKDIR /streamfair_token_svc

# Copy the binary from the build stage
COPY --from=build /streamfair_token_svc/token_svc .

COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
COPY ssl .

EXPOSE   8082
EXPOSE   9092

CMD [ "/streamfair_token_svc/token_svc" ]
ENTRYPOINT [ "/streamfair_token_svc/start.sh" ]

RUN apk add --no-cache bash curl

HEALTHCHECK \
    --start-period=5s \
    --interval=1s \
    --timeout=1s \
    --retries=30 \
        CMD curl --fail -s  http://localhost:8082/streamfair/v1/healthz || exit 1