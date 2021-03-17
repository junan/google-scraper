# google-scraper
A web application to google search with multiple keywords 

## Prerequisite
* [Go - 1.15](https://golang.org/doc/go1.15)
* [Docker](https://docs.docker.com/get-docker/)


## Usage

Clone the repository

`git clone git@github.com:junan/google-scraper.git`


#### Build development dependencies

  ```sh
  make build-dependencies
  ```

#### Build assets

  ```sh
  make build-assets
  ```

#### Run the worker

  ```sh
  make start-worker
  ```

#### Run the application in your local machine

  ```sh
  make dev
  ```

It will be accessible at: `http://localhost:8080`

### Run tests

```sh
make test
```
