FROM golang:latest

RUN apt-get update

RUN go get github.com/gorilla/mux

RUN mkdir -p /home/ApiServer

WORKDIR /home/Apiserver

ADD src .

ENTRYPOINT ["./run.sh"]
