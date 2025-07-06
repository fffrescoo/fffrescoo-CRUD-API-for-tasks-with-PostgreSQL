# Путь к базе
DB_DSN := "postgres://appuser:12345@localhost:5432/taskdb?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создать новую миграцию
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Применить миграции (создаст таблицу)
migrate:
	$(MIGRATE) up

# Откатить миграции (удалит таблицу)
migrate-down:
	$(MIGRATE) down

# Запустить приложение
run:
	go run cmd/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
lint:
	golangci-lint run