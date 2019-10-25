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
CREATE TABLE beers (id BIGSERIAL PRIMARY KEY, name TEXT, tastingID int)
```

Install dependencies
```
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/google/uuid
```

## Run locally
```
> go build
> ./gobeer
```

## Deploy to beer.abilia-gbg.se

* Stop service `sudo service gobeer stop`
* Copy `./gobeer` to `/build/beer/gobeer`
* Start service `sudo service gobeer start`

#### Todo on beer.abilia-gbg.se
- [x] Setup database

`docker run -p 5454:5432 --name gobeer -e POSTGRES_PASSWORD=secret -dt postgres`

- [x] Create systemd service

```
[Unit]
Description=gobeer
After=docker.service

[Service]
Type=simple
Restart=on-failure
RestartSec=5
ExecStart=/build/beer/gobeer

[Install]
WantedBy=multi-user.target
```

- [x] Create nginx configuration

```
server {
    listen 80;
    server_name beer.abilia-gbg.se;

    root /build/beer;

    error_log /build/error.beer.log;
    access_log /build/access.beer.log;

    location / {
        proxy_pass http://127.0.0.1:8000;
    }
}
```

- [x] Use cert for beer.abilia-gbg.se

## Deploy to google app engine
Google App Engine at Abilia innovation project

- [ ] Connect GO Google App Engine with GloudSQL
- [ ] Setup proxy for development