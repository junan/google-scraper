.PHONY: build-dependencies test dev

build-dependencies:
	go get github.com/beego/bee/v2
	go get github.com/wellington/wellington/wt
	go mod tidy

build-assets:
	wt compile assets/stylesheets/application.scss -s compressed -b static/css

dev:
	docker-compose -f docker-compose.dev.yml up -d
	bee run

test:
	docker-compose -f docker-compose.test.yml up -d
	BEEGO_ENV=test go test -v -p 1 ./...
	docker-compose -f docker-compose.test.yml down
