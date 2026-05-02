COMPOSE_PATH=deploy/docker-compose.yml
BUILD_PATH=build/api
EXE_PATH=$(BUILD_PATH)/shopapi
SRC_PATH=cmd/api/*.go

PROTO_ROOT=../contract/proto
PROTO_PATH=$(PROTO_ROOT)/auth/v1/auth.proto
PROTO_GEN_GO_OUT=gen

compose:
	docker-compose -p "shop-api" -f $(COMPOSE_PATH) up -d

compose-api:
	docker-compose -p "shop-api" -f $(COMPOSE_PATH) up -d api

compose-postgres:
	docker-compose -p "shop-api" -f $(COMPOSE_PATH) up -d postgres

compose-clear:
	docker compose -f $(COMPOSE_PATH) down -v

#  grpc
generate-grpc:
	mkdir -p $(PROTO_GEN_GO_OUT)
	protoc -I $(PROTO_ROOT) $(PROTO_PATH) --go_out=$(PROTO_GEN_GO_OUT) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_GEN_GO_OUT) --go-grpc_opt=paths=source_relative

# certificate
generate-cert:
	cd deploy/nginx/ssl && openssl req -x509 -nodes -days 730 -newkey rsa:2048 -config openssl.cnf -sha256 -out shopapi.local.crt

# run local
run:
	mkdir -p $(BUILD_PATH)
	go build -o $(EXE_PATH) $(SRC_PATH)
	./$(EXE_PATH)

.PHONY: run compose