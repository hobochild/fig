.PHONY: install 
install: go.sum
	go get -v -t

.PHONY: build 
build: install
	go build

.PHONY: test
test: build
	gotest -v ./utils

.PHONY: test_watch
test_watch:
	git ls-files | entr -r go test

.PHONY: dist
dist:
	env GOOS=linux GOARCH=amd64 go build -o ./dist/bump_linux_amd64
