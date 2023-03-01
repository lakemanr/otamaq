postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
	
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root otamaq

dropdb:
	docker exec -it postgres15 dropdb otamaq

migrateup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/otamaq?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/otamaq?sslmode=disable" -verbose down 1

.PHONY: postgres createdb dropdb migrateup1 migratedown1
