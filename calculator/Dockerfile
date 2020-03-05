FROM alpine:3.11

WORKDIR /usr/local
ENV GOPATH /usr/local/go

RUN apk add --no-cache go git collectd collectd-network supervisor

COPY src $GOPATH/src
COPY config/collectd.conf /etc/collectd/
COPY config/supervisord.conf /etc/

RUN go get -u github.com/gorilla/mux
RUN go build -o bin/calculator go/src/calculator.go

ENTRYPOINT [ "supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf" ]
