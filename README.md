# setup

### 下記インストール

- docker (version 20.10.5)
- docker-compose (version 1.24.0)

### 立ち上げ

```
make build
make up
```

# xo

brew install xo/xo/xo

# clean architecture

<img src="doc/ca.png">

# ddd (value object)

### backend/internal/4_domain に凝集。

## [primitive_object](backend/internal/4_domain/primitive_object)

string や int の基本的なデータ型に対するロジックを定義。  
文字列の長さチェックや nil 判定など

## [value_object](backend/internal/4_domain/value_object)

primitive_object を利用して個別のデータ型を定義。  
メールアドレスフォーマットチェックなど

## [struct_object](backend/internal/4_domain/struct_object)

value_object を複数組み合わせてドメインモデルを定義。

# grpc server client

## [grpc server](backend/internal/1_framework/in/go-grpc/person.go)

## [grpc client](backend/internal/1_framework/out/grpc_client/logic.go)

# rest server

## [rest server](backend/internal/1_framework/in/go-echo/v1/person/viaGRPC.go)

http request を grpc に変換して grpc サーバーにリクエストを送信。

# auth0 client

[auth0 client](backend/internal/1_framework/out/auth0_client/logic.go)

# microservice data candidate

### [マイクロサービス間の共通データ項目](backend/internal/1_framework/parameter/grpc/person.proto)

マイクロサービス間の共通データ項目を定義。

# dev environment

hot reload  
debug mode
