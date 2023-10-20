#!/bin/sh

# masterDB ---------------------------------
. ../.env
masterDB="master.sqlite3"

sqlite3 ${masterDB} < drop.sql
sqlite3 ${masterDB} < create.sql
sqlite3 ${STORE_ID}.sqlite3 < create.sql

cat data.tsv | sqlite3 ${masterDB} ".mode tabs" ".import /dev/stdin products"

