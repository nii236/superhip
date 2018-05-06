# Superhip

A webapp boilerplate for the lazy. And hip.

Notable tech:

* Postgrest wrapping Postgresql
* Go server for fancier non crud queries (e.g. sign up, sign in, forgot password)
* react-admin for management of data

# Spinup

## Dependencies

* Install docker
* Install Postgrest
* Install Go
* Install NVM
* Install yarn

```
docker pull postgres:10.3-alpine
```

## Run

* Run DB

```
docker run -d -p 5432:5432 --name devdb -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=dev -e POSTGRES_DB=dev postgres:10.3-alpine
```

## Run Postgrest

```
cd db
postgrest postgrest.conf
```

## Run server

```
cd server
go run \*.go
```

## Run client

```
cd client
yarn start
```

```

```
