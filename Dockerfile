FROM golang:1.10

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /tmp

ENTRYPOINT ["textql"]
