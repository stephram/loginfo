#!/usr/bin/env bash

echo "Running tests..."
go test ./cmd/... --cover

echo "Building..."
go build -o ./main cmd/main.go
ls -al ./main