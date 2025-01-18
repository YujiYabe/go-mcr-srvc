#!/usr/bin/env bash
set -e

# Host and port hardcoded for LocalStack
host_name="localstack"
# host_name="localhost"
port="4566"
url="http://${host_name}:${port}"

# Command to execute after LocalStack is ready
cmd="$@"

# Wait for LocalStack to be available
until curl -s -o /dev/null -w "%{http_code}" "$url" | grep "200" >/dev/null; do
	echo >&2 "Waiting for LocalStack at $url to be available..."
	sleep 2
done

echo >&2 "LocalStack is available - executing command"

if [ -n "$cmd" ]; then
	exec $cmd
fi

# ---------------------------

if [ "$DEBUG_MODE" = "true" ]; then
	echo "Starting Delve Debugger..."
	exec dlv debug ./cmd/main.go --headless --listen=:2345 --api-version=2 --log --build-flags="-mod=mod" --output=./debug/__debug_bin

else
	echo "Starting Air for Hot Reloading..."
	exec ./bin/air -c ./script/air.toml
fi
