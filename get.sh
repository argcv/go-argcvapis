#!/usr/bin/env bash

if [ -z $GOPATH ]; then
  echo "error: env: GOPATH not exists!!"
  exit 1
fi

ORG_DIR="$GOPATH/src/github.com/argcv"
PROJ_DIR="$ORG_DIR/go-argcvapis"


if [ -d $PROJ_DIR ]; then
  echo "error: project is exists!!"
  exit 2
fi

git clone git@github.com:argcv/go-argcvapis.git $PROJ_DIR

echo "cloned to $PROJ_DIR"

