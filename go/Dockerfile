FROM alpine:3.11

WORKDIR /usr/local
ENV GOPATH /usr/local/go

COPY src $GOPATH/src

RUN apk add --no-cache go git collectd
RUN go get -u github.com/gorilla/mux
RUN go build -o bin/calculator go/src/calculator.go

COPY config/collectd.conf /etc/collectd/

RUN apk add --no-cache collectd-network
RUN apk add --no-cache supervisor

COPY config/supervisord.conf /etc/

ENTRYPOINT [ "supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf" ]
