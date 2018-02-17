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
	@go get -u -v github.com/kardianos/govendor
	@go get -u -v gopkg.in/godo.v2/cmd/godo
	@go get -u -v github.com/axw/gocov/gocov
	@go get -u -v github.com/matm/gocov-html
	@go get -u -v github.com/mattn/goveralls
	@go get -u -v github.com/golang/lint/golint
	go get -u -v github.com/mitchellh/gox
	@dep ensure

lint:
	@for pkg in $(PACKAGES); do \
		go tool vet $$(basename $$pkg); \
		golint $$(basename $$pkg); \
	done

test:
	@go test -v $(PACKAGES)

bench:
	@go test $(PACKAGES) -bench . -benchtime 2s -benchmem

cover:
	@gocov test $(PACKAGES) | gocov report

cover-html:
	@- mkdir -p ${COVER_DIR}
	@gocov test $(PACKAGES) | gocov-html > ${COVER_DIR}/profile.html

coveralls:
	@- mkdir -p ${COVER_DIR}
	@go test $(PACKAGES) -coverprofile="${COVER_DIR}/profile.cov"
	@goveralls -coverprofile=${COVER_DIR}/profile.cov -service=travis-ci

workflow: build
	@godo -- --version=$(VERSION)
