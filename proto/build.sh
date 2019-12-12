#! /bin/bash

cd "$(dirname "$0")"

for proto in $(find ./github.com/smoug25/tensorflow_serving_helper/tensorflow -iname "*.proto"); do
echo processing: $proto step 2
protoc \
    -I . \
    --go_out=plugins=grpc:$GOPATH/src/ \
    $proto
done

#--ruby_out=plugins=grpc:../gen \
for proto in $(find ./github.com/smoug25/tensorflow_serving_helper/tensorflow_serving -iname "*.proto"); do
echo processing: $proto step 2
protoc \
    -I . \
    --go_out=plugins=grpc:$GOPATH/src/ \
    $proto
done
