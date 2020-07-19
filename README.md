# Simple HTTP server on golang.

## About

Simple server build on `net/http`. Using `github.com/jinzhu/gorm` as ORM.

## Installation

### Requirements

Golang >= 1.14
Postgres database.

### Build

Clone current repository, get in it and run:

```bash
go build .
```

Be sure to create `hw_db` database before run.

### Docker

You can use docker-compose commands to run service:

```bash
docker-compose up
```

Service listens to port 5000 by default.

## Tests

Tests are in `game_handler_test.go` and `repo_test.go` files.
To run tests:

```bash
go test
```

There is `request.sh` file containing curl request to the service.
