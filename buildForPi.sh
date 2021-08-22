#!/usr/bin/env bash

# Usage: ./buildForPi.sh GoScriptToBuild.go NameOfBinary

env GOOS=linux GOARCH=arm GOARM=5 go build -o $2 $1