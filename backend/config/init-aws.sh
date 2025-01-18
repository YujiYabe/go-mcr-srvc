#!/bin/bash
# デフォルトリージョンを東京（ap-northeast-1）に設定
export AWS_DEFAULT_REGION=ap-northeast-1
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test

# aws --endpoint-url=http://localhost:4566 secretsmanager create-secret --name my-local-secret --secret-string ./localstack_secrets.json
aws --endpoint-url=http://localhost:4566 secretsmanager create-secret --name my-local-secret --secret-string file:///etc/localstack/init/ready.d/secrets.json

# aws --endpoint-url=http://localhost:4566 secretsmanager list-secrets
# aws --endpoint-url=http://localhost:4566 secretsmanager get-secret-value --secret-id my-local-secret
