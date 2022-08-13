#!/bin/bash

set -eu


until curl -s -f -o /dev/null "http://accountapi:8080/v1/organisation/accounts"; do
    >&2 echo "accountapi is unavailable - sleeping"
    sleep 3
done;

>&2 echo "accountapi is available"

cd /app
go test -v ./...

