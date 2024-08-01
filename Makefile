include .env
export
run: .env
	go run .

run_api:
	docker compose build
	docker compose up

create_db:
	docker run -d --name postgres-docker -p 5433:5432 -e POSTGRES_PASSWORD=Redberry2024 postgres:bullseye

run_db:
	docker start postgres-docker

stop_db:
	docker stop postgres-docker

create-migration:
	migrate create -ext sql -dir migrations -seq $(name)

execute-migrations-up:
	migrate -database ${DB_STRING} -path migrations up

execute-migrations-down:
	migrate -database ${DB_STRING} -path migrations down
