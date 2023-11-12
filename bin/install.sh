#!/bin/bash

docker pull iimeta/iim-server:1.1.0

mkdir -p /data/iim-server/manifest/config

wget -P /data/iim-server/manifest/config https://github.com/iimeta/iim-server/raw/docker/manifest/config/config.yaml
wget https://github.com/iimeta/iim-server/raw/docker/bin/start.sh
