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
	@go install $(PACKAGES)

fmt:
	@go fmt $(PACKAGES)

restore:
	@go get -u -v github.com/kardianos/govendor
	@go get -u -v gopkg.in/godo.v2/cmd/godo
	@go get -u -v github.com/axw/gocov/gocov
	@go get -u -v github.com/matm/gocov-html
	@go get -u -v github.com/wadey/gocovmerge
	@go get -u -v github.com/mattn/goveralls
	@go get -u -v github.com/golang/lint/golint
	@govendor sync

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

workflow: install
	@godo

readme:
	@npm run gitdown
