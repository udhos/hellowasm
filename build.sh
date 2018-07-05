#!/bin/bash

gofmt -s -w ./*/*.go
go tool fix ./*/*.go
go tool vet ./console
go tool vet ./hellowasm_server

go test ./hellowasm_server
go install ./hellowasm_server

# run/console/wasm_exec.html requires test.wasm
GOARCH=wasm GOOS=js go build -o ./run/console/test.wasm ./console
