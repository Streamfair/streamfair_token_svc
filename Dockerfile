# Build Stage
FROM golang:1.22.0-alpine3.19 AS build
WORKDIR /streamfair_token_svc
COPY . .
RUN go build -o token_svc main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.19
WORKDIR /streamfair_token_svc

# Copy the binary from the build stage
COPY --from=build /streamfair_token_svc/token_svc .
# Copy the downloaded migration binary from the build stage
COPY --from=build /streamfair_token_svc/migrate ./migrate

COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8082
EXPOSE 9092

CMD [ "/streamfair_token_svc/token_svc" ]
ENTRYPOINT [ "/streamfair_token_svc/start.sh" ]