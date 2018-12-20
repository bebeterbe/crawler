#!/usr/bin/env bash
cwd=$(cd `dirname $0`; pwd)
export GOPATH=${GOPATH}:${cwd}
