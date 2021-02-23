FROM golang:1.16-alpine

WORKDIR /hasGoriraAPI

RUN apk update  \
&& apk add --no-cache git  \
&& go get -u github.com/cosmtrek/air

EXPOSE 5050

CMD air