# google-scraper
A web application to google search with multiple keywords 

## Demo

- [Staging](https://nimble-google-scraper-staging.herokuapp.com/)
- [Production](https://nimble-google-scraper.herokuapp.com/)
- [Postman collection](https://documenter.getpostman.com/view/11835486/TzJsfy75})

## Prerequisite
* [Go - 1.15](https://golang.org/doc/go1.15)
* [Docker](https://docs.docker.com/get-docker/)


## Usage

Clone the repository

`git clone git@github.com:junan/google-scraper.git`

### Create the `.env` file

Create a `.env` file and copy the contents of `.env.example` file into the `.env` file

#### Build development dependencies

  ```sh
  make build-dependencies
  ```

#### Build assets

  ```sh
  make build-assets
  ```

#### Run the migration, and the application in your local machine

  ```sh
  make dev
  ```

#### Run the worker

  ```sh
  make start-worker
  ```

It will be accessible at: `http://localhost:8080`

### Run tests

```sh
make test
```
