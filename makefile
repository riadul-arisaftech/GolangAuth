DB_URL=postgresql://root:2654@localhost:5432/simple_auth?sslmode=disable

network:
	docker network create auth-network

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=2654 -d postgres:15

createDB:
	docker exec -it postgres15 createdb --username=root --owner=root simple_auth

dropDB:
	docker exec -it postgres15 dropdb simple_auth

db_schema:
	dbml2sql --postgres -o doc/db_schema.sql doc/db.dbml

new_migration:
	migrate create -ext sql -dir src/db/migration -seq ${name}

migrateUp:
	migrate -path src/db/migration -database "${DB_URL}" -verbose up ${v}

migrateDown:
	migrate -path src/db/migration -database "${DB_URL}" -verbose down ${v}

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: network postgres createDB dropDB db_schema new_migrations migrateUp migrateDown sqlc test server