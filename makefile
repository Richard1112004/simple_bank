postgres:
	docker run -d -p 5433:5432 --name my-postgres-12 -e POSTGRES_PASSWORD=mysecretpassword postgres
createdb:
	docker exec -it my-postgres-12 createdb --username=postgres --owner=postgres golangpro
dropdb:
	docker exec -it my-postgres-12 dropdb --username=postgres golangpro
migrateup:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5433/golangpro?sslmode=disable" up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5433/golangpro?sslmode=disable" down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test