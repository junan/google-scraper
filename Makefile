# Include variables from ENV file
ENV =
-include .env
ifdef ENV
-include .env.$(ENV)
endif
export

.PHONY: build-dependencies test dev start-worker db-migrate db-rollback

build-dependencies:
	go get github.com/beego/bee/v2
	go get github.com/wellington/wellington/wt
	go mod tidy

build-assets:
	wt watch assets/stylesheets/application.scss -s compressed -b static/css

start-worker:
	go run worker/main.go

db-migrate:
	bee migrate -driver=postgres -conn="$(DATABASE_URL)"

db-rollback:
	bee migrate rollback -driver=postgres -conn="$(DATABASE_URL)"

dev:
	docker-compose -f docker-compose.dev.yml up -d
	make db-migrate
	bee run

test:
	docker-compose -f docker-compose.test.yml up -d
	BEEGO_ENV=test go test -v -p 1 ./...
	docker-compose -f docker-compose.test.yml down
