#!/usr/bin/env bash
. ./build_config.sh

#docker run -it -v `pwd`:/root/app /bin/bash
sudo docker run --restart=always -p 8181:8181 -it -v "`pwd`":/root/app -v "`pwd`/logs":/logs -d  $IMAGE_NAME
