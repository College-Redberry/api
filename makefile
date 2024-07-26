include .env
export
run: .env
	go run .

run_db:
	docker run -d --name postgres-docker -p 5433:5432 -e POSTGRES_PASSWORD=Redberry2024 postgres:bullseye