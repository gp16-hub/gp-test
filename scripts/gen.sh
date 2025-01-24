#!/bin/bash

set -e
# Find all .proto files in the protocol directory
proto_files=$(find ../protocol -name "*.proto")
goctl env check --install --verbose --force
# Loop through the files and format each one
rm -rf ../proto/gen
mkdir -p ../proto/gen
for file in $proto_files; do
  dir=$(dirname "$file") # Get the directory of the file
  # Generate code with protoc
  echo "Generating code for $file..."
  protoc --go_out=../proto/gen --go-grpc_out=../proto/gen -I="$dir" "$file"
done