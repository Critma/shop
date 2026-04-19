ENV_PATH ?=
include $(ENV_PATH)
COMPOSE_PATH=deploy/docker-compose.yml
BUILD_PATH=build/api
EXE_PATH=$(BUILD_PATH)/shopapi
SRC_PATH=cmd/api/*.go

compose:
# 	docker-compose -p "shop-api" -f $(COMPOSE_PATH) --env-file $(ENV_PATH) up -d
	docker-compose -p "shop-api" -f $(COMPOSE_PATH) $(if $(ENV_PATH),--env-file $(ENV_PATH)) up -d

compose-postgres:
	docker-compose -p "shop-api" -f $(COMPOSE_PATH) --env-file $(ENV_PATH) up -d postgres

compose-clear:
	docker compose -f $(COMPOSE_PATH) --env-file $(ENV_PATH) down -v

# run local
run:
	mkdir -p $(BUILD_PATH)
	go build -o $(EXE_PATH) $(SRC_PATH)
	./$(EXE_PATH)

.PHONY: run compose