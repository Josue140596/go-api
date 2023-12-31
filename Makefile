postgres:
	docker run --name postgresbank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine
createdb:
	docker exec -it postgresbank createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgresbank dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc: 
	sqlc generate
test: 
	go test -v -cover ./...
serve: 
	go run main.go

.PHONY: postgres createdb dropdb migrateup sqlc serve