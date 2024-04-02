###  VARIABLES  ###
ifndef GITHUB_ACTIONS
-include ./env/db.env
endif

# Docker Network
DB_NETWORK := db_access_network
SERVICE_NETWORK := streamfair_internal_network

# Service Container
## ADJUST FOR EACH SERVICE ##
SERVICE_IMAGE := streamfair_token_svc
SERVICE_TAG := latest

GRPC_PORT := 9092
GRPC_HOST_PORT := 9390

GRPC_GATEWAY_PORT := 8082
GRPC_GATEWAY_HOST_PORT := 8380

# Database Container
## ADJUST FOR EACH SERVICE ##
DB_IMAGE := postgres:16-alpine
DB_CONTAINER_NAME := token_service_db

DB_PORT := 5432
DB_HOST_PORT := 5434

DB_NAME := $(or $(DB_NAME),$(shell grep POSTGRES_DB ./env/db.env | cut -d '=' -f2))
DB_USER := $(or $(DB_USER),$(shell grep POSTGRES_USER ./env/db.env | cut -d '=' -f2))
DB_PASSWORD := $(or $(DB_PASSWORD),$(shell grep POSTGRES_PASSWORD ./env/db.env | cut -d '=' -f2))

DB_HOST := localhost
DB_SOURCE_SERVICE := "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_NETWORK}:${DB_HOST_PORT}/${DB_NAME}?sslmode=disable"
DB_SOURCE_MIGRATION := "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_HOST_PORT}/${DB_NAME}?sslmode=disable"

# Migrate
MIGRATION_DIR := db/migration
MIGRATION_NAME := init_schema

# Proto
## ADJUST FOR EACH SERVICE ##
COMMON_PROTO_DIR := ../CommonProto
COMMON_PROTO_ERROR_DIR := ../CommonProto/error
PROTO_DIR := ../CommonProto/TokenService/proto
PB_DIR := ../CommonProto/TokenService/pb
TOKEN_DIR := token
REFRESH_TOKEN_DIR := refresh_token

# Test
TEST_DIR := ./...
DB_TEST_DIR := ./db/sqlc
API_TEST_DIR := ./api
UTIL_TEST_DIR := ./util
SERVER_TEST_DIR := ./gapi
# Test Output
TEST_FILE := tests.log
DB_TEST_FILE := db_tests.log
API_TEST_FILE := api_tests.log
UTIL_TEST_FILE := util_tests.log
SERVER_TEST_FILE := server_tests.log
# Output Flag
OUT ?= 0

# Go
ENTRY_POINT := main.go

# Mock-gen
MOCK_SOURCE := db/sqlc/store.go
MOCK_DEST := db/mock/store_mock.go

# Documentation
SWAGGER_DIR := doc/swagger
SWAGGER_DOC_NAME := streamfair_token_service


###  TARGETS  ###
# Network Management
network:
	@docker network inspect ${DB_NETWORK} >/dev/null 2>&1 || \
	docker network create --driver bridge ${DB_NETWORK}
	@docker network inspect ${SERVICE_NETWORK} >/dev/null 2>&1 || \
	docker network create --driver bridge ${SERVICE_NETWORK}

# Service Management
service_image:
	@docker build -t ${SERVICE_IMAGE}:${SERVICE_TAG} .

service_container:
	@docker run --name ${SERVICE_IMAGE} --network ${DB_NETWORK} --network ${SERVICE_NETWORK} -p ${GRPC_GATEWAY_HOST_PORT}:${GRPC_GATEWAY_PORT} -p ${GRPC_HOST_PORT}:${GRPC_PORT} -e DB_SOURCE=${DB_SOURCE_SERVICE} ${SERVICE_IMAGE}:${SERVICE_TAG}

# DB Management
db_container: network
	@docker start ${DB_CONTAINER_NAME} >/dev/null 2>&1 || \
	docker run --name ${DB_CONTAINER_NAME} --network ${DB_NETWORK} -p ${DB_HOST_PORT}:${DB_PORT} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d ${DB_IMAGE}

createdb:
	@sleep 1
	@docker exec -it ${DB_CONTAINER_NAME} psql -U ${DB_USER} -lqt | cut -d \| -f 1 | grep -qw ${DB_NAME} >/dev/null 2>&1 || \
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	@docker exec -it ${DB_CONTAINER_NAME} dropdb ${DB_NAME}

createmigration:
	migrate create -ext sql -dir ${MIGRATION_DIR} -seq ${MIGRATION_NAME}

migrateup:
	@migrate -path ${MIGRATION_DIR} -database ${DB_SOURCE_MIGRATION} -verbose up

migrateup1:
	@migrate -path ${MIGRATION_DIR} -database ${DB_SOURCE_MIGRATION} -verbose up 1

migratedown:
	@migrate -path ${MIGRATION_DIR} -database ${DB_SOURCE_MIGRATION} -verbose down

migratedown1:
	@migrate -path ${MIGRATION_DIR} -database ${DB_SOURCE_MIGRATION} -verbose down 1

dbclean: migratedown migrateup
	clear


# Execution
server: network db_container createdb
	@go run ${ENTRY_POINT}

