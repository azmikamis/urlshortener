FROM golang:1.5.2

ADD . /go/src/github.com/azmikamis/urlshortener

RUN go install github.com/azmikamis/urlshortener

ENTRYPOINT /go/bin/urlshortener

EXPOSE 8080
