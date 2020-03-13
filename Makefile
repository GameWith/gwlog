setup:
	@go get -u golang.org/x/lint/golint
	@go mod tidy

# go fmt
fmt:
	@go fmt ./...

# go vet
vet:
	@go vet ./...

# golint
lint:
	@golint -set_exit_status ./...

# go test
test:
	@go test -v ./...

check:
	@make fmt
	@make vet
	@make lint
	@make test

.PHONY: setup fmt vet lint test check