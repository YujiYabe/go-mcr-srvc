#!/bin/sh
set -e
proxy_server_port=$1
web_server_port=$2

if [ ! -f "Gopkg.toml" ]; then
    dep init
fi

dep ensure

# bash

gin -p 7171 -a 7070
# gin -p ${proxy_server_port} -a ${web_server_port}
