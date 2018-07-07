.PHONY: all test clean man godep fast release install

GO15VENDOREXPERIMENT=1
GODEP := $(shell command -v dep 2> /dev/null)

all: textql

textql: deps test
	go build -ldflags "-X main.VERSION=`cat VERSION` -s" -o ./build/textql ./textql/main.go

fast:
	go build -i -ldflags "-X main.VERSION=`cat VERSION`-dev -s" -o ./build/textql ./textql/main.go

deps: godep
	dep ensure

godep:
ifndef GODEP
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif

test:
	go test ./inputs/
	go test ./storage/

clean:
	rm -fr ./build

release: textql
	git tag -a `cat VERSION`
	git push origin `cat VERSION`

install: deps test
	go install -ldflags "-X main.VERSION=`cat VERSION` -s" ./textql/main.go

man:
	ronn man/textql.1.ronn
