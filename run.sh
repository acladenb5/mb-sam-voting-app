#!/usr/bin/env bash

# GOOS=linux go build -o main
make clean && make
sam local start-api
