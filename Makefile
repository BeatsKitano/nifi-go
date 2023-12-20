# 定义默认的架构为 amd
ARCH ?= amd64

# 定义一个变量，用于存储当前时间
CURRENT_TIME := $(shell date +"%Y-%m-%d %H:%M:%S")

.PHONY: build
# build
build:
	@echo "start build $(ARCH)..."; 
	@mkdir -p bin;
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -ldflags "main.Build=$(CURRENT_TIME)" -o ./bin ./...
	@echo "build finish";

.PHONY: run
# run
run:
	@echo "start run...";
	go run ./...