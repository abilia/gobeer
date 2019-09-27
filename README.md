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
- [ x ] Setup database

`docker run -p 5454:5432 --name gobeer -e POSTGRES_PASSWORD=secret -dt postgres`

- [ ] Create systemd service
- [ x ] Create nginx configuration
- [ ] Use cert for beer.abilia-gbg.se

For be able to deploy with:
* Copy `./gobeer` to `/var/www/beer/gobeer`
* Restart service `sudo systemctl restart gobeer`


## Deploy to google app engine
Google App Engine at Abilia innovation project

- [ ] Connect GO Google App Engine with GloudSQL
- [ ] Setup proxy for development