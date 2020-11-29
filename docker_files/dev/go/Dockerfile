FROM golang:latest as builder

WORKDIR /go/src/app/server_side

# Go Modulesの有効化
ENV GO111MODULE=on

# なぜかgooseだけ、Go modulesが機能しないため、個別インストール
RUN go get bitbucket.org/liamstask/goose/cmd/goose