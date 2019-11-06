#!/bin/bash

# Setting up the GO PATH and bin directory to PATH

echo " Setting up GOPATH and GOBIN variable to your current Workspace"

CURR_DIR="$(pwd)"

export GOPATH=${CURR_DIR}
export GOBIN=${GOPATH}/bin
export PATH=$PATH:${GOPATH}/bin

echo "GO PATH --> $(go env GOPATH)"
echo "GO BIN --> $(go env GOBIN)"

echo "Please Verify the Output of GOPATH and GOBIN"
echo " Run this file using . ./setup.sh to set the Go Path Correctly"