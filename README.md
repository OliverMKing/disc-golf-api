# disc-golf-api

### Setup

Run `go get` to install dependencies. Clone the `.env_example` file `.env`.

Make sure you have [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) installed.

## Usage

### Docker

Run `docker-compose up --build`.

After starting Postgres, use the following commands to create the database tables.

```bash
set -o allexport; source .env; set +o allexport
export POSTGRESQL_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable"
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
