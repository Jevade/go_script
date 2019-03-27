SHELL := /bin/bash
BASEDIR = $(shell pwd)

versionDir = "main"
#$(shell pwd)/pkg/version
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: gotool
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 58spy.linux -v  -ldflags ${ldflags} .
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o 58spy.exe -v  -ldflags ${ldflags} .
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o 58spy.mac  -v  -ldflags ${ldflags} .
clean:
	rm -rf 58spy.mac
	rm -rf 58spy.linux
	rm -rf 58spy.exe
	find . -name "[._]*.s[a-w][a-z]" | xargs rm -f {}
gotool:doc
	@echo "Formating and make code"
	gofmt -w .
	go tool vet .|	grep -v vendor;true
ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=liujiawei0524@gmail.com"
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ssl cert and key"
doc:
	@echo "Generate swagger api doc"
	swag init 
.PHONY: clean gotool ca doc help
