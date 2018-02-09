FROM golang:latest@sha256:d1056842395f50cb8994764bac51ee07c4ca69575d649cd623ae2b62c4632a58 AS build

ENV PROJECT /go/src/github.com/gost-c/gost

RUN mkdir -p $PROJECT
COPY . $PROJECT

WORKDIR ${PROJECT}

RUN make install.dev \
 && CGO_ENABLED=0 make build

FROM alpine:3.6@sha256:3d44fa76c2c83ed9296e4508b436ff583397cac0f4bad85c2b4ecc193ddb5106

ENV PROJECT /go/src/github.com/gost-c/gost

WORKDIR /opt/gost

COPY --from=build $PROJECT/bin/* /usr/local/bin/

RUN ln -s /usr/local/bin/gost* / \
 && ln -s /usr/local/bin/gost* /bin/

CMD ["gost"]
