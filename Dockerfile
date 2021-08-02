FROM golang:1.16-alpine

WORKDIR /app/

ADD . /app

EXPOSE 8080
