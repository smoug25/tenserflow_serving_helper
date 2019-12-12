#! /bin/bash

cd "$(dirname "$0")"

for proto in $(find ./tensorflow -iname "*.proto"); do
echo processing: $proto step 2
protoc \
    -I . \
    --go_out=plugins=grpc:../ \
    $proto
done

#--ruby_out=plugins=grpc:../gen \
for proto in $(find ./tensorflow_serving -iname "*.proto"); do
echo processing: $proto step 2
protoc \
    -I . \
    --go_out=plugins=grpc:../ \
    $proto
done
