postgres:
	docker run --name postgresdb -p 5432:5432 -e POSTGRES_USER=coinmaster -e POSTGRES_PASSWORD=coinpass -e POSTGRES_DB=coin -d postgres:11

createdb:
	docker exec -it postgresdb createdb --username=coinmaster coin

dropdb:
	docker exec -it postgresdb dropdb --username=coinmaster coin

migrate-up:
	migrate -path ./internal/infra/store/db/migration -database postgres://coinmaster:coinpass@localhost:5432/coin?sslmode=disable -verbose up

migrate-down:
	migrate -path ./internal/infra/store/db/migration -database postgres://coinmaster:coinpass@localhost:5432/coin?sslmode=disable -verbose down 1

sqlc:
	sqlc generate -f ./internal/infra/store/db/sqlc.yaml --experimental

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrate-up migrate-down sqlc test