.PHONY: install,run

install:
	cd /go/src/golang-echo-docker-compose/ && go mod download



run:
	go run ./cmd/main.go
