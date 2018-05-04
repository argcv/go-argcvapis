#!/usr/bin/env bash

if [ -z $GOPATH ]; then
  echo 'error: env $GOPATH not exists'
  exit 1
fi

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "BASE_DIR=$BASE_DIR"

# Prepare data
DELETE_ARGCV_REPO_ON_EXIT=0
ARGCV_REPO_URL=${ARGCV_REPO_URL:-'https://github.com/argcv/argcv.git'}
ARGCV_REPO_PATH=${ARGCV_REPO_PATH:-''}

PROTO_DIR="$ARGCV_REPO_PATH/argcv"

SRC_DIR="$GOPATH/src"

cmd_check() {
    hash $1 2>/dev/null
}

cmd_required() {
    if $(cmd_check $1); then
        echo "CHECK: $1 OK"
    else
        echo "CHECK: $1 FAILED"
        echo "Note: "$2
        exit 1
    fi
}

cmd_required protoc "You can install it using 'brew install protobuf' in macOS, or search protobuf for other platform"
cmd_required protoc-gen-go "You can install it using 'go get -u github.com/golang/protobuf/protoc-gen-go'"
cmd_required protoc-gen-grpc-gateway "You can install it using 'go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway'"
#cmd_required protoc-gen-swagger "You can install it using 'go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger'"
#cmd_required protoc-gen-grpc_python "You can install it using 'go get -u github.com/golang/protobuf/protoc-gen-go'"


if [[ -z $ARGCV_REPO_PATH ]]; then
  cmd_required git "command git is required, please refer to https://git-scm.com/book/en/v2/Getting-Started-Installing-Git and install it"
  # Not Exists
  ARGCV_REPO_PATH=$(mktemp -d)
  echo "git clone $ARGCV_REPO_URL $ARGCV_REPO_PATH"
  git clone $ARGCV_REPO_URL $ARGCV_REPO_PATH
  DELETE_ARGCV_REPO_ON_EXIT=1
fi

for file in $(find ${BASE_DIR} -type f -name "*.pb.go" -o -name "*.pb.gw.go"); do
    echo "Cleaning file: $file"
    rm -f ${file}
done

pushd $ARGCV_REPO_PATH
for file in $(find ${PROTO_DIR} -type f -name "*.proto"); do
    echo "Processing: $file"
    protoc \
        -I $ARGCV_REPO_PATH \
        --go_out=plugins=grpc:${SRC_DIR}\
        --grpc-gateway_out=logtostderr=true:${SRC_DIR} \
        ${file}
#        --swagger_out=logtostderr=true:${SRC_DIR} \
done
popd # $ARGCV_REPO_PATH


if [ -d $ARGCV_REPO_PATH ] && [[ $DELETE_ARGCV_REPO_ON_EXIT == "1" ]]; then
  echo "Clean Up: $ARGCV_REPO_PATH"
  rm -rf $ARGCV_REPO_PATH
else
  echo "Skip: Clean Up: $ARGCV_REPO_PATH"
fi


