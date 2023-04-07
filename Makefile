run_postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

start_postgres:
	docker start postgres15
	
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root otamaq

dropdb:
	docker exec -it postgres15 dropdb otamaq

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/otamaq?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/otamaq?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/otamaq?sslmode=disable" -verbose down 1

test:
	go test -v --cover ./...

.PHONY: run_postgres start_postgres createdb dropdb migrateup migrateup1 migratedown1 test
