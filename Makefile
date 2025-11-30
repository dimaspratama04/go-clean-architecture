run-api-local:
	go run cmd/web/main.go

run-with-docker: 
	docker-compose up --build