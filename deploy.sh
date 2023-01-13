#!/bin/bash
BUILD_ID=dontKillMe
export GOAPTH=/usr/local/gopath
export GOROOT=/usr/local/go
PATH=$GOROOT/bin:$PATH
source /etc/profile
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn
source /etc/profile

go build 
nohup ./tiktok &
