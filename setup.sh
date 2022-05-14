#!/bin/sh

docker build -t server ./server
docker build -t worker ./worker

docker run --network="host" server &
docker run worker &
