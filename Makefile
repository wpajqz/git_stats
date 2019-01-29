help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dep: ## Get build dependencies
	go get -v -u github.com/oxequa/realize
	go get -u github.com/swaggo/swag/cmd/swag

build: ## Build the app
	go build -ldflags "-X main.apiVersion=1.0  -X 'main.gitCommit=`git rev-parse HEAD`' -X 'main.built=`date`'"

test: ## Launch tests
	go test -v ./...

http: ## Run default http service with hot reload
	realize start --name "http-server"