#!/bin/sh

docker build -t gcloud . && docker run -p 2610:8080 gcloud