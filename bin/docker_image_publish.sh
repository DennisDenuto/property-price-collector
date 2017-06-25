#!/usr/bin/env bash
set -xeu

pushd k8s
    docker build --no-cache -t property-price-collector .
    docker tag property-price-collector dennisdenuto/property-price-collector
    docker push dennisdenuto/property-price-collector
popd