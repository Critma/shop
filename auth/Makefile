CONTRACT_DIRECTORY=../contract
PROTO_ROOT=$(CONTRACT_DIRECTORY)/proto
PROTO_PATH=$(PROTO_ROOT)/auth/v1/auth.proto
PROTO_GEN_GO_OUT=gen

COMPOSE_PATH=deploy/docker-compose.yml
APP_PATH=cmd/service/main.go

generate-grpc:
	mkdir -p $(PROTO_GEN_GO_OUT)
	protoc -I $(PROTO_ROOT) $(PROTO_PATH) --go_out=$(PROTO_GEN_GO_OUT) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_GEN_GO_OUT) --go-grpc_opt=paths=source_relative

compose:
	docker compose -p "auth" -f $(COMPOSE_PATH) up -d --build

run:
	go run $(APP_PATH)