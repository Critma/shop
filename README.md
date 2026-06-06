# Монорепозиторий магазина (Go + gRPC + Kafka + Redis + PostgreSQL)

Монорепозиторий, содержащий:
- **api-gateway** — API gateway (Go, Gin, Huma, PostgreSQL, Redis, Kafka).
- **auth** — gRPC микросервис авторизации и управления сессиями (Go, gRPC, JWT, PostgreSQL).
- **message-generator** — микросервис, который генерирует изменение цены товара и отправляет их в Kafka.
- **contract** — общий модуль с gRPC контрактами (`.proto`) и сгенерированным кодом.

## Общая структура репозитория

```
.
├── shop/
│   ├── api-gateway/      # Go backend API gateway (Gin, Huma, Redis, Kafka)
│   ├── auth/             # gRPC сервис авторизации (JWT, PostgreSQL)
│   ├── message-generator/# Генератор событий в Kafka
│   └── contract/         # Общие gRPC контракты и генерация кода
└── README.md
```

Подробнее по каждому модулю:

- `shop/api-gateway/README.MD` — документация по API gateway (запуск и тд).
- `shop/contract/` — описание gRPC контрактов, генерация кода для Go.

## Быстрый старт (3 docker-compose)

```bash
make all
```

После успешного старта:

- **API gateway, OpenAPI (Go)**: `http://localhost/api/v1`
- **Swagger**: `http://localhost/swagger/`
- **pgAdmin**: `http://localhost/admin/`
- **Nginx status**: `http://localhost/status`
- **Auth gRPC сервис**: порт 9090 (внутри docker-сети), наружу не проброшен
- **Kafka (Redpanda)**: `localhost:9092` (внешний), `kafka:9092` (внутри docker-сети)
- **Redis**: `localhost:6379` (внешний), `redis:6379` (внутри docker-сети)
- **Kafka UI**: `http://localhost:9095`

Дополнительно можно запустить только PostgreSQL:

```bash
make compose-postgres
```

## Взаимодействие сервисов

- **api-gateway** (Go) — REST API gateway, интегрирован с `auth` по gRPC, потребляет сообщения из Kafka, использует Redis для кэширования.
- **auth** — общий gRPC-сервис авторизации (JWT, Login, Register, ChangePassword, RecoverPassword, ValidateToken).
- **message-generator** — генерирует периодические изменения товаров и отправляет события в Kafka.

**Объединённая авторизация:** все API gateway экземпляры используют один auth-service. Регистрация/авторизация в одном экземпляре даёт валидный токен для всех.

- Контракты в `shop/contract/proto/auth/v1/auth.proto`, сгенерированный код — в `shop/api-gateway/gen/auth/v1/...` и `shop/auth/gen/auth/v1/...`.

## Тестирование и разработка

- Для изменения gRPC контрактов:
  1. Изменение `.proto` в `shop/contract/proto/auth/v1/`.
  2. Генерация код: `make gen-grpc` в `shop/auth/` или `shop/api-gateway/`.
  3. Обновление в `auth` и `api-gateway`.

- Генерация JWT секрета (auth):
  ```bash
  make gen-jwt_secret
  ```

- Генерация сертификатов (api):
  ```bash
  make gen-cert
  ```

## Полезные ссылки внутри репозитория

- `shop/api-gateway/README.MD` — документация по API gateway.
- `shop/contract/` — работа с контрактами и генерацией gRPC-кода.
- `shop/auth/` — реализация сервиса авторизации, JWT и gRPC-сервер.
- `shop/message-generator/` — реализация генератора событий и отправки в Kafka.
- `shop/api-gateway/deploy/nginx/conf.d/default.conf` — конфигурация Nginx.
- `shop/api-gateway/web/static/` — статический веб.

## Конфигурация окружения

Конфигурация каждого сервиса загружается из переменных окружения с поддержкой значений по умолчанию.

### api-gateway

- `SERVICE_HOST` — хост API gateway (default: `localhost`)
- `POSTGRES_*` — параметры подключения к PostgreSQL
- `KAFKA_READER_ADDR` — адрес Kafka reader (default: `kafka:9092`)
- `KAFKA_READER_TOPIC` — топик для чтения (default: `product-updated`)
- `KAFKA_READER_GROUP_ID` — group ID для Kafka consumer (default: `api-gateway-local`)
- `REDIS_ADDR` — адрес Redis (default: `redis:6379`)
- `AUTH_HOST` — хост gRPC auth сервиса (default: `localhost`)
- `AUTH_PORT` — порт gRPC auth сервиса (default: `9090`)

### auth

- `AUTH_HOST` — хост gRPC сервиса (default: `localhost`)
- `AUTH_PORT` — порт gRPC сервиса (default: `9090`)
- `POSTGRES_*` — параметры подключения к PostgreSQL
- `TOKEN_TTL` — время жизни JWT токена (default: `1h`)
- `TOKEN_SECRET` — секрет для JWT (считывается из `/run/secrets/jwt_secret` или генерируется)

### message-generator

- `TICK_INTERVAL` — интервал генерации событий (default: `10s`)
- `POSTGRES_*` — параметры подключения к PostgreSQL
- `KAFKA_WRITER_ADDR` — адрес Kafka writer (default: `host.docker.internal:9093`)

## Базы данных

### api-gateway

- Таблицы: `address`, `client`, `supplier`, `image`, `product`
- Схема: `shop/api-gateway/deploy/postgres/docker-entrypoint-initdb.d/init.sql`

### auth

- Таблица: `users` с триггером для автоматического обновления `updated_at`
- Схема: `shop/auth/deploy/postgres/docker-entrypoint-initdb.d/init.sql`

## Журналы и мониторинг

- Nginx access/error логи: `/var/log/nginx/access_http.log`, `/var/log/nginx/error_http.log`
- Статус Nginx: `http://localhost/status`
- Статус PostgreSQL: проверяется через `pg_isready`
- Журналы контейнеров: `docker logs <container_name>`

## Статические файлы и демо-клиенты

- `index.html` — главная страница с ссылками на API документацию и демо-клиенты
- `sse-echo.html` — клиент для SSE (Server-Sent Events) для просмотра обновлений товаров
- `websocket-echo.html` — клиент для WebSocket для просмотра обновлений товаров

## Используемые технологии

### api-gateway
- Go 1.26
- Gin (веб-фреймворк)
- Huma v2 (веб-фреймворк, OpenAPI генерация)
- PostgreSQL (база данных)
- Redis (кэширование)
- Kafka (event streaming)
- gRPC (взаимодействие с auth сервисом)
- JWT (аутентификация)
- Zerolog (логирование)

### auth
- Go 1.26
- gRPC (сервис авторизации)
- PostgreSQL (база данных)
- JWT (токены)
- bcrypt (хеширование паролей)
- Uber Fx (инъекция зависимостей)
- Zap (логирование)

### message-generator
- Go 1.26
- PostgreSQL (чтение данных)
- Kafka (отправка событий)
- Uber Fx (инъекция зависимостей)
- slog (логирование)
