
deps:
	go mod tidy
	go mod vendor

lint:
	go golint ./...