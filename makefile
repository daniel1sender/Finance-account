.PHONY: format
format:
	goimports -l -w ./
	
.PHONY: lint
lint: format
	golangci-lint run

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build main.go

.PHONY: build-image
build-image:
	docker build -f dockerfile -t senderdan/desafio .

.PHONY: run-local
run-local:
	./run.sh
