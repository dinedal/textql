FROM golang:1.10

# install sqlite3 for option "-console"
RUN apt-get update && apt-get install -y sqlite3

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /tmp

ENTRYPOINT ["textql"]
