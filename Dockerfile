FROM zcong/golang:1.10.3 AS build
WORKDIR /go/src/github.com/gost-c/gost
COPY . .
RUN dep ensure -vendor-only -v && \
    CGO_ENABLED=0 go build -o ./bin/gost main.go

FROM alpine:3.7
WORKDIR /opt
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/gost-c/gost/bin/* /usr/bin/
EXPOSE 9393
CMD ["gost"]
