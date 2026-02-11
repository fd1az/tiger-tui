.PHONY: help build run clean test fmt vet tidy lint install-tools dev tb-up tb-down tb-reset tb-logs

# Load .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Project variables
BINARY_NAME=tiger-tui
BUILD_DIR=bin
MAIN_PATH=./cmd/tiger-tui
GO=go
GOFLAGS=-v

# Colors for terminal output
BLUE := \033[0;34m
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m
BOLD := \033[1m

help: ## Show this help message
	@echo "$(BOLD)$(BLUE)tiger-tui Makefile Commands$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'
	@echo ""

# ============================================================================
# Build & Run
# ============================================================================

build: ## Build the tiger-tui binary
	@echo "$(BLUE)Building $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)Build complete: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

run: build ## Build and run tiger-tui
	@./$(BUILD_DIR)/$(BINARY_NAME)

dev: ## Run in development mode with hot reload (requires air)
	@if ! command -v air > /dev/null; then \
		echo "$(YELLOW)Installing air for hot reload...$(NC)"; \
		go install github.com/cosmtrek/air@latest; \
	fi
	@air

clean: ## Remove build artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "$(GREEN)Clean complete$(NC)"

# ============================================================================
# Testing
# ============================================================================

test: ## Run all tests
	@echo "$(BLUE)Running tests...$(NC)"
	@$(GO) test -v -race ./...
	@echo "$(GREEN)Tests complete$(NC)"

test-coverage: ## Run tests with coverage report
	@echo "$(BLUE)Running tests with coverage...$(NC)"
	@$(GO) test -v -race -coverprofile=coverage.out ./...
	@$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

test-short: ## Run short tests only
	@echo "$(BLUE)Running short tests...$(NC)"
	@$(GO) test -v -short ./...

bench: ## Run benchmarks
	@echo "$(BLUE)Running benchmarks...$(NC)"
	@$(GO) test -bench=. -benchmem ./...

# ============================================================================
# Code Quality
# ============================================================================

fmt: ## Format Go code
	@echo "$(BLUE)Formatting code...$(NC)"
	@$(GO) fmt ./...
	@echo "$(GREEN)Format complete$(NC)"

vet: ## Run go vet
	@echo "$(BLUE)Running go vet...$(NC)"
	@$(GO) vet ./...
	@echo "$(GREEN)Vet complete$(NC)"

lint: ## Run golangci-lint
	@if ! command -v golangci-lint > /dev/null; then \
		echo "$(YELLOW)Installing golangci-lint...$(NC)"; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@echo "$(BLUE)Running linter...$(NC)"
	@golangci-lint run ./...
	@echo "$(GREEN)Lint complete$(NC)"

check: fmt vet lint ## Run all code quality checks
	@echo "$(GREEN)All checks passed$(NC)"

# ============================================================================
# Dependencies
# ============================================================================

deps: ## Download dependencies
	@echo "$(BLUE)Downloading dependencies...$(NC)"
	@$(GO) mod download
	@echo "$(GREEN)Dependencies downloaded$(NC)"

deps-update: ## Update dependencies
	@echo "$(BLUE)Updating dependencies...$(NC)"
	@$(GO) get -u ./...
	@$(GO) mod tidy
	@echo "$(GREEN)Dependencies updated$(NC)"

tidy: ## Tidy Go modules
	@echo "$(BLUE)Tidying Go modules...$(NC)"
	@$(GO) mod tidy
	@echo "$(GREEN)Tidy complete$(NC)"

# ============================================================================
# Development Tools
# ============================================================================

install-tools: ## Install all development tools
	@echo "$(BLUE)Installing development tools...$(NC)"
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "$(GREEN)Tools installed$(NC)"

setup: install-tools deps ## Initial project setup
	@echo "$(GREEN)Project setup complete$(NC)"

# ============================================================================
# TigerBeetle (Docker)
# ============================================================================

tb-up: ## Start TigerBeetle via Docker Compose
	@echo "$(BLUE)Starting TigerBeetle...$(NC)"
	@docker compose up -d tigerbeetle
	@echo "$(GREEN)TigerBeetle running on port 3000$(NC)"

tb-down: ## Stop TigerBeetle
	@echo "$(YELLOW)Stopping TigerBeetle...$(NC)"
	@docker compose down
	@echo "$(GREEN)TigerBeetle stopped$(NC)"

tb-reset: ## Stop TigerBeetle and delete data
	@echo "$(RED)Resetting TigerBeetle (deleting all data)...$(NC)"
	@docker compose down -v
	@echo "$(GREEN)TigerBeetle reset complete$(NC)"

tb-logs: ## Show TigerBeetle logs
	@docker compose logs -f tigerbeetle

# ============================================================================
# Utility Commands
# ============================================================================

info: ## Show project information
	@echo "$(BOLD)$(BLUE)tiger-tui Project Information$(NC)"
	@echo ""
	@echo "  $(YELLOW)Binary:$(NC)      $(BINARY_NAME)"
	@echo "  $(YELLOW)Main Path:$(NC)   $(MAIN_PATH)"
	@echo "  $(YELLOW)Build Dir:$(NC)   $(BUILD_DIR)"
	@echo "  $(YELLOW)Go Version:$(NC)  $$($(GO) version | cut -d' ' -f3)"
	@echo ""

version: ## Show Go version
	@$(GO) version

all: clean check build test ## Clean, check, build, and test

# Default target
.DEFAULT_GOAL := help
