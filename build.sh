#!/bin/bash

IMAGE="localhost/script-output-converter-build:latest"
docker build -t $IMAGE . && docker run -d --rm -v `pwd`:/build --name script-output-converter $IMAGE
