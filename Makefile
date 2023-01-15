GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

CC=go
NM="venera"

all: help


## Run
run: ## Build project in volatile mode
	$(CC) run $(NM)


## Build
build: ## Build project and organize files
	[ -d bin/ ] || `echo "Creating bin/" && mkdir bin/`
	$(CC) build -o bin/$(NM) $(NM)

build-op: ## Build with optimization flags
	[ -d bin/ ] || `echo "creating bin/" && mkdir bin/`
	$(CC) build -o bin/$(NM) -ldflags="-s -w" $(NM)

build-run: ## Build and run project
	[ -d bin/ ] || `echo "creating bin/" && mkdir bin/`
	$(CC) build -o bin/$(NM) $(NM)
	./bin/$(NM)

## Test
test: ## Run tests of the project
	@echo "Test"

## Clear
clear: ## Clear compilation garbage
	@echo "Cleaning bin/"
	@rm bin/*

## Help
help: ## Show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
