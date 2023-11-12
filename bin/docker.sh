#!/bin/bash
cd `dirname $0`
cd ../

docker build -f ./bin/Dockerfile -t iimeta/iim-server:1.1.0 .
