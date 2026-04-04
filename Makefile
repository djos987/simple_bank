postgres:
	docker run -p 5432:5432 --name postgres18 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18-alpine

createdb:
	docker exec -it postgres18 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres18 dropdb --username=root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

DOWN_STEP ?= 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down $(DOWN_STEP)

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mocdb -destination db/mock/store.go github.com/djos987/simple_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test mock server migratedown1