#!/bin/bash
# Compile for Windows, 32-bit version
GOOS=windows GOARCH=386 build -o bin/wnoise_32.exe src/wnoise.go
