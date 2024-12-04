run:	
	@go run .

init:
	@echo "Installing dependencies"
	@go mod tidy
	@sleep 1
	@echo "Running docker compose"
	@docker compose up -d db

clean:
	@docker compose down
	@sleep 1
	@echo "Removed postgres container"