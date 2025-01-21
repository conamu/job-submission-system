.PHONY: run-server
run-server:
	go run src/cmd/server/main.go

.PHONY: run-clients
run-clients:
	go run src/cmd/client/main.go

build-server:
	go build -o tmp/server src/cmd/server/main.go

build-clients:
	go build -o tmp/client-simulator src/cmd/client/main.go

build: build-clients build-server