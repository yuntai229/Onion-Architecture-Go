DB_NAME = example
DB_PORT = 3307
MYSQL_ROOT_PASSWORD = pass

# hello:
# 	@if [ -z "$(h)" ]; then \
# 		echo "hhhh"; \
# 	else \
# 		echo "not found"; \
# 	fi

up:
	docker compose up -d

down:
	docker compose down

clean:
	docker compose down
	docker rmi onion-architecture-go-app:latest
	docker volume rm onion-architecture-go_mysql
	rm -rf migrations

test:
	go test -v ./...

check:
	go vet

mock:
	mockgen -source=./ports/$(file) -destination=./mocks/$(file)

migrate-init:
	@atlas migrate diff init \
 		--dir "file://migrations" \
		--to "file://schema.hcl" \
  	--dev-url "docker://mysql/8/$(DB_NAME)"
	@atlas migrate apply \
  	--url "mysql://root:$(MYSQL_ROOT_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)" \
  	--dir "file://migrations"

migrate-status:
	@atlas migrate status \
  	--url "mysql://root:$(MYSQL_ROOT_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)" \
  	--dir "file://migrations"

migrate-commit:
	@atlas migrate diff $(name) \
 		--dir "file://migrations" \
		--to "file://schema.hcl" \
  	--dev-url "docker://mysql/8/$(DB_NAME)"

migrate-up:
	@atlas migrate apply \
  	--url "mysql://root:$(MYSQL_ROOT_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)" \
  	--dir "file://migrations"

migrate-down:
	@atlas migrate down \
  	--url "mysql://root:$(MYSQL_ROOT_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)" \
  	--dir "file://migrations" \
		--dev-url "docker://mysql/8/$(DB_NAME)"

migrate-clean:
	@atlas schema clean -u "mysql://root:$(MYSQL_ROOT_PASSWORD)@:$(DB_PORT)/$(DB_NAME)"
	@rm -rf migrations

# Reverting all schema changes
migrate-reload:
	@atlas schema clean -u "mysql://root:$(MYSQL_ROOT_PASSWORD)@:$(DB_PORT)/$(DB_NAME)"
	@atlas migrate apply -u "mysql://root:$(MYSQL_ROOT_PASSWORD)@:$(DB_PORT)/$(DB_NAME)"
