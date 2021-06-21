TEST_COMMAND= GIN_MODE=test go test -v -cover ./src/controller/...

start:
	POSTGRES_HOST=localhost go run ./src/main.go
setEnv:
	cp .env.example .env
upDB:
	docker compose up db migrate
up:
	docker compose up -d
down:
	docker compose down
test:
	$(TEST_COMMAND)
testAll:
	go clean -testcache && $(TEST_COMMAND)
lint:
	golangci-lint run ./src/...
