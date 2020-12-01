#!/bin/bash

cd app
# point to minikube dockerd
eval $(minikube -p minikube docker-env)
docker build . -t cf/ip-app:v0.1.0
# reset to host dockerd
eval $(minikube docker-env -u)
