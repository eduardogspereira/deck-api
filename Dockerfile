FROM golang:1.15-alpine

RUN apk add curl

RUN mkdir /deck-api
ADD . /deck-api
WORKDIR /deck-api

RUN go get
RUN go build -o deck-api

CMD ["./deck-api"]