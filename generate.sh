#!/usr/bin/env bash

PADDED=$(printf %02d $1)

mkdir -p day$PADDED/input
cd day$PADDED
touch input/inp.txt
touch input/test.txt
cd ..
go run generate/day.go $PADDED