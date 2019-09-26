# gobeer

Server for [flutterbeer](https://github.com/abilia/flutterbeer)

Running (soon) at https://innovation-242008.appspot.com/

## Setup

### Prerequisits
* Install [go](https://golang.org/doc/install)
* Install [postgresql](https://www.postgresql.org/download/)

### Development
Check out the code

Suggested IDE: [vscode](https://code.visualstudio.com/download)

Create database
```
CREATE USER gobeeruser WITH PASSWORD 'thisisthepassword';
CREATE DATABASE beerdb OWNER gobeeruser;
```

Create table(s)
```
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, username TEXT);
```

Install dependencies
```
go get github.com/gorilla/mux
go get github.com/lib/pq
```

### Run locally
```
> go run main.go database.go models.go api.go
```

## Deploy
Running in Abilia innovation project in a google app engine