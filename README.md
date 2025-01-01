# setup

### 下記インストール

- docker (version 20.10.5)
- docker-compose (version 1.24.0)

### 立ち上げ

```
make build
make up
```

# air

cd backend/  
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b ./bin

# golangci-lint

cd backend/  
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.62.2

# gosec

cd backend/  
curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.21.4 ./bin

# xo

brew install xo/xo/xo

# grpc

PROTOC_VERSION=24.4
wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip
unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d backend/
[[]]
