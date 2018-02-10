FROM golang:latest AS build

ENV PROJECT /go/src/github.com/gost-c/gost
# install xz
RUN apt-get update && apt-get install -y \
    xz-utils \
&& rm -rf /var/lib/apt/lists/*
# install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx
RUN mkdir -p $PROJECT
COPY . $PROJECT

WORKDIR ${PROJECT}

RUN make install.dev \
 && CGO_ENABLED=0 make build
# strip and compress the binary
RUN strip --strip-unneeded ./bin/gost
RUN upx ./bin/gost

FROM alpine:3.7

ENV PROJECT /go/src/github.com/gost-c/gost

WORKDIR /opt/gost

COPY --from=build $PROJECT/bin/* /usr/local/bin/

RUN ln -s /usr/local/bin/gost* / \
 && ln -s /usr/local/bin/gost* /bin/

CMD ["gost"]
