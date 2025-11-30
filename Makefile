run-local:
	go run cmd/web/main.go

run-docker: 
	docker-compose up --build