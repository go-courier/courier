PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git describe --always)-devel
GOTEST = go test -v -race

test:
	$(GOTEST) ./...

cover:
	$(GOTEST) -coverprofile=coverage.txt -covermode=atomic ./...

release:
	git push
	git push origin $(VERSION)
