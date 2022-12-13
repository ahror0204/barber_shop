DB_URL=postgresql://postgres:1@localhost:5432/users_db?sslmode=disable

CURRENT_DIR=$(shell pwd)

proto-gen:
	rm -rf genproto
	./scripts/gen-proto.sh ${CURRENT_DIR}

swag-init:
	swag init -g api/server.go -o api/docs

start:
	go run main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

.PHONY: start migrateup migratedown