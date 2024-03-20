#!/bin/bash

source src/bash/vars/vars_go.sh

echo "Changing working dir $PWD"
echo

cd src/go/cli
echo "$PWD"

echo "Building go program to $go_output"
echo


go build -o  $go_file_path

echo "$go_file_name built"
echo