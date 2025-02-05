#!/bin/bash

make build
buildah build --annotation "module.wasm.image/variant=compat" -t quay.io/jkeam/wasi
