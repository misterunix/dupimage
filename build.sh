#!/bin/sh

env GOOS=linux GOARCH=amd64 go build -o bin/dupimage

cp bin/dupimage ~/bin/dupimage

