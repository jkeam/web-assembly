#!/bin/bash

make build
podman build -t quay.io/jkeam/wasm -f ./Containerfile .
