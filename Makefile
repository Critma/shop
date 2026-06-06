.PHONY: all api-gateway auth message-generator up down stop start restart logs clean help

SHELL := /bin/bash

help:
	@echo "Доступные команды:"
	@echo "  make all             - Запустить все сервисы (api-gateway, auth, message-generator)"
	@echo "  make api-gateway     - Запустить api-gateway (API gateway + Nginx + PostgreSQL + Redis + Kafka)"
	@echo "  make auth            - Запустить auth сервис (gRPC авторизация + PostgreSQL)"
	@echo "  make message-generator - Запустить message-generator (генератор событий)"
	@echo "  make down            - Остановить и удалить все контейнеры"
	@echo "  make stop            - Остановить все контейнеры"
	@echo "  make start           - Запустить остановленные контейнеры"
	@echo "  make restart         - Перезапустить все сервисы"
	@echo "  make logs            - Показать логи всех сервисов"
	@echo "  make clean           - Удалить все volumes и контейнеры"

all: api-gateway auth message-generator

api-gateway:
	@echo "Запуск api-gateway..."
	$(MAKE) -C shop/api-gateway compose

auth:
	@echo "Запуск auth сервиса..."
	$(MAKE) -C shop/auth compose

message-generator:
	@echo "Запуск message-generator..."
	$(MAKE) -C shop/message-generator compose

stop:
	@echo "Остановка всех контейнеров..."
	$(MAKE) -C shop/api-gateway compose-stop || true
	$(MAKE) -C shop/auth compose-stop || true
	$(MAKE) -C shop/message-generator compose-stop || true

restart: down up

up: all

logs:
	@echo "Логи api-gateway:"
	docker-compose -p "shop-api" -f shop/api-gateway/deploy/docker-compose.yml logs -f || true
	@echo ""
	@echo "Логи auth:"
	docker-compose -p "auth" -f shop/auth/deploy/docker-compose.yml logs -f || true
	@echo ""
	@echo "Логи message-generator:"
	docker-compose -p "message-generator" -f shop/message-generator/deploy/docker-compose.yml logs -f || true

clean:
	@echo "Удаление всех volumes и контейнеров..."
	$(MAKE) -C shop/api-gateway compose-clear || true
	$(MAKE) -C shop/auth compose-clear || true
	$(MAKE) -C shop/message-generator compose-clear || true
