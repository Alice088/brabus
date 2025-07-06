#!/bin/bash

export PATH=$PATH:/home/gosha/go/bin

# shellcheck disable=SC2164
cd ../cmd/brabus
go build -o ../../prod/bin/brabus

# shellcheck disable=SC2164
cd ../banana
go build -o ../../prod/bin/banana
