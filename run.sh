#!/bin/sh

docker build -t gcloud . && docker run -e DIREKTIV_TEST=true -p 9191:8080 -p 9292:9292 gcloud