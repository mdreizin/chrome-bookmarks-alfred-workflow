PACKAGES = $$(go list ./... | grep -v /vendor/)
ifdef TRAVIS_TAG
	VERSION=$(TRAVIS_TAG)
else
	VERSION=latest
endif
BIN_NAME:=chrome-bookmarks
GOBUILD_ARGS:=-ldflags "-X main.version=$(VERSION)"

.PHONY: clean build fmt deps lint test bench cover cover-html coveralls workflow

clean:
	@go clean $(PACKAGES)
	@- rm -rf build coverage gododir/godobin-* $$GOPATH/bin/$(BIN_NAME)

build:
	@GOOS=darwin GOARCH=amd64 go build $(GOBUILD_ARGS) -o $$GOPATH/bin/$(BIN_NAME) ./cli/...;

fmt:
	@go fmt $(PACKAGES)

deps:
	@go get -u -v github.com/kardianos/govendor
	@go get -u -v gopkg.in/godo.v2/cmd/godo
	@go get -u -v github.com/axw/gocov/gocov
	@go get -u -v github.com/matm/gocov-html
	@go get -u -v github.com/wadey/gocovmerge
	@go get -u -v github.com/mattn/goveralls
	@go get -u -v github.com/golang/lint/golint
	@glide install

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
	@- mkdir -p coverage
	@gocov test $(PACKAGES) | gocov-html > coverage/profile.html

coveralls:
	@- mkdir -p coverage
	@for pkg in $(PACKAGES); do \
		go test $$pkg -coverprofile="coverage/$$(basename $$pkg)-profile.cov"; \
	done
	@gocovmerge coverage/*-profile.cov > coverage/profile.cov
	@goveralls -coverprofile=coverage/profile.cov -service=travis-ci

workflow: build
	@godo -- --version=$(VERSION)
