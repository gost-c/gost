# gost
[![Go Report Card](https://goreportcard.com/badge/github.com/gost-c/gost)](https://goreportcard.com/report/github.com/gost-c/gost)
[![CircleCI Build Status](https://circleci.com/gh/gost-c/gost.svg?style=shield)](https://circleci.com/gh/gost-c/gost)
[![](https://images.microbadger.com/badges/version/zcong/gost.svg)](https://microbadger.com/images/zcong/gost "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/image/zcong/gost.svg)](https://microbadger.com/images/zcong/gost "Get your own image badge on microbadger.com")

> simple gist like service in go

## Docker

```sh
# run a mysql
$ docker run --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=true -v `pwd`/create-gost.sql:/docker-entrypoint-initdb.d/create-db.sql  -d mysql
# start services
$ docker run --name gost -d --link mysql:mysql -p 8000:8000 -e JWT_SECRET=secret \
-e GIN_MODE=debug \
-e MYSQL_DB_URL="root:@tcp(mysql:3306)/gost?charset=utf8&parseTime=True&loc=Local" \
zcong/gost
```

## Docs

see [http://gost-docs.congz.pw](http://gost-docs.congz.pw)

## License

MIT &copy; zcong1993
