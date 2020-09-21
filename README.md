# Go Gin Jwt Authorizaion Example

Go Gin Gorm REST API example using most useful features.

## Required

- Go
- Mysql
- Redis

## Setup & Installation

```
$ cd $GOPATH/src

$ git clone https://github.com/MuhammadTamzid/go-gin-jwt-authorization-example.git
```

### Add ENV variables

You need to set ENV, also you can add `.env` file in project directory.

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=123456
DB_NAME=online_course

REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=

TOKEN_SECRET=98hbun98hsdfsdwesdfs
ACCESS_TIME_EXPIRED=900000000000
REFRESH_TIME_EXPIRED=604800000000000

PORT=4000
```

### Fix PATH for Swagger

run `$HOME/go/bin/swag` or add `$HOME/go/bin` to your `$PATH`

### SQL script

Run [SQL script](https://github.com/MuhammadTamzid/go-gin-jwt-authorization-example/blob/master/docs/sql/schema.sql)

### Run application

```
$ go run main.go
```

## API documents

After run application, browser to http://localhost:4000/swagger/index.html, you can see Swagger Api documents.
![image](https://i.imgur.com/6b4IZQo.jpg)
