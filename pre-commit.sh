#!/bin/sh -eux

goimports -w .
go tool vet -all .
if ! golint ./... \
        | grep -v "\-gen.go" \
        | grep -v ".pb.go" \
        ; then
    echo "golint ok!"
else
    exit 1
fi
