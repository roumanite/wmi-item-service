setup:
	@echo "--- Setup and generate config yaml files ---"
	@cp config.sample.yaml config.yaml

server:
	@echo "--- Run httpd server ---"
	@go run cmd/server/main.go

migrate-up:
	@echo "--- Run db migration ---"
	@go run cmd/db/migrate/main.go --direction=up --step=$(step) --schema=$(schema)

migrate-down:
	@echo "--- Run db migration ---"
	@go run cmd/db/migrate/main.go --direction=down --step=$(step) --schema=$(schema)

migrate-create-pg:
	@echo "--- Create db file migration ---"
	@migrate create -dir=db/migrations -ext sql -format "unix" $(name)
