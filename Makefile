include ./backend/internal/env/local.env

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
	DEBUG_MODE=true docker compose up


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
.PHONY: gosec
gosec:
	cd backend && ./bin/gosec  -exclude=G115 -conf ./bin/gosec.json ./...

# ----------------------------
.PHONY: golint
golint:
	cd ./backend ./bin/golangci-lint run ./...


# ----------------------------
.PHONY: staticcheck
staticcheck:
	cd ./backend && ./bin/staticcheck ./...


# ----------------------------
.PHONY: deadcode
deadcode:
	cd ./backend && ./bin/deadcode 




# ----------------------------
.PHONY: xo
xo:
	cd backend/internal/1_framework/db/postgres && xo schema postgres://user:user@localhost:15432/app?sslmode=disable


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

# # 指定ディレクトリに移動してからコンパイル ----------------------------
# .PHONY: gen-grpc
# gen-grpc:
# 	cd backend/internal/1_framework/parameter/grpc && \
# 	PATH=$(PWD)/backend/bin:$$PATH protoc \
# 		--go_out=. \
# 		--go_opt=paths=source_relative \
# 		--go-grpc_out=. \
# 		--go-grpc_opt=paths=source_relative \
# 		*.proto


# ----------------------------
.PHONY: install-tools
install-tools:
	# Create bin directory if it doesn't exist
	mkdir -p backend/bin
	
	# Install protoc compiler
	# GOBIN=$(PWD)/backend/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	# GOBIN=$(PWD)/backend/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	
	# Install other tools
	GOBIN=$(PWD)/backend/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOBIN=$(PWD)/backend/bin go install github.com/securego/gosec/v2/cmd/gosec@latest
	GOBIN=$(PWD)/backend/bin go install honnef.co/go/tools/cmd/staticcheck@latest
	GOBIN=$(PWD)/backend/bin go install github.com/air-verse/air@latest
	GOBIN=$(PWD)/backend/bin go install golang.org/x/tools/cmd/deadcode@latest
	
	


