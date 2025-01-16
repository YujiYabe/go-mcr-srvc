export AWS_DEFAULT_REGION=ap-northeast-1&& export AWS_ACCESS_KEY_ID=test && export AWS_SECRET_ACCESS_KEY=test
aws --endpoint-url=http://localhost:4566 secretsmanager create-secret --name my-local-secret --secret-string file://backend/config/localstack_secrets.json

export AWS_DEFAULT_REGION=ap-northeast-1 && export AWS_ACCESS_KEY_ID=test && export AWS_SECRET_ACCESS_KEY=test
aws --endpoint-url=http://localhost:4566 secretsmanager list-secrets

aws --endpoint-url=http://localhost:4566 secretsmanager get-secret-value --secret-id my-local-secret

```
batch-get-secret-value                   | cancel-rotate-secret
create-secret                            | delete-resource-policy
delete-secret                            | describe-secret
get-random-password                      | get-resource-policy
get-secret-value                         | list-secret-version-ids
list-secrets                             | put-resource-policy
put-secret-value                         | remove-regions-from-replication
replicate-secret-to-regions              | restore-secret
rotate-secret                            | stop-replication-to-replica
tag-resource                             | untag-resource
update-secret                            | update-secret-version-stage
validate-resource-policy                 | help
```
