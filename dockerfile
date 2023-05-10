#syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main

EXPOSE 8082

CMD [ "./main" ]