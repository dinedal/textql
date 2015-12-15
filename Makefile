all: textql

textql: deps test
	go build -ldflags "-X main.VERSION=`cat VERSION`" -o ./build/textql ./textql/main.go

fast:
	go build -i -ldflags "-X main.VERSION=`cat VERSION`-dev" -o ./build/textql ./textql/main.go

deps: glide
	./glide install

glide:
ifeq ($(shell uname),Darwin)
	curl -L https://github.com/Masterminds/glide/releases/download/0.5.0/glide-darwin-amd64.zip -o glide.zip
	unzip glide.zip
	mv ./darwin-amd64/glide ./glide
	rm -fr ./darwin-amd64
	rm ./glide.zip
else
	curl -L https://github.com/Masterminds/glide/releases/download/0.5.0/glide-linux-386.zip -o glide.zip
	unzip glide.zip
	mv ./linux-386/glide ./glide
	rm -fr ./linux-386
	rm ./glide.zip
endif

test:
	go test ./inputs/
	go test ./storage/

clean:
	rm ./glide
	rm -fr ./build

release: textql
	git tag -a `cat VERSION`
	git push origin `cat VERSION`

install: deps test
	go install -ldflags "-X main.VERSION=`cat VERSION`" ./textql/textql.go
