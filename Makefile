PKG = github.com/ilius/chunk32
TEST_ASSETS = $(GOPATH)/src/${PKG}/assets-test

default: build

build: chunk32

chunk32: *.go
	go build

assets_test.go: ${TEST_ASSETS}/*
	GOBIN=$(GOPATH)/bin go get github.com/a-urth/go-bindata/...
	$(GOPATH)/bin/go-bindata -prefix ${TEST_ASSETS} -pkg main -o assets_test.go ${TEST_ASSETS}/...

test: assets_test.go
	go test

