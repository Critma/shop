COMPOSE_PATH=deploy/docker-compose.yml
BUILD_PATH=build/service
EXE_PATH=$(BUILD_PATH)/msg_gen
SRC_PATH=cmd/service/*.go

compose:
	docker-compose -p "message-generator" -f $(COMPOSE_PATH) up -d --build

compose-clear:
	docker compose -f $(COMPOSE_PATH) down -v

# run local
run:
	mkdir -p $(BUILD_PATH)
	go build -o $(EXE_PATH) $(SRC_PATH)
	./$(EXE_PATH)

.PHONY: run compose