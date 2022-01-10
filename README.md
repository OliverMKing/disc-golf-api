# disc-golf-api

### Setup

Run `go get` to install dependencies. Clone the `.env_example` file `.env`.

Make sure you have [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) installed.

## Usage

### Local

Generate docs with `swag init`. Run the server with `go run main.go`.

Both commands can be run with a single `make run` command.

### Docker

Run `docker-compose up --build`.
