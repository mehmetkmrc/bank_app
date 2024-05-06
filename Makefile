postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root new_banka

dropdb:
	docker exec -it postgres15 dropdb new_banka

migrateup:
	migrate -path db/migration -database "postgresql://postgres:ITM-2020@localhost:5432/new_banka?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:ITM-2020@localhost:5432/new_banka?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server