# ================== 可配置参数 ==================
SCAN_DIR       ?= ./app,./infrastructure
OUTPUT_FILE    ?= auto_import.go
MODULE_NAME    ?=
CONCURRENCY    ?= 8
IGNORE_PATTERNS ?= **/vendor/**,**/.git/**,**/*_test.go,**/testdata/**,**/node_modules/**

# 构建相关
BINARY_NAME    ?= xianFuBlogGo
BUILD_DIR      ?= bin
LDFLAGS        ?= -s -w -extldflags "-static" -X "main.UserName="
GO_BUILD_CMD   = CGO_ENABLED=0 go build -v -a \
                 -ldflags '$(LDFLAGS)' \
                 -gcflags="all=-trimpath=$(CURDIR)" \
                 -asmflags="all=-trimpath=$(CURDIR)" \
                 -trimpath

# 自动检测当前系统环境
CURRENT_OS     := $(shell go env GOOS)
CURRENT_ARCH   := $(shell go env GOARCH)

# ================== 设置默认目标 ==================
.DEFAULT_GOAL := run

# ================== 目标定义 ==================
.PHONY: all generate run build build-all clean

# 完整流程（生成 + 运行），不设为默认
all: generate run

# 生成导入文件
generate:
	@echo "===== 运行扫描工具 ====="
	go run ./auto_tool/main.go \
		-module=$(MODULE_NAME) \
		-dir=$(SCAN_DIR) \
		-output=$(OUTPUT_FILE) \
		-concurrency=$(CONCURRENCY) \
		-ignore='$(IGNORE_PATTERNS)'

# 直接运行（默认目标）
run: generate
	@echo "===== 启动项目 ====="
	go run .

# 智能编译当前平台
build:
	@echo "===== 检测到当前系统: $(CURRENT_OS)/$(CURRENT_ARCH) ====="
	@echo "===== 编译当前平台可执行文件 ====="
	$(GO_BUILD_CMD) -o $(BINARY_NAME) .

# 交叉编译所有平台
build-all: build-linux build-mac-intel build-mac-arm

build-linux:
	@mkdir -p $(BUILD_DIR)
	@echo "===== 构建 Linux (amd64) ====="
	GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .

build-mac-intel:
	@mkdir -p $(BUILD_DIR)
	@echo "===== 构建 macOS Intel (amd64) ====="
	GOOS=darwin GOARCH=amd64 $(GO_BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .

build-mac-arm:
	@mkdir -p $(BUILD_DIR)
	@echo "===== 构建 macOS Apple Silicon (arm64) ====="
	GOOS=darwin GOARCH=arm64 $(GO_BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .

# 清理
clean:
	@echo "===== 清理构建产物 ====="
	rm -f $(BINARY_NAME)
	rm -rf $(BUILD_DIR)