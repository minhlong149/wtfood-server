.PHONY: run db

db:
	docker-compose up db

run:
	go run cmd/server/main.go
