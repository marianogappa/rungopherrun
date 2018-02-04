#!/bin/bash
docker run -v /Users/marianol/Code/go:/mnt/go -e GOPATH=/mnt/GOPATH -ti frolvlad/alpine-go sh
