# ====================================================================================
# Project-wide variables
# ====================================================================================
BINARY       := load-balancer
MAIN_CMD     := cmd/main.go
BUILD_DIR    := bin

# Default ports
LB_ADDR      := :8080

# Go tooling
GO           := go
GOTEST       := $(GO) test
GOFMT        := $(GO) fmt
GOVET        := $(GO) vet

# ====================================================================================
# Targets
# ====================================================================================
.PHONY: all build run backends fmt vet test clean help

all: build

## build        → compile the load-balancer binary into $(BUILD_DIR)/
build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY) $(MAIN_CMD)

## run          → build & run local load-balancer (on $(LB_ADDR))
run: build
	@echo "▶ Running load-balancer on $(LB_ADDR)"
	$(BUILD_DIR)/$(BINARY) --addr=$(LB_ADDR)


## fmt          → run gofmt over the entire module
fmt:
	@echo "▶ Formatting code..."
	$(GOFMT) ./...

## vet          → run govet over the entire module
vet:
	@echo "▶ Running go vet..."
	$(GOVET) ./...

## test         → run all unit tests
test:
	@echo "▶ Running tests..."
	$(GOTEST) ./...

## clean        → remove built binaries
clean:
	@echo "▶ Cleaning up..."
	rm -rf $(BUILD_DIR)

## help         → display this help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = "[: ]+## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'
