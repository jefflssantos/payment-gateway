FROM golang:1.22

WORKDIR /go/src

CMD ["tail", "-f", "/dev/null"]
