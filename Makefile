up:
	docker-compose build postgres
	docker-compose up --build -d postgres
	docker-compose run --rm wait -c postgres:5432 -t 15
	docker-compose up --build deck-api

down:
	docker-compose down

build:
	go build -o deck-api

test:
	go test ./...