FROM golang:1.15.2-alpine

# go getコマンドは裏でgitコマンドを使用しているのでgitを入れる
RUN apk update && apk add git

RUN mkdir /go/src/app

WORKDIR /go/src/app
