postgres:
	docker run --name db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345 -p 9876:5432 -d postgres

createdb:
	docker exec -it db createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it db dropdb simple_bank 

migrateup:
	migrate -path db/migration -database "postgresql://root:12345@127.0.0.1:9876/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:12345@127.0.0.1:9876/simple_bank?sslmode=disable" -verbose down


sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc