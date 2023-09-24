DEFAULT: help

update: ## Update the go dependencies
	go mod tidy
	go mod vendor

godoc.run: ## Run godoc test
	CGO=0 go run cmd/godoc/server/main.go

godoc.client: ## Run godoc gRPC client
	CGO=0 go run cmd/godoc/client/main.go

godoc.compile: ## Compile gRPC files from proto
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    godoc-test/pb/godoc.proto

help: ## Show commands of the makefile (and any included files)
	@awk 'BEGIN {FS = ":.*?## "}; /^[0-9a-zA-Z_.-]+:.*?## .*/ {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)