.PHONY: build run seed db

build:
	go build -o bin/go-movies main.go

run:
	go run main.go

seed:
	go run ./import/install.go

db:
	docker compose up -d

wait-db:
	@echo "Waiting for postgres..." && \
	until docker compose exec -T db pg_isready -q; do sleep 1; done && \
	echo "Postgres is ready."

setup: db wait-db seed
