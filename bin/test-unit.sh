#!/bin/bash

set -ex

bin=$(dirname $0)

reformatted_packages="$($bin/go fmt $($bin/go list github.com/DennisDenuto/property-price-collector/... 2>/dev/null | grep -v /vendor/ ))"

if [[ $reformatted_packages = *[![:space:]]* ]]; then
  echo "go fmt reformatted the following packages:"
  echo $reformatted_packages
  exit 1
fi

echo -e "\n Running unit tests..."
$bin/env ginkgo -r $race -trace -skipPackage="main,vendor,integration" $@