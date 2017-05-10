FROM golang:1.7.5-alpine

ENV REDIS_ADDRESS="localhost:6379"

EXPOSE 8080


RUN apk --update add git

ADD . /go/src/app

WORKDIR /go/src/app

RUN go get -d -v

RUN go install

CMD ["/go/bin/app"]