.PHONY: build-dependencies test dev

build-dependencies:
	go get github.com/beego/bee/v2
	go mod tidy

dev:
	docker-compose -f docker-compose.dev.yml up -d
	bee run

test:
	go test -v -p 1 ./...

