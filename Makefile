.PHONY: checkhash
checkhash: ## checkhash
	@curl -X POST -H "Content-Type: application/json" -d '{"payload": "your_payload_here"}' http://localhost:8080/checkhash

.PHONY: gethash
gethash: ## gethash
	@curl -X POST -H "Content-Type: application/json" -d '{"payload": "your_payload_here"}' http://localhost:8080/gethash

.PHONY: createhash
createhash: ## createhash
	@curl -X POST -H "Content-Type: application/json" -d '{"payload": "your_payload_here"}' http://localhost:8080/createhash

.PHONY: test
test: ## run tests
	@go test -v -cover hashing/...
	@go test -v -cover gateway/...

.PHONY: deploy
deploy: ## deploy
	@docker compose up -d --build

.PHONY: destroy
destroy: ## destroy
	@docker compose down

.DEFAULT_GOAL := help
.PHONY: help
help: ## show help
	@awk '{ if (NF == 2 && $$1 == "include") { while ((getline line < $$2) > 0) print line ; close($$2) } else print }' $(firstword $(MAKEFILE_LIST)) \
		| grep -E '^[a-zA-Z_-]+:.*?## .*$$' \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'