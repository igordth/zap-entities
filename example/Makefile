# List off commands
list:
	@grep '^[^[:space:]].*:' Makefile | grep -Eo '^[^:]*'

# Run rotation example
run-stdout:
	@cd .. && go run ./example/stdout

# Run rotation example
run-rotation:
	@cd .. && go run ./example/rotation

# Run multi example
run-multi:
	@cd .. && go run ./example/multi

# Run elk example
run-elk:
	@cd .. && go run ./example/elk