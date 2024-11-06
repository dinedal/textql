.PHONY: all test clean man godep fast release install

all: textql

textql: test textql_macos_arm64 textql_macos textql_linux textql.exe

textql_macos_arm64:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -ldflags "-X main.VERSION=`cat VERSION` -s" -o ./build/textql_macos_arm64 ./textql/main.go

textql_macos:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "-X main.VERSION=`cat VERSION` -s" -o ./build/textql_macos ./textql/main.go

textql_linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "-X main.VERSION=`cat VERSION` -s" -o ./build/textql_linux ./textql/main.go

textql.exe:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "-X main.VERSION=`cat VERSION` -s" -o ./build/textql.exe ./textql/main.go

fast:
	go build -i -ldflags "-X main.VERSION=`cat VERSION`-dev -s" -o ./build/textql ./textql/main.go

test: test-tools
	go test -v -race -coverprofile=c.out -covermode=atomic ./... | go-junit-report -iocopy -set-exit-code -out unit-test-report.xml

test-tools: ## Install tools required for generating test reports
	go install github.com/jstemmer/go-junit-report/v2@v2.1.0

clean:
	rm -fr ./build

release: textql
	git tag -a `cat VERSION`
	git push origin `cat VERSION`

install: test
	go install -ldflags "-X main.VERSION=`cat VERSION` -s" ./textql/main.go

man:
	ronn man/textql.1.ronn