# Cleanup
down:
	@if [ "$$(docker ps -aq -f name=$(DB_CONTAINER_NAME))" ]; then \
		docker stop $(DB_CONTAINER_NAME); \
		docker rm $(DB_CONTAINER_NAME); \
	fi
	@if [ "$$(docker ps -aq -f name=$(SERVICE_IMAGE))" ]; then \
		docker stop $(SERVICE_IMAGE); \
		docker rm $(SERVICE_IMAGE); \
	fi
	@if [ "$$(docker network ls -q -f name=$(DB_NETWORK))" ]; then \
		docker network rm $(DB_NETWORK); \
	fi
	@if [ "$$(docker network ls -q -f name=$(SERVICE_NETWORK))" ]; then \
		docker network rm $(SERVICE_NETWORK); \
	fi


# SQLC Generation
sqlc:
	sqlc generate


# Mock Generation
mock:
	mockgen -source=${MOCK_SOURCE} -destination=${MOCK_DEST}


# Proto Generation
## ADJUST FOR EACH SERVICE ##
proto: proto_core

proto_core: clean_pb proto_token proto_rftoken proto_errors
	protoc \
		--proto_path=$(PROTO_DIR) \
		--proto_path=${COMMON_PROTO_DIR} \
		--go_out=$(PB_DIR) \
		--go_opt=paths=source_relative,Mtoken_svc.proto=github.com/Streamfair/streamfair_token_svc/common_proto/TokenService/pb \
		--go-grpc_out=${PB_DIR} \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=${PB_DIR} \
		--grpc-gateway_opt=paths=source_relative \
		${PROTO_DIR}/*.proto
	statik -src=./$(SWAGGER_DIR) -dest=./doc

proto_token: clean_token_dir
	protoc \
		--proto_path=${PROTO_DIR} \
		--proto_path=${COMMON_PROTO_DIR} \
		--go_out=${PB_DIR} \
		--go_opt=paths=source_relative,Mtoken/rpc_create_token.proto=github.com/Streamfair/streamfair_token_svc/common_proto/TokenService/pb/token \
		--go-grpc_out=${PB_DIR} \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=${PB_DIR} \
		--grpc-gateway_opt=paths=source_relative \
		$(PROTO_DIR)/$(TOKEN_DIR)/*.proto

proto_rftoken: clean_refresh_token_dir
	protoc \
		--proto_path=${PROTO_DIR} \
		--proto_path=${COMMON_PROTO_DIR} \
		--go_out=${PB_DIR} \
		--go_opt=paths=source_relative,Mrefresh_token/rpc_create_refresh_token.proto=github.com/Streamfair/streamfair_token_svc/common_proto/TokenService/pb/refresh_token \
		--go-grpc_out=${PB_DIR} \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=${PB_DIR} \
		--grpc-gateway_opt=paths=source_relative \
		$(PROTO_DIR)/$(REFRESH_TOKEN_DIR)/*.proto

proto_errors:
	protoc \
		--proto_path=${COMMON_PROTO_ERROR_DIR} \
		--go_out=${COMMON_PROTO_ERROR_DIR} \
		--go_opt=paths=source_relative \
		${COMMON_PROTO_ERROR_DIR}/*.proto

clean_pb:
	rm -f $(PB_DIR)/*.go
	rm -f $(SWAGGER_DIR)/*.swagger.json

clean_token_dir:
	rm -f $(PB_DIR)/$(TOKEN_DIR)/*.go

clean_refresh_token_dir:
	rm -f $(PB_DIR)/$(REFRESH_TOKEN_DIR)/*.go


# Evans GRPC Client
evans:
	evans --host ${DB_HOST} --port ${GRPC_PORT} -r repl


# Testing
test:
	@if [ $(OUT) -eq 1 ]; then \
		go test -v -cover ${TEST_DIR} > ${TEST_FILE}; \
	else \
		go test -v -cover ${TEST_DIR} ; \
	fi

dbtest:
	@if [ $(OUT) -eq 1 ]; then \
		go test -v -cover ${DB_TEST_DIR} > ${DB_TEST_FILE}; \
	else \
		go test -v -cover ${DB_TEST_DIR} ; \
	fi

apitest:
	@if [ $(OUT) -eq 1 ]; then \
		go test -v -cover ${API_TEST_DIR} > ${API_TEST_FILE}; \
	else \
		go test -tags=-coverage -v -cover ${API_TEST_DIR} ; \
	fi

utiltest:
	@if [ $(OUT) -eq  1 ]; then \
		go test -v -cover -count=1 ${UTIL_TEST_DIR} > ${UTIL_TEST_FILE}; \
	else \
		go test -v -cover -count=1 ${UTIL_TEST_DIR} ; \
	fi

servertest:
	@if [ $(OUT) -eq  1 ]; then \
		go test -v -cover -count=1 ${SERVER_TEST_DIR} > ${SERVER_TEST_FILE}; \
	else \
		go test -v -cover -count=1 ${SERVER_TEST_DIR} ; \
	fi

coverage_html:
	go test -coverprofile=coverage.out ${TEST_DIR}
	go tool cover -html=coverage.out
	rm coverage.out

clean:
	rm -f *_tests.log


# PHONY Targets
.PHONY: network db_container createdb dropdb createmigration migrateup migrateup1 migratedown migratedown1 dbclean service_image service_container server sqlc mock proto_core proto_token proto_rftoken clean_pb clean_token_dir clean_refresh_token_dir evans test dbtest apitest utiltest servertest coverage_html clean down vault_container vault_image
