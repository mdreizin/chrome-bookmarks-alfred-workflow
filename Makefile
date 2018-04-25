PACKAGES = $$(go list ./... | grep -v /vendor/)
ifdef TRAVIS_TAG
VERSION=$(TRAVIS_TAG)
else
VERSION=latest
endif
BUILD_DIR:=build
COVER_DIR:=coverage
BIN_DIR:=$$GOPATH/bin
BIN_NAME:=chrome-bookmarks
GOBUILD_ARGS:=-ldflags "-X main.version=$(VERSION)"

.PHONY: clean build fmt deps lint test bench cover cover-html coveralls workflow

clean:
	@go clean $(PACKAGES)
	@- rm -rf ${COVER_DIR} ${BUILD_DIR} gododir/godobin-* ${BIN_DIR}/$(BIN_NAME)

build:
ifeq ($(TRAVIS),true)
	@gox $(GOBUILD_ARGS) -os="darwin" -arch="amd64" -osarch="!darwin/arm64" -output="${BIN_DIR}/${BIN_NAME}" ./cli/...
else
	@go build $(GOBUILD_ARGS) -o ${BIN_DIR}/$(BIN_NAME) ./cli/...
endif

fmt:
	@go fmt $(PACKAGES)

deps:
	@go get -u -v gopkg.in/godo.v2/cmd/godo
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v github.com/mitchellh/gox
	@dep ensure

lint:
	@go vet $(PACKAGES)
	@golint $(PACKAGES)

test:
	@go test -v $(PACKAGES)

bench:
	@go test $(PACKAGES) -bench . -benchtime 2s -benchmem

cover:
	@- rm -rf c.out
	@go test $(PACKAGES) -coverprofile=c.out

cover-html:
	@cover && go tool cover -html=c.out

workflow: build
	@godo -- --version=$(VERSION)
