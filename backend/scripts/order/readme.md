
install prerequisites
https://grpc.io/docs/languages/python/quickstart/

python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./delivery.proto


