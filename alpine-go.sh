#!/bin/bash
docker run -v /Users/marianol/Code/go:/mnt/go -e GOPATH=/mnt/go -ti frolvlad/alpine-go sh
