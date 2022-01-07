project_name=discgolfapitest
build_path=bin/$(project_name)

run: docs server

.PHONY: docs
docs:
	echo "building docs"
	swag init

.PHONY: server
server:
	echo "starting server"
	go run main.go

# compiled build and run
build-and-run: build run-build

.PHONY: build
build: docs
	echo "building go code"
	go build -o $(build_path)

.PHONY: run-build
run-build:
	echo "running go code"
	./$(build_path)
