.PHONY: test dev

build-dependencies:
	go get github.com/beego/bee

dev:
	bee run

test:
	go test -v -p 1 ./...
