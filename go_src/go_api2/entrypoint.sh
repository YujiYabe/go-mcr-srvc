#!/bin/sh
set -e
proxy_server_port=$1
web_server_port=$2

if [ ! -f "Gopkg.toml" ]; then
    dep init
fi

dep ensure

gin -p ${proxy_server_port} -a ${web_server_port}
