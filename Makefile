# Environment variables set when running the Go compiler.
GOOS ?= $(shell ${GO} env GOOS)
GOARCH ?= $(shell ${GO} env GOARCH)
GO_BUILD_ENVVARS = \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	CGO_ENABLED=0 \

# The go commands and the minimum Go version that must be used to build the app.
GO ?= go
GOFMT ?= $(shell ${GO} env GOROOT)/bin/gofmt
GO_VERSION_KIALI = 1.17

# Identifies the current build.
# These will be embedded in the app and displayed when it starts.
VERSION ?= v0.0.1
COMMIT_HASH ?= $(shell git rev-parse HEAD)

# A default go GOPATH if it isn't user defined
GOPATH ?= ${HOME}/go


build:
	@echo Building...
	${GO_BUILD_ENVVARS} ${GO} build \
		-o ${GOPATH}/bin/kiali -ldflags "-X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH}"

## clean: Clean ${GOPATH}/bin/kiali, ${GOPATH}/pkg/*, ${OUTDIR}/docker and the kiali binary
clean:
	@echo Cleaning...
	@rm -f kiali
	@rm -rf ${GOPATH}/bin/kiali
	@[ -d ${GOPATH}/pkg ] && chmod -R +rw ${GOPATH}/pkg/* || true
	@rm -rf ${GOPATH}/pkg/*
	@rm -rf ${OUTDIR}/docker


clean-all: clean
	@rm -rf ${OUTDIR}

## install: Install missing dependencies. Runs `go install` internally
install:
	@echo Installing...
	${GO_BUILD_ENVVARS} ${GO} install \
		-ldflags "-X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH}"


#
# Lint targets
#

## lint-install: Installs golangci-lint
lint-install:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(${GO} env ${GOPATH})bin v1.23.8

## lint: Runs golangci-lint
# doc.go is ommited for linting, because it generates lots of warnings.
lint:
	golangci-lint run --skip-files "doc\.go" --tests --timeout 5m		