-include ./backend/internal/env/local.env

GO_TOOLCHAIN ?= go1.27.0
GOLANGCI_LINT ?= github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
GOVULNCHECK ?= golang.org/x/vuln/cmd/govulncheck@latest
OAPI_CODEGEN ?= github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.16.3
override PROJECT_ROOT := $(abspath $(dir $(firstword $(MAKEFILE_LIST))))
override POSTGRES_DATA_DIR := $(PROJECT_ROOT)/db/engine/postgres/data
override REDIS_DATA_DIR := $(PROJECT_ROOT)/db/engine/redis/data
DATA_CLEANER_IMAGE ?= alpine:3.22

# ----------------------------
.PHONY: gomod

gomod:
	cd backend && go mod tidy && go mod vendor


# ----------------------------
.PHONY: stop
stop:
	docker compose --env-file ./backend/internal/env/local.env stop


# 破壊的操作: このComposeプロジェクトとローカルDBデータだけを削除する。
# 誤実行防止のため DESTROY_LOCAL_DATA=yes の明示指定を必須とする。
# ----------------------------
.PHONY: destroy-local-data
destroy-local-data:
	@if [ "$(DESTROY_LOCAL_DATA)" != "yes" ]; then \
		echo "Destructive operation blocked."; \
		echo "Run: make destroy-local-data DESTROY_LOCAL_DATA=yes"; \
		exit 1; \
	fi
	@set -eu; \
	for data_dir in "$(POSTGRES_DATA_DIR)" "$(REDIS_DATA_DIR)"; do \
		resolved_dir="$$(realpath -m "$$data_dir")"; \
		case "$$resolved_dir" in \
			"$(POSTGRES_DATA_DIR)"|"$(REDIS_DATA_DIR)") ;; \
			*) echo "Refusing to delete unexpected path: $$resolved_dir"; exit 1 ;; \
		esac; \
	done; \
	docker compose \
		--project-directory "$(PROJECT_ROOT)" \
		--env-file "$(PROJECT_ROOT)/backend/internal/env/local.env" \
		down --remove-orphans --volumes; \
	for data_dir in "$(POSTGRES_DATA_DIR)" "$(REDIS_DATA_DIR)"; do \
		resolved_dir="$$(realpath -m "$$data_dir")"; \
		if [ -d "$$resolved_dir" ]; then \
			echo "Deleting local data under $$resolved_dir"; \
			docker run --rm --volume "$$resolved_dir:/data" $(DATA_CLEANER_IMAGE) \
				sh -c 'find /data -mindepth 1 -maxdepth 1 -exec rm -rf -- {} +'; \
		fi; \
	done

# 旧ターゲットは意図が不明瞭なため廃止。上記の明示的なターゲットを使用する。
.PHONY: removeAll
removeAll:
	@echo "removeAll was removed because it was unsafe."
	@echo "Run: make destroy-local-data DESTROY_LOCAL_DATA=yes"
	@exit 1


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

# 破壊的操作: ローカルDBデータを削除してbuild後に再起動する。
# ----------------------------
.PHONY: reset-local-data
reset-local-data:
	$(MAKE) destroy-local-data DESTROY_LOCAL_DATA=$(DESTROY_LOCAL_DATA)
	$(MAKE) build
	$(MAKE) up



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
