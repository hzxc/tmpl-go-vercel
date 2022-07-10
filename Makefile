GOPROXY:=$(shell go env GOPROXY)
AUTH:= "app/auth/proto"
PING:= "app/pingpong/proto"

.PHONY:hello
hello:
	echo hello\
	 world

.PHONY:go.proxy
go.proxy:
	echo $(GOPROXY)

.PHONY:auth.proto
auth.proto:
	protoc --go_out=$(AUTH) \
	--go-grpc_out=$(AUTH) \
	$(AUTH)/auth.proto

.PHONY:ping.proto
ping.proto:
	protoc --go_out=$(PING) \
	--go-grpc_out=$(PING) \
	$(PING)/pingpong.proto

.PHONY:proto
proto:
	protoc --go_out=proto \
	--go-grpc_out=proto \
	proto/conf.proto proto/job.proto proto/log.proto proto/order.proto proto/worker.proto

.PHONY:build
build: # linux下运行
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user main/main.go

.PHONY:install
	go install ./main/...

