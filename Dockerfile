FROM golang:latest AS build

ENV PROJECT /go/src/github.com/gost-c/gost

RUN mkdir -p $PROJECT
COPY . $PROJECT

WORKDIR ${PROJECT}

RUN make install.dev \
 && CGO_ENABLED=0 make build

FROM alpine:3.6

ENV PROJECT /go/src/github.com/gost-c/gost

WORKDIR /opt/gost

COPY --from=build $PROJECT/bin/* /usr/local/bin/

RUN ln -s /usr/local/bin/gost* / \
 && ln -s /usr/local/bin/gost* /bin/

CMD ["gost"]
