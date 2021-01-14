.PHONY: build-dependencies test dev

build-dependencies:
	go get github.com/beego/bee/v2
	psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'google_scraper_development'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE google_scraper_development"

dev:
	bee run
	docker-compose -f docker-compose.dev.yml up -d

test:
	go test -v -p 1 ./...

