#!/usr/bin/env bash

# Exit Once failed
# set -Eeuxo pipefail

if [ ! -f "$PROTO_DIR/bin/protoc" ]; then
  # Cleanup Origin Data
  rm -rf $PROTO_BASE
  mkdir -p $PROTO_BASE
  wget -O - "https://github.com/google/protobuf/archive/v${PROTO_VERSION}.tar.gz" | tar xz -C /tmp
  cd "/tmp/protobuf-${PROTO_VERSION}"
  ./autogen.sh
  ./configure --prefix="$PROTO_DIR" --disable-shared
  make -j 4
  make install
fi



