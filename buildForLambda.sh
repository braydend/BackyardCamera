#!/usr/bin/env bash

# Usage: ./buildForLambda.sh GoScriptToBuild.go NameOfBinary

GOOS=linux GOARCH=amd64 go build -o $2 $1
zip "$2".zip $2
rm $2