export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

AETHER_ROC_API_VERSION := latest
ONOS_BUILD_VERSION := v0.6.3

build: # @HELP build the Go binaries and run all validations (default)
build:
	CGO_ENABLED=1 go build -o build/_output/openapi-gen ./cmd/openapi-gen
	CGO_ENABLED=1 go build -o build/_output/aether-roc-api ./cmd/aether-roc-api

test: # @HELP run the unit tests and source code validation
test: build deps linters license_check
	CGO_ENABLED=1 go test -race github.com/onosproject/aether-roc-api/pkg/...
	CGO_ENABLED=1 go test -race github.com/onosproject/aether-roc-api/cmd/...
#	CGO_ENABLED=1 go test -race github.com/onosproject/aether-roc-api/api/...

coverage: # @HELP generate unit test coverage data
coverage: build deps
	./build/bin/coveralls-coverage

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 30m

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/ cmd/ tests/)"

aether-roc-api-base-docker: # @HELP build aether-roc-api base Docker image
	@go mod vendor
	docker build . -f build/base/Dockerfile \
		--build-arg ONOS_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		--build-arg ONOS_MAKE_TARGET=build \
		-t onosproject/aether-roc-api-base:${AETHER_ROC_API_VERSION}
	@rm -rf vendor

aether-roc-api-docker: aether-roc-api-base-docker # @HELP build aether-roc-api Docker image
	docker build . -f build/aether-roc-api/Dockerfile \
		--build-arg AETHER_ROC_API_BASE_VERSION=${AETHER_ROC_API_VERSION} \
		-t onosproject/aether-roc-api:${AETHER_ROC_API_VERSION}

images: # @HELP build all Docker images
images: build aether-roc-api-docker

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/aether-roc-api:${AETHER_ROC_API_VERSION}

all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} onosproject/aether-roc-api

bumponosdeps: # @HELP update "onosproject" go dependencies and push patch to git.
	./../build-tools/bump-onos-deps ${VERSION}

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/aether-roc-api/aether-roc-api ./cmd/onos/onos
	go clean -testcache github.com/onosproject/aether-roc-api/...

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
