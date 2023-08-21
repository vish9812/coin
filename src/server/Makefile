createdb:
	createdb --username=coinmaster coin

dropdb:
	dropdb --username=coinmaster coin

migrate-up:
	migrate -path ./internal/db/migration -database postgres://coinmaster:coinpass@localhost:5432/coin?sslmode=disable -verbose up

migrate-down:
	migrate -path ./internal/db/migration -database postgres://coinmaster:coinpass@localhost:5432/coin?sslmode=disable -verbose down 1

.PHONY: createdb dropdb migrate-up migrate-down