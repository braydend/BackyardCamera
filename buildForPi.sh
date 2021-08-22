#!/usr/bin/env bash

# Usage: ./buildForPi.sh GoScriptToBuild.go NameOfBinary

env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o $2 $1