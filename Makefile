.PHONY: help generate test lint fmt dependencies clean check coverage service race .remove_empty_dirs .pre-check-go

SRCS = $(patsubst ./%,%,$(shell find . -name "*.go" -not -path "*vendor*" -not -path "*.pb.go"))
PBS = $(patsubst %.proto,%.pb.go,$(patsubst api%,pkg%,$(PROTOS)))

MOCK_PACKAGES = \
	internal/ingaccounts \
	internal/users \
	internal/savedmessages \
	internal/clients

MOCKED_FILES = $(shell find . -name DOES_NOT_EXIST_FILE $(patsubst %,-or -path "./%/mocks/*.go",$(MOCK_PACKAGES)))
MOCKED_FOLDERS = $(patsubst %,%/mocks,$(MOCK_PACKAGES))

#PROTO_MESSAGES = \
#	api/services/authentication/v1/authentication.proto


SQL_MODELS= \
	internal/users \
	internal/savedmessages \
	internal/ingaccounts

SQL_FOLDERS = $(patsubst %,%,$(SQL_MODELS))

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


generate: $(MOCKED_FILES) $(MOCKED_FOLDERS) | .remove_empty_dirs ## to generate all auto-generated files
.remove_empty_dirs:
	-find . -type d -print | xargs rmdir 2>/dev/null | true


generate_sql: .bin/sqlc
	$(foreach dir,$(SQL_FOLDERS), \
		.bin/sqlc generate -f $(dir)/sqlc.yaml; \
	)


get_list_files: $(SQL_FOLDERS) | .remove_empty_dirs ## to generate all auto-generated files
.remove_empty_dirs:
	-find . -type d -print | xargs rmdir 2>/dev/null | true

dependencies: | .pre-check-go .bin/golangci-lint ## to install the dependencies
	$(GO) mod download

clean: ## to remove generated files
	-rm -rf service
	-find . -type d -name mocks -exec rm -rf \{} +

service: $(SRCS)
	go build -o $@ -ldflags="$(LD_FLAGS)" ./cmd

test: | generate  ## to run tests
	go test ./...

lint: .bin/golangci-lint ## to lint the files
	.bin/golangci-lint run --config=.golangci-lint.yml --skip-files tables.go --skip-files dashboard.go --skip-files experiment_list --fix ./...

.bin/golangci-lint:
	if [ -z "$$(which golangci-lint)" ]; then curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b .bin/ $(LINTER_VERSION); else mkdir -p .bin; ln -s "$$(which golangci-lint)" $@; fi

.bin/sqlc:
	if [ -z "$(wildcard .bin/sqlc)" ]; then mkdir -p .bin; cd .bin; curl -sfL https://github.com/kyleconroy/sqlc/releases/download/v1.11.0/sqlc_1.11.0_linux_amd64.tar.gz --output sqlc.tar.gz; tar -xf sqlc.tar.gz; fi

collect_migrations:
	-mkdir -p internal/database/migrations
	$(foreach dir,$(SQL_FOLDERS), \
		cp -a $(dir)/migrations/. internal/database/migrations; \
	)

fmt: ## to run `go fmt` on all source code
	gofmt -s -w $(SRCS)

check: | generate ## Run tests
	go test ./...

race: | generate ## to run data race detector
	go test -timeout 30s -race ./...

coverage: coverage.cover coverage.html ## to run tests and generate test coverage data
	gocov convert $< | gocov report

coverage.html: coverage.cover
	go tool cover -html=$< -o $@

coverage.cover: $(SRCS) $(PBS) Makefile | generate
	-rm -rfv .coverage
	mkdir -p .coverage
	$(foreach pkg,$(PACKAGES),go test -timeout 30s -short -covermode=count -coverprofile=.coverage/$(subst /,-,$(pkg)).cover $(pkg)${\n})
	echo "mode: count" > $@
	grep -h -v "^mode:" .coverage/*.cover >> $@

.SECONDEXPANSION:
$(PBS): $$(patsubst %.pb.go,%.proto,$$(patsubst pkg%,api%,$$@)) | .pre-check-go
	$(PROTOC) $(PROTOC_OPTIONS) --go_out=plugins=grpc:$(GOPATH)/src ./$<

.SECONDEXPANSION:
$(MOCKED_FOLDERS): | .pre-check-go
	cd $(patsubst %/mocks,%,$@) && mockery --all --keeptree --outpkg mocks --output mocks

.SECONDEXPANSION:
$(MOCKED_FILES): $$(shell find $$(patsubst %/mocks,%,$$(patsubst %/mocks/,%,$$(dir $$@))) -maxdepth 1 -name "*.go") | $(MOCKED_FOLDERS)
	rm -rf $(dir $@)
	cd $(patsubst %/mocks,%,$(patsubst %/mocks/,%,$(dir $@))) && mockery --all --keeptree --outpkg mocks --output mocks

.pre-check-go:
	if [ -z "$$(which mockery)" ]; then go get -v github.com/vektra/mockery/v2/.../; fi

# Variables
ROOT := git.cafebazaar.ir/divar/divar-cloud-sand-boxing

PROTOC ?= protoc
PROTOC_OPTIONS ?= -I.
LINTER_VERSION = v1.39.0
SQLC_VERSION:=1.10.0
GoAdmin_VERSION = v1.39.0
GIT ?= git
COMMIT := $(shell $(GIT) rev-parse HEAD)
CI_COMMIT_TAG ?=
VERSION ?= $(strip $(if $(CI_COMMIT_TAG),$(CI_COMMIT_TAG),$(shell $(GIT) describe --tag 2> /dev/null || echo "$(COMMIT)")))
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
LD_FLAGS := -X $(ROOT)/pkg/info.Version=$(VERSION) -X $(ROOT)/pkg/info.Commit=$(COMMIT) -X $(ROOT)/pkg/info.BuildTime=$(BUILD_TIME)


