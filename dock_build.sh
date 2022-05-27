#!/usr/bin/env bash

. ./build_config.sh
docker build . -t $IMAGE_NAME

