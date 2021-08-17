#!/bin/bash

source .env
go build
CompileDaemon  --command="./web-cert-checker"