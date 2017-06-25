FROM golang:latest
RUN go get github.com/DennisDenuto/property-price-collector
WORKDIR /go/src/github.com/DennisDenuto/property-price-collector

ENTRYPOINT ["go", "run", "main.go"]