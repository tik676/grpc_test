# README

## Описание

Простой пример сервиса задач (TODO) на Go с использованием gRPC и PostgreSQL.  
Сервис умеет создавать, изменять, удалять и получать задачи.  
В gRPC-методы обёрнуты базовые CRUD-операции через proto-файл.
В ближайшем будущем события будут отправлятся в kafka.



## Структура проекта

- `cmd/main.go` — точка входа, запуск сервера.
- `internal/api/todo.proto` — описание gRPC-сервиса и сообщений.
- `internal/delivery/grpc/pb/` — сгенерированные protobuf-файлы.
- `internal/delivery/grpc/server.go` — реализация gRPC методов.
- `internal/domain/models.go` — структуры данных (Task).
- `internal/infrastructure/postgres/postgres.go` — работа с базой данных.
- `internal/usecase/usecase.go` — бизнес-логика.
- `.env` — переменные среды (коннект к базе).
- `docker-compose.yml` — деплой Postgres и backend.
- `Dockerfile` — билд Go-контейнера.
- `init.sql` — автоматическое создание таблицы задач.
- `go.mod`, `go.sum` — зависимости.



## Как запустить

1. **Заполни `.env`** (пример уже есть):
    ```
    HOST=postgres
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    POSTGRES_DB=app_db
    ```
2. **Собери и подними сервисы**:
    ```
    docker-compose up --build
    ```
    Это развернёт backend и базу. Таблица создаётся автоматически через `init.sql`.

3. **Проверь работу gRPC на порту `8080`.**


## Как тестировать

- Используй [grpcurl](https://github.com/fullstorydev/grpcurl) или [BloomRPC](https://bloomrpc.com/) для ручных gRPC-запросов.
- Импортируй ваш proto-файл для работы через GUI-клиент.



## Методы сервиса

- **CreateTask** — создать задачу.
- **ListTasks** — получить список задач.
- **EditTask** — изменить задачу.
- **DeleteTask** — удалить задачу.


## Прочее

- Код без лишних зависимостей, структурирован по слоям.
- PostgreSQL требует образ `postgres:14-alpine`.
- Все основные настройки — через `.env` и docker-compose.

