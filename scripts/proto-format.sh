#!/bin/bash

set -e
set -x
# Find all .proto files in the protocol directory
proto_files=$(find ../protocol -name "*.proto")

# Loop through the files and format each one
for file in $proto_files; do
  clang-format -i -style=Google $file
done
