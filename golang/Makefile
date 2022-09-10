postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=aaron -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username aaron --owner=aaron code_editor

dropdb:
	docker exec -it postgres14 dropdb code_editor

migrateup:
	migrate -path pkg/database/migrations/ -database "postgresql://aaron:secret@localhost:5432/code_editor?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/database/migrations/ -database "postgresql://aaron:secret@localhost:5432/code_editor?sslmode=disable" -verbose down


sqlc:
	sqlc generate

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, sqlc
