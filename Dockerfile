FROM golang:latest

RUN apt-get update

RUN mkdir -p /home/ApiServer

WORKDIR /home/Apiserver

ADD src .

CMD["go run main.go"]