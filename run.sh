#!/usr/bin/env bash

mkdir build
go build -o build/ ./day*
cd build
for f in *; do
    time ./$f
    echo
done
