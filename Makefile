DB_NAME = example
DB_PORT = 3307
MYSQL_ROOT_PASSWORD = pass

# hello:
# 	@if [ -z "$(h)" ]; then \
# 		echo "hhhh"; \
# 	else \
# 		echo "not found"; \
# 	fi

# Start
.PHONY: up
up:
	docker compose up -d

# Shutdown
.PHONY: down
down:
	docker compose down

# restart
.PHONY: restart
restart:
	docker compose restart

# Clean all the resource created from this app
.PHONY: clean
clean:
	docker compose down
	docker rmi onion-architecture-go-app:latest
	docker volume rm onion-architecture-go_mysql
	rm -rf migrations

# Unittest
.PHONY: test
test:
	go test -v ./...

# Check your code
.PHONY: check
check:
	go vet

# Run this if there is any file changes in ports folder
# Command ex: make mock file=<filename> (filename need to be exist under the ports folder)
.PHONY: mock
mock:
	mockgen -source=./ports/$(file) -destination=./mocks/$(file)

# Init the db from empty state. Run this command after you first run, want to start again
# when you kill all the resource or clean the db data and migration files
.PHONY: migrate-init
migrate-init:
	@atlas migrate diff init \
 		--dir "file://migrations" \
		--to "file://schema.hcl" \
  	--dev-url "docker://mysql/8/$(DB_NAME)"
	@atlas migrate apply \
  	--url "mysql://root:$(MYSQL_ROOT_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)" \
  	--dir "file://migrations"

# Lookup the current db migrate status
.PHONY: migrate-status
migrate-status:
	@atlas migrate status \
  	--url "mysql://root:$(MYSQL_ROOT_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)" \
  	--dir "file://migrations"

# Commit the schema, commit sql will be stored into migrations folder (auto generated if not exist)
# Command ex: migrate-commit name=<commit name>
.PHONY: migrate-commit
migrate-commit:
	@atlas migrate diff $(name) \
 		--dir "file://migrations" \
		--to "file://schema.hcl" \
  	--dev-url "docker://mysql/8/$(DB_NAME)"

# Apply the schema changes to the db
.PHONY: migrate-up
migrate-up:
	@atlas migrate apply \
  	--url "mysql://root:$(MYSQL_ROOT_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)" \
  	--dir "file://migrations"

# Revert to the previous version
.PHONY: migrate-down
migrate-down:
	@atlas migrate down \
  	--url "mysql://root:$(MYSQL_ROOT_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)" \
  	--dir "file://migrations" \
		--dev-url "docker://mysql/8/$(DB_NAME)"

# Clean all the db's data, schema and migration files
.PHONY: migrate-clean
migrate-clean:
	@atlas schema clean -u "mysql://root:$(MYSQL_ROOT_PASSWORD)@:$(DB_PORT)/$(DB_NAME)"
	@rm -rf migrations

# Reverting all schema changes
.PHONY: migrate-reload
migrate-reload:
	@atlas schema clean -u "mysql://root:$(MYSQL_ROOT_PASSWORD)@:$(DB_PORT)/$(DB_NAME)"
	@atlas migrate apply -u "mysql://root:$(MYSQL_ROOT_PASSWORD)@:$(DB_PORT)/$(DB_NAME)"
