SHELL:=/bin/bash
ifdef TRAVIS_TAG
VERSION=$(TRAVIS_TAG)
else
VERSION=dev
endif
BUILD_DIR:=build
COVER_DIR:=c.out
CONF_DIR:=configs
ASSET_DIR:=assets
THIRD_DIR:=third_party
WORKFLOW_NAME:=workflow
ARCHIVE_NAME:=chrome-bookmarks.alfredworkflow
GOBUILD_ARGS:=-ldflags "-X main.version=$(VERSION)"

.PHONY: clean build fmt deps lint test bench cover cover-html

clean:
	@rm -rf ${COVER_DIR} ${BUILD_DIR}

build: clean
	@mkdir -p ${BUILD_DIR}
	@go run $(GOBUILD_ARGS) cmd/workflow-gen/main.go -workflow-tmpl-file="configs/info.plist.gohtml" -workflow-file="configs/workflow.yml" -browser-file="configs/browser.yml" -asset-dir="${ASSET_DIR}" -out-dir="${BUILD_DIR}"
ifeq ($(TRAVIS),true)
	@gox $(GOBUILD_ARGS) -os="darwin" -arch="amd64" -osarch="!darwin/arm64" -output="${BUILD_DIR}/${WORKFLOW_NAME}" ./cmd/workflow
else
	@go build $(GOBUILD_ARGS) -o ${BUILD_DIR}/$(WORKFLOW_NAME) ./cmd/workflow
endif
	@cp ${CONF_DIR}/browser.yml ${BUILD_DIR}
	@cp ${ASSET_DIR}/*.* ${BUILD_DIR}/
	@cp ${ASSET_DIR}/chrome.png ${BUILD_DIR}/icon.png
	@cp ${THIRD_DIR}/normalise ${BUILD_DIR}
	@pushd ${BUILD_DIR} &> /dev/null && \
		zip -rX ${ARCHIVE_NAME} ./* -x ${ARCHIVE_NAME} &> /dev/null && \
	popd &> /dev/null

format:
	@go fmt ./...

deps:
	@grep -o '".*"' tools.go | tr -d '"' | tr -s '\r\n' ' ' | go install `xargs -0`

lint:
	@go vet ./...
	@golint ./...

test:
	@go test -v ./...

bench:
	@go test ./... -bench . -benchtime 2s -benchmem

cover:
	@- rm -rf c.out
	@go test ./... -coverprofile=c.out

cover-html: cover
	@go tool cover -html=c.out
