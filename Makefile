PACKAGES = $(shell go list ./... | grep -v '/integration_tests')

ifndef DCL_VERSION
DCL_VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
endif

ifndef DCL_COMMIT
DCL_COMMIT := $(shell git log -1 --format='%H')
endif

UID := $(shell id -u)
#GID := $(shell id -g)

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=DcLedger \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=dcld \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(DCL_VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(DCL_COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'
OUTPUT_DIR ?= build

LOCALNET_DIR ?= .localnet
LOCALNET_DOCKER_NETWORK = "distributed-compliance-ledger_localnet"

remove_containers = $(if $(1),docker rm -f $(1),true)
localnet_containers = $(shell docker ps --format '{{.ID}}' --filter network=$(LOCALNET_DOCKER_NETWORK))
clean_network = $(call remove_containers,$(call localnet_containers))

LICENSE_TYPE = "apache"
COPYRIGHT_YEAR = "2020"
COPYRIGHT_HOLDER = "DSR Corporation"
LICENSED_FILES = $(shell find . -type f -not -path '*/.*' -not -name '*.md' -not -name 'requirements.txt')

all: install

build: go.sum
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT_DIR)/dcld ./cmd/dcld

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/dcld

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	go test -v $(PACKAGES)

lint:
	golangci-lint run ./... --timeout 5m0s

license:
	addlicense -l ${LICENSE_TYPE} -y ${COPYRIGHT_YEAR} -c ${COPYRIGHT_HOLDER} ${LICENSED_FILES}

license-check:
	addlicense -l ${LICENSE_TYPE} -y ${COPYRIGHT_YEAR} -c ${COPYRIGHT_HOLDER} -check ${LICENSED_FILES}

clean:
	rm -rf $(OUTPUT_DIR)

# Docker

image:
	docker build -t dcledger --build-arg TEST_UID=${UID} \
		--build-arg DCL_VERSION=${DCL_VERSION} --build-arg DCL_COMMIT=${DCL_COMMIT} .

localnet_init:
	/bin/bash ./genlocalnetconfig.sh

localnet_start:
	@if [ -d "${LOCALNET_DIR}/observer0" ]; then\
		docker-compose --profile observers up -d;\
	else\
		docker-compose up -d;\
	fi

localnet_stop:
	docker-compose down || ($(call clean_network) && docker-compose down)

localnet_export: localnet_stop
	docker-compose run node0 dcld export --for-zero-height  >genesis.export.node0.json
	docker-compose run node1 dcld export --for-zero-height  >genesis.export.node1.json
	docker-compose run node2 dcld export --for-zero-height  >genesis.export.node2.json
	docker-compose run node3 dcld export --for-zero-height  >genesis.export.node3.json
	@if [ -d "${LOCALNET_DIR}/observer0" ]; then\
		docker-compose run observer0 dcld export --for-zero-height  >genesis.export.observer0.json;\
	fi


localnet_reset: localnet_stop
	docker-compose run node0 dcld unsafe-reset-all
	docker-compose run node1 dcld unsafe-reset-all
	docker-compose run node2 dcld unsafe-reset-all
	docker-compose run node3 dcld unsafe-reset-all
	@if [ -d "${LOCALNET_DIR}/observer0" ]; then\
		docker-compose run observer0 dcld unsafe-reset-all;\
	fi

localnet_clean: localnet_stop
	rm -rf $(LOCALNET_DIR)
	rm -rf $(HOME)/.dcl


localnet_rebuild: localnet_clean localnet_init


.PHONY: all build install test lint clean image localnet_init localnet_start localnet_stop localnet_clean localnet_export localnet_reset license license-check localnet_rebuild
