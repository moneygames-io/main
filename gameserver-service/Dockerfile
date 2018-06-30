FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -v ./...
RUN go build

CMD ["app"]
