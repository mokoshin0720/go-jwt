# sample

## Usage
1. run server

Docker
```
$ make migrate
$ make run
```

Local
```
$ mysql -u root
mysql> create database sample

$ set -a; source .env.migrate; set +a; # set environment variable
$ dbmate up
$ set -a; source .env; set +a; # set environment variable
$ go run github.com/ispec-inc/sample/cmd/api/server
```

2. healthcheck

```
$ curl localhost:9000/health
{"message":"success"}
```
