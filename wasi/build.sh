#!/bin/bash

make build
buildah build --annotation "module.wasm.image/variant=compat" --platform wasi/wasm -t quay.io/jkeam/wasi
