#/bin/bash

go mod tidy
go vet
go fix
go build -o bin/get_access_token
