#!/usr/bin/env bash
. ./build_config.sh

docker run -it -p 8181:8181 -v "`pwd`":/root/app $IMAGE_NAME /bin/bash
#docker run -it --network=host -v `pwd`:/root/app $IMAGE_NAME /bin/bash
