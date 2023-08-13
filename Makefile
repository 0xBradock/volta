postgres15:
	podman run --name pg15 -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -d postgres:15.4-alpine

createdb:
	podman exec -it pg15 createdb --username=root --owner=root volta

dropdb:
	podman exec -it pg15 psql -U root volta

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/volta?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/volta?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres15 createdb dropdb postgres15 migrateup migratedown sqlc
