# Devstore

## Package

* Web: [gin](https://github.com/gin-gonic/gin)
* Database access: [sqlx](github.com/jmoiron/sqlx)
* Data validation: [validator](github.com/go-playground/validator)
* Configuration: [viper](github.com/spf13/viper)
* Logging: [logrus](github.com/sirupsen/logrus)
* JWT: [golang-jwt](github.com/golang-jwt/jwt)
* Database: PostgreSQL

## Prerequisites

- Install [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).

## Running Locally

1. `make environment`
2. `make migration-up`
3. `cp app.env.sample app.env`, adjust the value
5. `make server`
6. App running!

## Other Commands :

You can run `make help` for showing all available commands :

```bash
❯ make help
environment                    Setup environment.
migrate-all                    Rollback migrations, all migrations
migrate-create                 Create a DB migration files e.g `make migrate-create name=migration-name`
migrate-down                   Rollback migrations, latest migration (1)
migrate-up                     Run migrations UP
server                         Running application
```
