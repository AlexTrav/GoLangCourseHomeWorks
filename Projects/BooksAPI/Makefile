APP_NAME=books-api

start:
	docker-compose up --build

down:
	docker-compose down

swagger:
	swag init -g cmd/app/main.go

migrate:
	go run cmd/app/main.go

logs:
	docker-compose logs -f

restart:
	docker-compose down
	docker-compose up --build