PACKAGES = $$(go list ./... | grep -v /vendor/)

.PHONY: clean build install lint test bench cover cover-html coveralls readme workflow

clean:
	@go clean $(PACKAGES)
	@- rm -rf dist
	@- rm -rf coverage
	@- rm -rf gododir/godobin-*

build:
	@go build $(PACKAGES)

install:
	@go get $(PACKAGES)

lint:
	@go vet $(PACKAGES)
# @golint $(PACKAGES)

test:
	@go test $(PACKAGES) -v

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

workflow:
	@go install $(PACKAGES)
	@godo
