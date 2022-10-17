FROM golang:1.16.0-buster AS backend-builder
COPY . /app
RUN cd /app && env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-linkmode external -extldflags -static -s -w' -o ovpn-admin

FROM alpine:3.14
WORKDIR /app
COPY --from=backend-builder /app/ovpn-admin /app
RUN apk add --update bash easy-rsa  && \
    ln -s /usr/share/easy-rsa/easyrsa /usr/local/bin && \
    rm -rf /tmp/* /var/tmp/* /var/cache/apk/* /var/cache/distfiles/*
RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone
RUN apk del tzdata
