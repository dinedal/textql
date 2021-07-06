.PHONY: all test clean man godep fast release install

all: textql

textql: test
	go build -ldflags "-X main.VERSION=`cat VERSION` -s" -o ./build/textql ./textql/main.go

fast:
	go build -i -ldflags "-X main.VERSION=`cat VERSION`-dev -s" -o ./build/textql ./textql/main.go

test:
	go test ./...

clean:
	rm -fr ./build

release: textql
	git tag -a `cat VERSION`
	git push origin `cat VERSION`

install: test
	go install -ldflags "-X main.VERSION=`cat VERSION` -s" ./textql/main.go

man:
	ronn man/textql.1.ronn
