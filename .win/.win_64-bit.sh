#!/bin/bash
# Compile for Windows, 64-bit version
GOOS=windows GOARCH=amd64 go build -o bin/wnoise_64.exe src/wnoise.go
