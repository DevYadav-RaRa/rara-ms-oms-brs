#!/usr/bin/env bash

. ./build_config.sh
#docker build --progress=plain --no-cache ./_dev_docker -t $IMAGE_NAME
docker build ./_dev_docker -t $IMAGE_NAME

