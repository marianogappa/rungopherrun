#!/bin/sh
rm *.so
go build -buildmode=plugin -ldflags "-pluginpath=plugin/hot-$(date +%s)" -o plugin-$(date +%s).so .
