init:
	cp compose.override.yml.sample conpose.override.yml

run:
	go run ./cmd

test:
	go test ./...

build:
	docker build -o goexe .

fmt:
	go fmt ./...


.PHONY: init run test build