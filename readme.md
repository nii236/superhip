# Superhip

A webapp boilerplate for the lazy. And hip.

# Spinup

## Dependencies

* Install docker
* Install Go
* Install NVM
* Install yarn

```
docker pull postgres:10.3-alpine
```

## Run

## Run DB

```
docker run -d -p 5432:5432 --name devdb -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=dev -e POSTGRES_DB=dev postgres:10.3-alpine
```

## Create DB

```
pgcli -U dev -W -h localhost
CREATE DATABASE superhip;
```

## Migrate data

```
cd server
psq;
migrate -database "postgres://dev:dev@localhost:5432/superhip?sslmode=disable" -path ./migrations up
```

## Run server

```
cd server
go run *.go
```

## Run client

```
cd client
yarn start
```
