#!/bin/bash

set -xeu

function start() {
    minikube start
    set +e
        pachctl deploy local
    set -e
}

function stop() {
    minikube stop
}

start

echo 'done'