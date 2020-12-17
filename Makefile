.PHONY: test dev production

build-dependencies:
	go get github.com/beego/bee/v2

dev:
	bee run

test:
	go test -v -p 1 ./...

production:
	BEEGO_ENV=${BEEGO_ENV} bee run
