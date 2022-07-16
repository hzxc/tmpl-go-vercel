GOPROXY:=$(shell go env GOPROXY)

.PHONY:hello
hello:
	echo hello\
	 world

.PHONY:go.proxy
go.proxy:
	echo $(GOPROXY)

.PHONY:go.build
go.build: # linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user main/main.go

.PHONY:go.install
go.install:
	go install ./main/...

.PHONY:mysql.docker
mysql.docker:
	docker container run --name mysql57 -p 3306:3306 \
	-v G:/mysql/data:/var/lib/mysql \
	-v G:/mysql/mysqld:/var/run/mysqld \
	-v G:/mysql/my.cnf:/etc/mysql/my.cnf \
	-e MYSQL_ROOT_PASSWORD=password \
	-d mysql:5.7

