#!/bin/bash

./tools/go-licenses check ./... \
  --ignore "github.com/stellatarum/docker-machine-driver-pve" \
  --disallowed_types=forbidden,restricted,reciprocal,unknown
