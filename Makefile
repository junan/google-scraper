.PHONY: test dev

build-dependencies:
	go get github.com/beego/bee/v2

dev:
	bee run

test:
	go test -v -p 1 ./...
