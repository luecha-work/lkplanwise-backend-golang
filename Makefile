DB_URL=postgresql://planwise:Password@1@localhost:5432/lkplanwisedb?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name lkplanwise-postgres -p 5432:5432 -e POSTGRES_USER=planwise -e POSTGRES_PASSWORD=Password@1 -d postgres:latest
	
# docker run --name lkplanwise-postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=planwise -e POSTGRES_PASSWORD=Password@1 -d postgres:latest

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=Password@1 -d mysql:8

createdb:
	docker exec -it lkplanwise-postgres createdb --username=planwise --owner=planwise lkplanwisedb

dropdb:
	docker exec -it lkplanwise-postgres dropdb lkplanwisedb

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --lkplanwise-postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/lkplanwisedb/db/sqlc Store

	mockgen -package mockwk -destination worker/mock/distributor.go github.com/techschool/lkplanwisedb/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=lkplanwisedb \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

# proto:
# 	rm -f pb/*.go
# 	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
# 	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
# 	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
# 	proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl
	

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis