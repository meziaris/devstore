# Devstore

## Package

* Web: [gin](https://github.com/gin-gonic/gin)
* Database access: [sqlx](https://github.com/jmoiron/sqlx)
* Data validation: [validator](https://github.com/go-playground/validator)
* Configuration: [viper](https://github.com/spf13/viper)
* Logging: [logrus](https://github.com/sirupsen/logrus)
* JWT: [golang-jwt](https://github.com/golang-jwt/jwt)
* Database: PostgreSQL
* Image Storage: [cloudinary-go](https://github.com/cloudinary/cloudinary-go/v2)

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
