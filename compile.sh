#!/bin/bash

if [ "${1:0:1}" != '-' ]; then
    exec $1
fi

go build -a -ldflags="-w -s" -gcflags="-trimpath="`pwd` -asmflags=-trimpath=`pwd` -o script-output-converter main.go
