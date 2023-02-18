GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

UNAME := $(shell uname) ## Prone to error (in windows)
ARCHT := $(shell uname -m)
UNAMEP:= $(shell uname -smr)

CC=go
NM="venera"
all: help

## Configure Golang
install-go: ## Install or update golang env automatically, 
  ifeq ($(OS),Windows_NT) ## OS (Windows_NT

  else

    ifeq ($(UNAME),Linux ) ## OS (Linux)
			$(info $(UNAMEP))

      ifneq ($(shell command -v pacman),)   ## pacman
			pacman -S go
        
      else ifneq ($(shell command -v apt),) ## apt-get
			apt install golang-go

      else ifneq ($(shell command -v yum),) ## yum
			sudo yum -y install epel-release   ## !!!!!!!!!
			sudo yum -y install golang         ## not tested

      else
			$(error No appliable package managers found, you'll have to install it manually)

      endif

    else ## OS (Other)

    endif
  endif

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
