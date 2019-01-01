FROM golang:1.11 AS build
WORKDIR /mnt
COPY . .
RUN CGO_ENABLED=0 go build -o ./bin/gost main.go

FROM alpine:3.7
WORKDIR /opt
RUN apk add --no-cache ca-certificates
COPY --from=build /mnt/bin/* /usr/bin/
EXPOSE 9393
CMD ["gost"]
