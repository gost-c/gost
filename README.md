# gost
[![Go Report Card](https://goreportcard.com/badge/github.com/gost-c/gost)](https://goreportcard.com/report/github.com/gost-c/gost)
[![CircleCI Build Status](https://circleci.com/gh/gost-c/gost.svg?style=shield)](https://circleci.com/gh/gost-c/gost)
[![](https://images.microbadger.com/badges/version/zcong/gost.svg)](https://microbadger.com/images/zcong/gost "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/image/zcong/gost.svg)](https://microbadger.com/images/zcong/gost "Get your own image badge on microbadger.com")

> simple gist like service in go

## Docker

```sh
# run a mongo
$ docker run --name mongo -d mongo
# start services
$ docker run --name gost -d --link mongo:mongo -p 8000:9393 -e JWTKEY=secret \
-e ENV=debug \
-e MONGOURL="mongo" \
zcong/gost
```

## Docs

see [http://gost-docs.congz.pw](http://gost-docs.congz.pw)

## License

MIT &copy; zcong1993
