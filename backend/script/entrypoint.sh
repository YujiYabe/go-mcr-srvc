#!/bin/sh
set -e

if [ "$DEBUG_MODE" = "true" ]; then
	echo "Starting Delve Debugger..."
	exec dlv debug ./cmd/main.go --headless --listen=:2345 --api-version=2 --log --build-flags="-mod=mod" --output=./debug/__debug_bin

else
	echo "Starting Air for Hot Reloading..."
	exec ./bin/air -c ./script/air.toml
fi
