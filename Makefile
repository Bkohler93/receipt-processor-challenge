run:
	@docker compose up -d db
	@echo "Waiting for PostgreSQL to be ready..."
	@until docker exec -it $(shell docker compose ps -q db) pg_isready -U dev_user -d dev_db >/dev/null 2>&1; do \
		echo "Waiting for the database..."; \
		sleep 1; \
	done
	@echo "Database is ready!"
	@echo "Running Go application..."
	go run .

init:
	go mod tidy

stop-db:
	docker compose down