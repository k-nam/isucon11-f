FROM golang:1.16.5-buster
WORKDIR /webapp/go

RUN apt-get update && apt-get install -y wget zip

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && go get -u github.com/cosmtrek/air


COPY go.mod go.sum ./
RUN go mod download

COPY air.toml .
