## コンパイルコマンド

protoc \
 --go_out=. \
 --go_opt=paths=source_relative \
 --go-grpc_out=. \
 --go-grpc_opt=paths=source_relative \
 ./person.proto

person.proto と生成されるファイルは別なリポジトリに置く

buf lint .
