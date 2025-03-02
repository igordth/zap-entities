# List off commands
list:
	@grep '^[^[:space:]].*:' Makefile | grep -Eo '^[^:]*'

# Run rotation example
run-stdout:
	@cd .. && go run ./example/stdout

# Run rotation example
run-rotation:
	@cd .. && go run ./example/rotation

# Run rgxp example
run-rgxp:
	@cd .. && go run ./example/rgxp

# Run elk example
run-elk:
	@cd .. && go run ./example/elk

# Run clickhouse example
run-clickhouse:
	docker compose -p zap-core-clickhouse down -v
	docker compose -p zap-core-clickhouse --project-directory ./clickhouse/build up -d
	sleep 3s # wait docker
	cd .. && go run ./example/clickhouse

# Run file example
run-file:
	@cd .. && go run ./example/file