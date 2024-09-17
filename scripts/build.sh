#/bin/bash

go mod tidy && \
go fmt && \
go vet && \
go fix && \
staticcheck -go 1.23.1 ./... && \
go build -o bin/get_access_token
