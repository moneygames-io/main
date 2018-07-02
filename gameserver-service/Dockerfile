FROM golang:1.8

WORKDIR /go/src/app
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . .

RUN dep ensure
RUN go build

CMD ["app"]
