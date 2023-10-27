#!/bin/bash

docker run -d \
  --network host \
  --restart=always \
  -p 9000:9000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/iim-server/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  --name iim-server \
  iimeta/iim-server:1.0.0
