postgresinit:
	sudo docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	sudo docker exec -it postgres15 psql

createdb:
	sudo docker exec -it postgres15 createdb --username=root --owner=root todolist

dropdb: 
	sudo docker exec -it postgres15 dropdb notes

createamigrationfile:
	../bin/migrate create -ext sql -dir ./db/migrations add_notes

migrateup:
	../bin/migrate -path db/migrations -database "postgresql://root:password@localhost:5433/todolist?sslmode=disable" -verbose up

migratedown:
	../bin/migrate -path db/migrations -database "postgresql://root:password@localhost:5433/todolist?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown