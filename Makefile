-include ./backend/internal/env/local.env

GO_TOOLCHAIN ?= go1.27.0
GOLANGCI_LINT ?= github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
GOVULNCHECK ?= golang.org/x/vuln/cmd/govulncheck@latest
OAPI_CODEGEN ?= github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.16.3

# ----------------------------
.PHONY: gomod

gomod:
	cd backend && go mod tidy && go mod vendor


# ----------------------------
.PHONY: stop
stop:
	docker compose --env-file ./backend/internal/env/local.env stop


# ----------------------------
.PHONY: removeAll
removeAll: stop
	docker compose --env-file ./backend/internal/env/local.env stop
	docker system prune -f
	sudo rm -rf db/engine/postgres/data
	sudo rm -rf db/engine/redis/data


# ----------------------------
.PHONY: build
build:
	docker compose --env-file ./backend/internal/env/local.env build
	# docker compose --env-file ./backend/internal/env/local.env build --no-cache

# ----------------------------
.PHONY: debug
debug:
	DEBUG_MODE=true docker compose --env-file ./backend/internal/env/local.env up

# ----------------------------
.PHONY: up
up:
	docker compose --env-file ./backend/internal/env/local.env up

# ----------------------------
.PHONY: reup
reup: stop build up

# ----------------------------
.PHONY: restart
restart: stop up

# ----------------------------
.PHONY: resetAll
resetAll: removeAll build up


# ----------------------------
.PHONY: gotest
gotest:
	cd backend && GOTOOLCHAIN=$(GO_TOOLCHAIN) go test ./...


# ----------------------------
.PHONY: lint
lint:
	cd ./backend && GOTOOLCHAIN=$(GO_TOOLCHAIN) go run $(GOLANGCI_LINT) run ./...

# ----------------------------
.PHONY: govulncheck
govulncheck:
	cd ./backend && GOTOOLCHAIN=$(GO_TOOLCHAIN) go run $(GOVULNCHECK) ./...

# ----------------------------
.PHONY: security
security: govulncheck

# ----------------------------
.PHONY: golint
golint: lint


# 指定ディレクトリ配下を再帰的に探してコンパイル ----------------------------
.PHONY: gen-grpc
gen-grpc:
	PATH=$(PWD)/backend/bin:$$PATH find backend/internal/1_framework/parameter/grpc -name "*.proto" -type f -exec \
		protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		{} \;

# ----------------------------
.PHONY: gen-openapi
gen-openapi:
	PATH=$(PWD)/backend/bin:$$PATH oapi-codegen \
	-generate types,server \
	-o backend/internal/1_framework/in/go-echo/openapi/api.gen.go \
	-package openapi \
	backend/internal/1_framework/in/go-echo/openapi/openapi.yaml 




# ----------------------------
.PHONY: install-tools
install-tools:
	# Create bin directory if it doesn't exist
	mkdir -p backend/bin
	
	# Install protoc compiler
	# GOBIN=$(PWD)/backend/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	# GOBIN=$(PWD)/backend/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	
	# Install other tools
	GOTOOLCHAIN=$(GO_TOOLCHAIN) GOBIN=$(PWD)/backend/bin go install $(GOLANGCI_LINT)
	GOTOOLCHAIN=$(GO_TOOLCHAIN) GOBIN=$(PWD)/backend/bin go install $(GOVULNCHECK)
	GOTOOLCHAIN=$(GO_TOOLCHAIN) GOBIN=$(PWD)/backend/bin go install github.com/air-verse/air@latest
	GOTOOLCHAIN=$(GO_TOOLCHAIN) GOBIN=$(PWD)/backend/bin go install $(OAPI_CODEGEN)
