#!/bin/bash
set -xe
go build -o ./main main.go
npx @modelcontextprotocol/inspector -- ./main
