postgres:
	docker run --name my-postgres -p 5000:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

create-db:
	docker exec -it my-postgres createdb --username=admin --owner=admin simple-bank

drop-db:
	docker exec -it my-postgres dropdb simple-bank

migrate-up:
	migrate -path ./db/migration -database "postgresql://admin:admin@localhost:5000/simple-bank?sslmode=disable" -verbose up

migrate-down:
	migrate -path ./db/migration -database "postgresql://admin:admin@localhost:5000/simple-bank?sslmode=disable" -verbose down

sqlc:
# require cmd instead of powershell on windows:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: all