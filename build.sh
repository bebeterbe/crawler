#!/usr/bin/env bash
# go交叉编译，详细文档参考如下链接
# https://golang.org/doc/install/source#environment

#输出go变量
cwd=$(cd `dirname $0`; pwd)
export GOPATH=${GOPATH}:${cwd}

distPath=${cwd}/dist
version='0.0.1'

#目标路径
rm -rf ${distPath}
mkdir -p ${distPath}

#编译
#mac
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ${distPath}/agent_darwin_amd64_${version} ./src/app.go

#linux
GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o ${distPath}/agent_linux_386_${version} ./src/app.go
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${distPath}/agent_linux_amd64_${version} ./src/app.go
GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ${distPath}/agent_linux_arm_${version} ./src/app.go
GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ${distPath}/agent_linux_arm64_${version} ./src/app.go

#openbsd
GOOS=openbsd GOARCH=amd64 go build -ldflags "-s -w" -o ${distPath}/agent_openbsd_amd64_${version} ./src/app.go

#plan9
GOOS=plan9 GOARCH=amd64 go build -ldflags "-s -w" -o ${distPath}/agent_plan9_amd64_${version} ./src/app.go

#freebsd
GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w" -o ${distPath}/agent_freebsd_amd64_${version} ./src/app.go

#solaris
GOOS=solaris GOARCH=amd64 go build -ldflags "-s -w" -o ${distPath}/agent_solaris_amd64_${version} ./src/app.go

#window
GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o ${distPath}/agent_window_386_${version}.exe ./src/app.go
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ${distPath}/agent_window_amd64_${version}.exe ./src/app.go

