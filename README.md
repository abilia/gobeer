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
CREATE TABLE tastings (id BIGSERIAL PRIMARY KEY, name TEXT);
```

Install dependencies
```
go get github.com/gorilla/mux
go get github.com/lib/pq
```

## Run locally
```
> go build
> ./gobeer
```

## Deploy to beer.abilia-gbg.se

#### Todo on beer.abilia-gbg.se
- [ ] Create service
- [ ] Create nginx configuration
- [ ] Update DNS
- [ ] Use cert for beer.abilia-gbg.se

For be able to deploy with:
* Copy `./gobeer` to `/var/www/gobeer`
* Restart service `sudo systemctl restart gobeer`


## Deploy to google app engine
Running in Abilia innovation project in a google app engine